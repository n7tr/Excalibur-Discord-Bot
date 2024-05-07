package bypass

import (
	"github.com/bwmarrin/discordgo"
)

func GetBotNicks(s *discordgo.Session, guildID string) ([]string, error) {
	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		return nil, err
	}

	botNicknames := []string{}

	for _, member := range members {
		if member.User.Bot {
			botNickname := member.Nick
			if botNickname == "" {
				botNickname = member.User.Username
			}

			botNicknames = append(botNicknames, botNickname)
		}
	}

	return botNicknames, nil
}
