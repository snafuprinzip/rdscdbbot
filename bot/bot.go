package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/snafuprinzip/renscbot/config"
)

var (
	BotID string
	goBot *discordgo.Session
)

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID
	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("SC Website Collector Bot is running...")
	return
}

func messageHandler(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if strings.HasPrefix(msg.Content, config.CmdPrefix) {
		// Drop our own messages
		if msg.Author.ID == BotID {
			return
		}

		if msg.Content == "!ping" {
			_, _ = s.ChannelMessageSend(msg.ChannelID, "pong")
		}
	}
	fmt.Println(msg.Content)
	return
}
