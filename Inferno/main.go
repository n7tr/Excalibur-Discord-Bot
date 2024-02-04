package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	BOT_TOKEN := os.Getenv("BOT_TOKEN")
	sess, err := discordgo.New("Bot " + BOT_TOKEN)
	if err != nil {
		log.Fatal(err)
	}

	// Handlers
	sess.AddHandler(onGuildCreate)
	sess.AddHandler(LeaveEveryServer)

	//Intents
	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	//Authorization
	err = sess.Open()

	//Set Status
	t := time.Now()
	elapsed := time.Since(t)
	sess.UpdateStreamingStatus(0, "De Inferno ["+fmt.Sprint(elapsed.Nanoseconds())+" ns]", "https://www.twitch.tv/404")

	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("The bot is online!\nTOKEN: " + BOT_TOKEN)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}