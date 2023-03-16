package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/meta-byte/rocketeer-discord-bot/api"
)

type Bot struct {
	Session         *discordgo.Session
	ApplicationID   string
	Commands        []*discordgo.ApplicationCommand
	CommandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
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
					Image: &discordgo.MessageEmbedImage{
						URL: responseObject.Results[0].Image,
					},
					//TODO: Add footer to embed
					/*
						Footer: &discordgo.MessageEmbedFooter{
							Text: footerString,
						},
					*/
					Color: 15883269,
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
	fmt.Println("adding commands...")

	for _, command := range b.Commands {
		_, err := b.Session.ApplicationCommandCreate(b.ApplicationID, "", command)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v /n", command.Name, err)
			return err
		}
	}
	return nil
}

func (b *Bot) UnregisterCommands() {
	log.Println("Removing commands...")

	for _, v := range b.Commands {
		err := b.Session.ApplicationCommandDelete(b.ApplicationID, "", v.ID)
		if err != nil {
			log.Fatalf("Cannot delete '%v' command: %v /n", v.Name, err)
		}
	}
}

func (b *Bot) RegisterHandlers() {
	fmt.Println("registering command handlers...")

	b.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		handler, ok := b.CommandHandlers[i.ApplicationCommandData().Name]
		if !ok {
			log.Fatal("Unable to register command handlers.")
			return
		}
		handler(s, i)
	})
}
