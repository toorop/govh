package main

import (
	"flag"
	"fmt"
	"github.com/Toorop/govh"
	//"github.com/toqueteos/webbrowser"
	//"github.com/wsxiaoys/terminal"
	//"github.com/wsxiaoys/terminal/color"
	"encoding/json"
	"os"
)

const (
	OVH_APP_KEY    = "SECRET"
	OVH_APP_SECRET = "SECRET"
)

var (
	ck           string // consumer key
	outputFormat string
	keyring      govh.Keyring
	cmd          Cmd
)

func init() {

	// Check and load config
	// export OVH_CONSUMER_KEY="WYAUsR31Z3dT9Y5f0arTHeZwpFRdcnz2"

	keyring.AppKey = OVH_APP_KEY
	keyring.AppSecret = OVH_APP_SECRET
	//flag.StringVar(&keyring.ConsumerKey, "ck", "", "Consumer Key")

	flag.StringVar(&ck, "ck", "", "Consumer Key")
	flag.StringVar(&outputFormat, "of", "JSON", "Output format")
	flag.Parse()

	// WYAUsR31Z3dT9Y5f0arTHeZwpFRdcnz2
	ck = os.Getenv("OVH_CONSUMER_KEY")

	// No CK
	if len(ck) == 0 {
		dieError("No consumer key found")
	}

	//fmt.Println(ck)

	/*if len(flag.Args()) > 1 {
		cmd = Cmd{
			Domain: flag.Arg(0),
			Action: flag.Arg(1),
		}
	}*/

	/*// Is there a consumer key ?
	if len(keyring.ConsumerKey) == 0 {
		// Check ENV
		keyring.ConsumerKey = os.Getenv("OVH_CONSUMER_KEY")
	}

	// if No ConsumerKey, request one
	if len(keyring.ConsumerKey) == 0 {
		fmt.Println("\r\nNo consummer key found. We will request one ....\r\n")
		ck, link, err := govh.AuthGetConsumerKey(keyring)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Your consumer key is %s\r\n\r\n", ck)
		fmt.Println("You need to validate it :")
		fmt.Printf("\t- If you have a browser available on this machine it will open to the validation page.\n\t- If not copy and past the link below in a browser to validate your key :\r\n%s\r\n", link)
		webbrowser.Open(link)
		fmt.Printf("\n\nWhen it will be done relaunch govh with parameter -ck:\n\tgovh [ACTION] -ck %s\n\n\n", ck)
		os.Exit(0)
	}*/
}

// Main
func main() {
	govh := govh.New(OVH_APP_KEY, OVH_APP_SECRET, ck)
	resp, _ := govh.GetIp()
	var ips []string
	_ = json.Unmarshal([]byte(resp), &ips)
	for _, ip := range ips {
		fmt.Println(ip)
	}
	//fmt.Println(resp)

	/*ovh := govh.NewClient(OVH_APP_KEY, OVH_APP_SECRET, ck)
	resp, _ := ovh.Do("GET", "ip", "")

	fmt.Println(resp)*/

	/*terminal.Stdout.Color("y").
	Print("Hello world").Nl().
	Reset().
	Colorf("@{kW}Hello world\n")
	*/
	//color.Println("@rHello worldiz")

	/*
		switch cmd.Domain {
		//case "sms":
		//	govh.SmsHandler(cmd)
		//	break
		case "help":
			fmt.Println("HELP")
			break
		default:
			fmt.Println("This domain is not valid or not implemented yet !\r\n\tgovh help\r\nfor... help.")
		}*/

}
