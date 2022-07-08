package util

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var messagesChan chan string
var prefix string

func DiscInit(token string, messages chan string, pr string) {
	messagesChan = messages
	prefix = pr

	log.Println("Watcher prefix is", prefix)

	discbot, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Println("Error creating Discord session: ", err)
	}
	log.Println("Authorised in Discord")

	discbot.AddHandler(messageCreate)

	discbot.Identify.Intents = discordgo.IntentsGuildMessages

	err = discbot.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	fmt.Println("Service is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	err = discbot.Close()
	if err != nil {
		return
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	log.Println("Входящее сообщение из Discord: ", m.Content)

	if strings.HasPrefix(m.Content, prefix) {
		messagesChan <- m.Content
	}
}
