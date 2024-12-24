package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: %s <path to .env>", os.Args[0])
		return
	}

	err := godotenv.Load(os.Args[1])
	token := os.Getenv("DISCORD_TOKEN")
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.Contains(strings.ToLower(m.Content), "john") {
		return
	}

	var responses = []string{
		"Better be seen at oxford than caught at john's",
		"rather go to oxford than st johns tbh",
		"watch your mouth before you get sent to St Johns",
		"johns is basically a prison",
		"when do we demolish st johns again?",
		"they say st john's was built to make the rest of cambridge look better.",
		"the best part of visiting st johns is when you leave",
		"mfw when someone mentions johns",
		"No Cantabrigian wind howleth fiercer than mine antipathy for the folly that is St John's",
	}

	var response = responses[rand.Intn(len(responses))]

	s.ChannelMessageSendReply(m.ChannelID, response, m.Reference())
}
