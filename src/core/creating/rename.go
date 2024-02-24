package creating

import (
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func GuildRename(s *discordgo.Session, event *discordgo.GuildCreate) {

	godotenv.Load()
	server_name := os.Getenv("SERVER_NAME")

	avatarData, _ := ioutil.ReadFile("avatar.webp")
	avatarBase64 := base64.StdEncoding.EncodeToString(avatarData)

	guildID := event.ID
	s.GuildEdit(guildID, &discordgo.GuildParams{
		Name: server_name,
		Icon: "data:image/png;base64," + avatarBase64,
	})
}
