package removing

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func EmojiDelete(s *discordgo.Session, event *discordgo.GuildCreate) {
	emojis, err := s.GuildEmojis(event.ID)

	if err != nil {
		log.Println(err)
		return
	}

	for _, emoji := range emojis {
		err := s.GuildEmojiDelete(event.ID, emoji.ID)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
