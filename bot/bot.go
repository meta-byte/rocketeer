package bot

import (
	"fmt"
	"log"
	"strings"

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

//	type Command struct {
//		Name        string
//		Description string
//		Handler     func(session *discordgo.Session, i *discordgo.InteractionCreate)
//	}
//
// TODO: split commands and handler functions for readability and maintenance.
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
			{
				Name:        "dm",
				Description: "I'll send you a dm with launch info.",
			},
		},
		CommandHandlers: map[string]func(session *discordgo.Session, i *discordgo.InteractionCreate){
			"hello": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Hello!",
					},
				})
			},
			"launch": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
				launch, err := api.GetLaunch()
				embed := api.BuildLaunchEmbed(launch, s, i)
				if err != nil {
					log.Fatal(err)
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "I was unable to retrieve launch data.",
						},
					})
				}

				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Here's the next upcoming launch! 🚀",
						Embeds: []*discordgo.MessageEmbed{
							&embed,
						},
					},
				})

				message, err := s.InteractionResponse(i.Interaction)

				err = s.MessageReactionAdd(message.ChannelID, message.ID, "🚀")

			},
			"dm": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
				user, err := s.User(i.Member.User.ID)

				launch, err := api.GetLaunch()
				embed := api.BuildLaunchEmbed(launch, s, i)
				if err != nil {
					log.Fatal(err)
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "I was unable to retrieve launch data.",
						},
					})
				}
				privateChannel, err := s.UserChannelCreate(user.ID)
				if err != nil {
					log.Fatalf("Error creating private channel: %v", err)
				}

				_, err = s.ChannelMessageSendEmbed(privateChannel.ID, &embed)

				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "I've sent you a direct message",
					},
				})
			},
		},
	}
}

// TODO: how to do we handle reactions
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
