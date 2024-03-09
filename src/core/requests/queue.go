package requests

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	queue      = make(chan string, 100)
	processing = false
)

func HandleQueue(s *discordgo.Session) {
	if processing {
		return
	}

	processing = true

	for {
		select {
		case guildID := <-queue:
			go func(guildID string) {
				time.Sleep(2 * time.Second)
			}(guildID)
		default:

			processing = false
			return
		}
	}
}
