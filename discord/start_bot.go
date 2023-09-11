package discord

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func StartBot() {
	// TODO: move out to config file
	dg, err := discordgo.New("Bot " + "INSERT AUTH TOKEN HERE")
	if err != nil {
		log.Fatal(err)
	}

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	guildID := "GUILD ID GOES HERE"
	RegisterCommands(dg, &guildID)

	defer dg.Close()
	log.Print("Discord bot started")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
