package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/meta-byte/rocketeer-discord-bot/types"
)

func makeGETRequest(endpoint string, target interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Move business logic for getting the singular next launch to here. Setup separate function for multiple launches etc.
func GetLaunch() (launch *types.Launch, err error) {
	endpoint := os.Getenv("ENDPOINT")
	err = makeGETRequest(endpoint, &launch)
	if err != nil {
		return nil, err
	}
	if len(launch.Results) > 0 {
		return launch, nil
	}
	return nil, errors.New("no launch results")
}

func BuildLaunchEmbed(l *types.Launch, s *discordgo.Session, i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	duration := l.Results[0].WindowStart.Sub(time.Now())
	days := int(duration.Hours() / 24)
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	countdown := fmt.Sprintf("%d:%d:%d:%d", days, hours, minutes, seconds)
	thumbnailURL := ""

	if len(l.Results[0].Program) > 0 {
		thumbnailURL = l.Results[0].Program[0].ImageURL
	}

	return discordgo.MessageEmbed{
		Title:       l.Results[0].Name,
		Description: l.Results[0].Mission.Description + "\n",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: thumbnailURL,
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "",
				Inline: false,
			},
			{
				Name:   "---Launch---",
				Inline: false,
			},
			{
				Name:   "Status",
				Value:  l.Results[0].Status.Abbrev,
				Inline: true,
			},
			{
				Name:   "Approx Countdown",
				Value:  countdown,
				Inline: true,
			},
			{
				Name:   "Current Time",
				Value:  time.Now().UTC().Format("01-02-2006 15:04:05") + "Z",
				Inline: true,
			},
			{
				Name:   "Launch Window Opens",
				Value:  l.Results[0].WindowStart.Format("01-02-2006 15:04:05") + "Z",
				Inline: true,
			},
			{
				Name:   "Launch Window Closes",
				Value:  l.Results[0].WindowEnd.Format("01-02-2006 15:04:05") + "Z",
				Inline: true,
			},
			{
				Name:   "Pad",
				Value:  fmt.Sprintf("[%v](%v)", l.Results[0].Pad.Name, l.Results[0].Pad.MapURL),
				Inline: true,
			},
			{
				Name:   "",
				Inline: false,
			},
			{
				Name:   "---Vehicle---",
				Inline: false,
			},
			{
				Name:   "Rocket",
				Value:  l.Results[0].Rocket.Configuration.FullName,
				Inline: true,
			},
			{
				Name:   "Launch Provider",
				Value:  l.Results[0].LaunchServiceProvider.Name,
				Inline: true,
			},
			{
				Name:   "Provider Type",
				Value:  l.Results[0].LaunchServiceProvider.Type,
				Inline: true,
			},
			{
				Name:   "---Mission---",
				Inline: false,
			},
			{
				Name:   "Program",
				Value:  l.Results[0].LaunchServiceProvider.Type,
				Inline: true,
			},
			{
				Name:   "Mission Type",
				Value:  l.Results[0].Mission.Type,
				Inline: true,
			},
			{
				Name:   "Orbit",
				Value:  l.Results[0].Mission.Orbit.Name,
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: l.Results[0].Image,
		},
		Footer: &discordgo.MessageEmbedFooter{
			IconURL: i.Member.Avatar,
			Text:    l.Results[0].Status.Description,
		},
		Color: 302553,
	}
}
