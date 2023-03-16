package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/meta-byte/rocketeer-discord-bot/api"
)

type Bot struct {
	Session            *discordgo.Session
	ApplicationID      string
	Commands           []*discordgo.ApplicationCommand
	RegisteredCommands []*discordgo.ApplicationCommand
	CommandHandlers    map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func NewBot(session *discordgo.Session, applicationID string) *Bot {
	return &Bot{
		Session:       session,
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
				responseObject, err := api.GetLaunches()
				duration := responseObject.Results[0].WindowStart.Sub(time.Now())

				days := int(duration.Hours() / 24)
				hours := int(duration.Hours()) % 24
				minutes := int(duration.Minutes()) % 60
				seconds := int(duration.Seconds()) % 60

				countdown := fmt.Sprintf("%d days, %d hours, %d minutes, %d seconds\n", days, hours, minutes, seconds)

				if err != nil {
					log.Fatal(err)
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "I was unable to retrieve launch data.",
						},
					})
				}
				// footerString := fmt.Sprintf("Requested by %v", i.Member)
				embed := discordgo.MessageEmbed{
					Title:       responseObject.Results[0].Name,
					Description: responseObject.Results[0].Mission.Description,
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "Status",
							Value:  responseObject.Results[0].Status.Abbrev,
							Inline: true,
						},
						{
							Name:   "Approx Countdown",
							Value:  countdown,
							Inline: true,
						},
						{
							Name:   "Launch Window Opens",
							Value:  responseObject.Results[0].WindowStart.Format("01-02-2006 15:04:05") + "Z",
							Inline: true,
						},
						{
							Name:   "Current Time",
							Value:  time.Now().UTC().Format("01-02-2006 15:04:05") + "Z",
							Inline: true,
						},
					},
					Image: &discordgo.MessageEmbedImage{
						URL: responseObject.Results[0].Image,
					},
					//TODO: Add footer to embed
					/*
						Footer: &discordgo.MessageEmbedFooter{
							Text: footerString,
						},
					*/
					Color: 302553,
				}
				// session.ChannelMessageSendEmbed(i.ChannelID, &embed)
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							&embed,
						},
					},
				})
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
}
