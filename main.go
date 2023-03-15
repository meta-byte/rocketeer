package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// https://mholt.github.io/json-to-go/ to save time here
type Launch struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		ID     string `json:"id"`
		URL    string `json:"url"`
		Slug   string `json:"slug"`
		Name   string `json:"name"`
		Status struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Abbrev      string `json:"abbrev"`
			Description string `json:"description"`
		} `json:"status"`
		LastUpdated           time.Time `json:"last_updated"`
		Net                   time.Time `json:"net"`
		WindowEnd             time.Time `json:"window_end"`
		WindowStart           time.Time `json:"window_start"`
		Probability           any       `json:"probability"`
		Holdreason            string    `json:"holdreason"`
		Failreason            string    `json:"failreason"`
		Hashtag               any       `json:"hashtag"`
		LaunchServiceProvider struct {
			ID   int    `json:"id"`
			URL  string `json:"url"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"launch_service_provider"`
		Rocket struct {
			ID            int `json:"id"`
			Configuration struct {
				ID       int    `json:"id"`
				URL      string `json:"url"`
				Name     string `json:"name"`
				Family   string `json:"family"`
				FullName string `json:"full_name"`
				Variant  string `json:"variant"`
			} `json:"configuration"`
		} `json:"rocket"`
		Mission struct {
			ID               int    `json:"id"`
			Name             string `json:"name"`
			Description      string `json:"description"`
			LaunchDesignator any    `json:"launch_designator"`
			Type             string `json:"type"`
			Orbit            struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Abbrev string `json:"abbrev"`
			} `json:"orbit"`
		} `json:"mission"`
		Pad struct {
			ID        int    `json:"id"`
			URL       string `json:"url"`
			AgencyID  int    `json:"agency_id"`
			Name      string `json:"name"`
			InfoURL   any    `json:"info_url"`
			WikiURL   any    `json:"wiki_url"`
			MapURL    any    `json:"map_url"`
			Latitude  string `json:"latitude"`
			Longitude string `json:"longitude"`
			Location  struct {
				ID                int    `json:"id"`
				URL               string `json:"url"`
				Name              string `json:"name"`
				CountryCode       string `json:"country_code"`
				MapImage          string `json:"map_image"`
				TotalLaunchCount  int    `json:"total_launch_count"`
				TotalLandingCount int    `json:"total_landing_count"`
			} `json:"location"`
			MapImage                  string `json:"map_image"`
			TotalLaunchCount          int    `json:"total_launch_count"`
			OrbitalLaunchAttemptCount int    `json:"orbital_launch_attempt_count"`
		} `json:"pad"`
		WebcastLive                    bool   `json:"webcast_live"`
		Image                          string `json:"image"`
		Infographic                    any    `json:"infographic"`
		Program                        []any  `json:"program"`
		OrbitalLaunchAttemptCount      int    `json:"orbital_launch_attempt_count"`
		LocationLaunchAttemptCount     int    `json:"location_launch_attempt_count"`
		PadLaunchAttemptCount          int    `json:"pad_launch_attempt_count"`
		AgencyLaunchAttemptCount       int    `json:"agency_launch_attempt_count"`
		OrbitalLaunchAttemptCountYear  int    `json:"orbital_launch_attempt_count_year"`
		LocationLaunchAttemptCountYear int    `json:"location_launch_attempt_count_year"`
		PadLaunchAttemptCountYear      int    `json:"pad_launch_attempt_count_year"`
		AgencyLaunchAttemptCountYear   int    `json:"agency_launch_attempt_count_year"`
	} `json:"results"`
}

func launchHandler() {
	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://lldev.thespacedevs.com/2.2.0/launch/upcoming/", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject Launch
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject.Results[0])
}

var (
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
	commands       = []*discordgo.ApplicationCommand{
		{
			Name:        "hello",
			Description: "Hello world!",
		},
		{
			Name:        "launch",
			Description: "Get info on next launch",
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"hello": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hello world",
				},
			})
		},
		"launch": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			launchHandler()
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "retrieved launch data",
				},
			})
		},
	}
	session *discordgo.Session
)

func init() { flag.Parse() }

func init() {
	var err error
	godotenv.Load(".env")
	token := os.Getenv("BOT_TOKEN")
	session, err = discordgo.New("Bot " + token)

	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func init() {
	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if name, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			name(s, i)
		}
	})
}

func main() {

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Adding commands...")
	})

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err := session.Open()
	if err != nil {
		log.Fatalf("unable to open session...", err)
	}

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
	defer session.Close()

	fmt.Println("Rocketeer is online!")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	if *RemoveCommands {
		log.Println("Removing commands...")

		for _, v := range registeredCommands {
			err := session.ApplicationCommandDelete(session.State.User.ID, "", v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}
}
