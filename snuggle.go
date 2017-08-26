package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// global
var dg *discordgo.Session

var (
	// Token : Discord API bot token
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()

}

func main() {

	// Create a new Discord session using the provided bot token (commandline arg).
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Message handler
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called every time a new message is created
// on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	cmdToken := "." // TODO: move to external config

	if m.Content[:1] == cmdToken && m.Content != "" {
		var msgArgs []string
		msgArgs = strings.Split(m.Content[1:], " ")
		if len(m.Content) > 0 {
			switch messageArgs[0] {
			case "butts":
				s.ChannelMessageSend(m.ChannelID, "wow that's hot")
				break
			}
		}
	}

}
