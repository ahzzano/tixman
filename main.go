package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var (
	Token = flag.String("token", "", "Auth Token")
	App   = flag.String("app", "", "Application ID")
	Guild = flag.String("guild", "", "GuildID")
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "echo",
		Description: "Hello World",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "message",
				Description: "what",
				Required:    true,
			},
		},
	},
}

func setupHandlers(sess *discordgo.Session) {
	sess.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as %s", r.User.String())
	})
}

func main() {
	flag.Parse()

	if *Token == "" {
		log.Fatalf("No token was supplied. Supply the token via the token flag")
		return
	}

	session, err := discordgo.New("Bot " + *Token)

	if err != nil {
		log.Fatalf("Unable to open session: %s", err)
	}

	setupHandlers(session)

	err = session.Open()

	if err != nil {
		log.Fatalf("Could not open session: %s", err)
	}

	log.Printf("Session is now currently running")

	sigch := make(chan os.Signal, 1)

	signal.Notify(sigch, os.Interrupt)
	<-sigch

	err = session.Close()

	if err != nil {
		log.Printf("Could not close gracefully: %s", err)
	}

	log.Printf("Session has now closed. Saynora...")

}
