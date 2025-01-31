package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("TIX_TOKEN")
	println(token)

	bot, err := discordgo.New("Bot " + token)

	if err != nil {
		panic(fmt.Errorf("Error: Unable to create bot %s", err))
	}

	fmt.Println("Hello World!")
	fmt.Println("%s", bot)
}
