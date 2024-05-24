package start_end

import (
	"Excalibur/core/requests"
	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

func Leave(s *discordgo.Session, event *discordgo.GuildCreate) {
	data := []byte{}
	jsonData, _ := json.Marshal(data)

	requests.Sendhttp("https://discord.com/api/v9/users/@me/guilds/"+event.ID, "DELETE", jsonData)
}
