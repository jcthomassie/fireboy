package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
)

// Variables used for command line parameters
var (
	Token      string
	TestStatement string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot token")
	flag.StringVar(&TestStatement, "s", "", "Should be hello world")
	flag.Parse()

	if Token == "" || TestStatement == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	_, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	fmt.Println(TestStatement)
}
