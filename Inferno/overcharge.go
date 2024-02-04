package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func LeaveEveryServer(s *discordgo.Session, m *discordgo.MessageCreate) {
	godotenv.Load()
	if m.Author.ID == s.State.User.ID {
		return
	}

	BOT_OWNER_ID := os.Getenv("BOT_OWNER_ID")
	if m.Content == ".overcharge_leave" {
		if m.Author.ID == BOT_OWNER_ID {
			s.ChannelMessageDelete(m.ChannelID, m.ID)

			guilds := s.State.Guilds
			for _, guild := range guilds {
				s.GuildLeave(guild.ID)
			}
		} else {
			s.ChannelMessageDelete(m.ChannelID, m.ID)
			s.ChannelMessageDelete(m.ChannelID, "Can't leave from all servers.")
		}
	}
}
