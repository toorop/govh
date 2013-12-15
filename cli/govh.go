package main

import (
	"flag"
	"fmt"
	"github.com/Toorop/govh/govh"
	"github.com/Toorop/govh/ovh"
	"github.com/toqueteos/webbrowser"
	"os"
)

var (
	keyring ovh.Keyring
	cmd     govh.Cmd
)

func init() {

	keyring.AppKey = "SECRET"
	keyring.AppSecret = "SECRET"
	flag.StringVar(&keyring.ConsumerKey, "ck", "", "Consumer Key")
	flag.Parse()

	if len(flag.Args()) > 1 {
		cmd = govh.Cmd{
			Domain: flag.Arg(0),
			Action: flag.Arg(1),
		}
	}

	// if No ConsumerKey, request one
	if len(keyring.ConsumerKey) == 0 {
		fmt.Println("\r\nNo consummer key found. We will request one ....\r\n")
		ck, link, err := ovh.AuthGetConsumerKey(keyring)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Your consumer key is %s\r\n\r\n", ck)
		fmt.Println("You need to validate it :")
		fmt.Printf("\t- If you have a browser available on this machine it will open to the validation page.\n\t- If not copy and past the link below in a browser to validate your key :\r\n%s\r\n", link)
		webbrowser.Open(link)
		fmt.Printf("\n\nWhen it will be done relaunch govh with parameter -ck:\n\tgovh [ACTION] -ck %s\n\n\n", ck)
		os.Exit(0)
	}
}

// Main
func main() {
	switch cmd.Domain {
	case "sms":
		govh.SmsHandler(cmd)
		break
	case "help":
		fmt.Println("HELP")
		break
	default:
		fmt.Println("This domain is not valid or not implemented yet !\r\n\tgovh help\r\nfor... help.")
	}

}
