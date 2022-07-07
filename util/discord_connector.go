package util

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func DiscInit(token string, messages chan string) {
	discbot, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
	}

	discbot.AddHandler(messageCreate)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "@Ждун") {

	}
}
