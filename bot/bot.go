package bot

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/meta-byte/rocketeer-discord-bot/database"
	"github.com/meta-byte/rocketeer-discord-bot/types"
)

type Bot struct {
	Session            *discordgo.Session
	DB                 database.Database
	ApplicationID      string
	Commands           []*discordgo.ApplicationCommand
	RegisteredCommands []*discordgo.ApplicationCommand
	CommandHandlers    map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func NewBot(session *discordgo.Session, applicationID string, db database.Database) *Bot {
	return &Bot{
		Session:       session,
		DB:            db,
		ApplicationID: applicationID,
		Commands: []*discordgo.ApplicationCommand{
			{
				Name:        "hello",
				Description: "Hello!",
			},
			{
				Name:        "launch",
				Description: "Get info on an upcoming launch.",
			},
		},
		CommandHandlers: map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
			"hello": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Hello!",
					},
				})
			},
			"launch": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
				launch := upcomingLaunch(db)
				embed := buildLaunchEmbed(launch, s, i)

				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Here's the next upcoming launch! 🚀",
						Embeds: []*discordgo.MessageEmbed{
							&embed,
						},
					},
				})
				message, _ := s.InteractionResponse(i.Interaction)
				s.MessageReactionAdd(message.ChannelID, message.ID, "🚀")

			},
		},
	}
}

func (b *Bot) RegisterCommands() error {
	log.Println("adding commands...")

	for _, command := range b.Commands {
		_, err := b.Session.ApplicationCommandCreate(b.ApplicationID, "", command)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v /n", command.Name, err)
			return err
		}
	}

	log.Println("fetching registered commands...")
	registeredCommands, err := b.Session.ApplicationCommands(b.ApplicationID, "")
	if err != nil {
		log.Fatalf("Error fetching registered commands: %v", err)
		registeredCommands = nil
		return err
	}

	b.RegisteredCommands = registeredCommands
	return nil
}

func (b *Bot) UnregisterCommands() error {

	log.Println("Removing commands...")

	for _, command := range b.RegisteredCommands {
		err := b.Session.ApplicationCommandDelete(b.ApplicationID, "", command.ID)
		if err != nil {
			log.Fatalf("Cannot delete '%v' command: %v /n", command.Name, err)
			return err
		}
	}
	return nil
}

func (b *Bot) RegisterHandlers() {
	log.Println("registering command handlers...")

	b.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		handler, ok := b.CommandHandlers[i.ApplicationCommandData().Name]
		if !ok {
			log.Fatal("Unable to register command handlers.")
			return
		}
		handler(s, i)
	})
	b.Session.AddHandler(func(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
		channel := r.ChannelID
		user := r.UserID
		messageID := r.MessageID

		message, err := s.ChannelMessage(channel, messageID)
		if err != nil {
			log.Fatalf("Unable to retrieve message: %v", err)
		}

		content := message.Content
		if (s.State.User.ID != user) && (r.Emoji.Name == "🚀") && (strings.Contains(content, "🚀")) {
			privateChannel, err := s.UserChannelCreate(user)
			if err != nil {
				log.Fatalf("Error creating private channel: %v", err)
			}

			_, err = s.ChannelMessageSend(privateChannel.ID, "I received your reaction to my message.")
		} else {
			fmt.Println("reaction was created by rocketeer...ignoring")
		}

	})

}

func upcomingLaunch(db database.Database) *types.Launch {
	launches := db.FetchLaunches()
	if len(launches) == 0 {
		return nil
	}
	//sort launches
	sort.Slice(launches, func(i, j int) bool {
		durationI := launches[i].Net.Sub(time.Now())
		durationJ := launches[j].Net.Sub(time.Now())
		return durationI < durationJ
	})

	return &launches[0]
}

func buildLaunchEmbed(l *types.Launch, s *discordgo.Session, i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	duration := l.Net.Sub(time.Now())
	days := int(duration.Hours() / 24)
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	countdown := fmt.Sprintf("%d:%d:%d:%d", days, hours, minutes, seconds)
	thumbnailURL := ""

	if len(l.Program) > 0 {
		thumbnailURL = l.Program[0].ImageURL
	}

	return discordgo.MessageEmbed{
		Title:       l.Name,
		Description: l.Mission.Description + "\n",
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
				Value:  l.Status.Abbrev,
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
				Value:  l.WindowStart.Format("01-02-2006 15:04:05") + "Z",
				Inline: true,
			},
			{
				Name:   "Launch Window Closes",
				Value:  l.WindowEnd.Format("01-02-2006 15:04:05") + "Z",
				Inline: true,
			},
			{
				Name:   "Pad",
				Value:  fmt.Sprintf("[%v](%v)", l.Pad.Name, l.Pad.MapURL),
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
				Value:  l.Rocket.Configuration.FullName,
				Inline: true,
			},
			{
				Name:   "Launch Provider",
				Value:  l.LaunchServiceProvider.Name,
				Inline: true,
			},
			{
				Name:   "Provider Type",
				Value:  l.LaunchServiceProvider.Type,
				Inline: true,
			},
			{
				Name:   "---Mission---",
				Inline: false,
			},
			{
				Name:   "Program",
				Value:  l.LaunchServiceProvider.Type,
				Inline: true,
			},
			{
				Name:   "Mission Type",
				Value:  l.Mission.Type,
				Inline: true,
			},
			{
				Name:   "Orbit",
				Value:  l.Mission.Orbit.Name,
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: l.Image,
		},
		Footer: &discordgo.MessageEmbedFooter{
			IconURL: i.Member.Avatar,
			Text:    l.Status.Description,
		},
		Color: 302553,
	}
}
