package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	fmt.Println("Hello World!")

	token := os.Getenv("TOKEN")

	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Errorf("Error: Unable to create bot %s", err)
		return
	}

	fmt.Println("%s", bot)
}
