<img src="https://i.imgur.com/aIcOxU8.jpeg">

# Excalibur Discord Bot
A powerful Discord nuke bot written in Go

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
    BOT_OWNER_ID: your id

    MASS_BAN: true or false

    WEBHOOK_URL: Webhook's URL
    AVATAR_URL: avatar url for webhook

    PREFERRED_LOCALE: check list of <a href="https://discord.com/developers/docs/reference#locales">locales</a>

    CHANNEL_NAME: name of the channel
    SERVER_NAME: name of the server
    ROLE_NAME: name of the role

    EMBED_TITLE: Embed's title
    EMBED_DESCRIPTION: Embed's description

}
</pre>
All variables have a string data type. Only the MASS_BAN variable has two possible values - true and false. Write them with a lowercase letter.

# Code Structure
All bot's functions are in core folder

# Bypass Anti Nuke bots
Yes, Excalibur can bypass anti nuke bots like Security, Wick and other. Bot checks members for anti-nuke bots. If they are there, then the bot runs a function that bypasses the protection.

# main.go & auto.go
main.go - starts the bot and onGuildCreate handler from auto.go
<hr>
auto.go - runs bot functions from core folder
<hr>

# sendhttp.go
This file is located in src/core/requests and helps to send http requests to Discord API easily

# smooth.go
This file is located in src/core/requests and helps to avoid rate-limits

# queue.go
This function is responsible for creating a nuke queue

# Nuking process
This bot nukes the server when you add it. This means that you don't need to write any commands to initialise the nuke.

# Installation guide
<pre>
	1. Clone or download the repository source code
	2. Install golang
	3. Go to src folder
	4. Change values in .env
	5. Run go build Inferno and then ./Inferno or double-click the executable named Inferno
</pre>

# Where to host?
We recommend you to use <a href="https://fl0.com">fl0.com</a>, <a href="https://back4app.com">back4app.com</a>, <a href="https://koyeb.com">koyeb.com</a> and <a href="https://render.com">render.com</a>. They're free and there you can host Dynamic and other discord bots. More information about other hostings are <a href="https://github.com/DmitryScaletta/free-heroku-alternatives">here</a>

# Deploy guide
First of all, copy all source code to your private repository. Then create an account on <a href="https://railway.app">railway.app</a> via github. Use Dockerfile for quick deployment. <a href="https://railway.app">Railway.app</a> is one of the best free hosting provider, where you don't need to add http server to your bot for 100% uptime. 

# Dockerfile example
<pre>
# For deployment on railway.app
FROM golang:latest

WORKDIR /

COPY . .

RUN go build Inferno

CMD [ "./Inferno" ]
</pre>

<pre>
# For deployment on render.com and others
FROM golang:latest

WORKDIR /

COPY . .

RUN go build Inferno

EXPOSE 8080

CMD [ "./Inferno" ]
</pre>

If you want to deploy your fork on render.com, add code snippet bellow to main.go
<pre>
// imports
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

//starts http server
func main() {
	go func() {
		http.HandleFunc("/", getRoot)
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Excalibur is at render.com now.. 🚀\n")
}
</pre>
