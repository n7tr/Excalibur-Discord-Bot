package start_end

import "github.com/bwmarrin/discordgo"

func BotLeave(s *discordgo.Session, event *discordgo.GuildCreate) {
	s.GuildLeave(event.ID)
}
