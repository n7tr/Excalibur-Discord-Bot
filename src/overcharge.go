package main

import (
	"Excalibur/core/requests"
	"os"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func LeaveEveryServer(s *discordgo.Session, m *discordgo.MessageCreate) {
	godotenv.Load()
	BOT_OWNER_ID := os.Getenv("BOT_OWNER_ID")

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content != ".overcharge" {
		return
	}

	if m.Content == ".overcharge" && m.Author.ID == BOT_OWNER_ID {
		s.ChannelMessageDelete(m.ChannelID, m.ID)

		guilds := s.State.Guilds
		smoothed := requests.Smooth(guilds)

		for _, ch := range smoothed {
			wg := new(sync.WaitGroup)
			wg.Add(len(ch))
			for _, guild := range ch {
				go func(guild *discordgo.Guild) {
					defer wg.Done()
					s.GuildLeave(guild.ID)
				}(guild)
			}
			wg.Wait()
		}
	} else {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
	}
}
