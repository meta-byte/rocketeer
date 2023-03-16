package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/meta-byte/rocketeer-discord-bot/bot"
)

func init() {
	// flag.Parse()
}

func main() {
	// Initialize Discord session
	godotenv.Load(".env")
	token := os.Getenv("BOT_TOKEN")
	appID := os.Getenv("APP_ID")
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("invalid bot parameters: %v", err)
	}
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	// Create bot instance
	bot := bot.NewBot(session, appID)

	// Register commands
	err = bot.RegisterCommands()
	if err != nil {
		log.Fatal(err)
	}

	//Register command handlers
	bot.RegisterHandlers()

	//Discord has responded that the bot is ready
	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("rocketeer is online!")
	})

	err = session.Open()
	if err != nil {
		log.Fatal("unable to open session...", err)
	}

	defer session.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
	fmt.Println("\ninterrupt received")

	bot.UnregisterCommands()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("rocketeer is now offline.")
}
