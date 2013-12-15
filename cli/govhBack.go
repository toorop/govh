package main

import (
	"flag"
	"fmt"
	"github.com/Toorop/govh/ovh"
	"github.com/toqueteos/webbrowser"
)

var keyring ovh.Keyring

func init() {
	keyring.AppKey = "SECRET"
	keyring.AppSecret = "SECRET"
	flag.StringVar(&keyring.ConsumerKey, "ck", "", "Consumer Key")
	flag.Parse()
}

func main() {

	// if No ConsumerKey, request one
	if len(keyring.ConsumerKey) == 0 {
		fmt.Println("No consummer key found. We will request one ....")
		ck, link, err := ovh.GetConsumerKeyAndValidationURL(keyring)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Your consumer key is %s\r\n", ck)
		fmt.Println("You need to validate it.")
		fmt.Printf("If you have a browser available on this machine it will open to the validation page. If not copy and past the link below in a browser to validate your key :\r\n%s\r\n", link)
		webbrowser.Open(link)
		//fmt.Println(err, ck, link)
	}
}
