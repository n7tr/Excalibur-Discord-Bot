<img src="https://media.discordapp.net/attachments/1109426745012142081/1183350870079971449/IMG_9567.png?ex=65880461&is=65758f61&hm=44475ba2ed4ff3a5e0a01e530660de151736d11b3038d632fbb672e66a0bd8cf&=&format=webp&quality=lossless&width=1202&height=657">

# Inferno-Discord-Bot
A powerful Discord nuke bot written on Go
<hr>

# .env file
There's many variables such as 

<pre>
{
    BOT_TOKEN: bot's token
    
    WEBHOOK_ID: Webhook's ID
    WEBHOOK_TOKEN: Webhook's Token
    AVATAR_URL: avatar url for webhook

    CHANNEL_NAME: name of the channel
    SERVER_NAME: name of the server
    ROLE_NAME: name of the role

    EMBED_TITLE: Embed's title
    EMBED_DESCRIPTION: Embed's description

}
</pre>
<p>/// All values are have string data type</p>

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
