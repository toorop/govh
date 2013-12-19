package main

import (
	"errors"
	"fmt"
	"github.com/Toorop/govh"
	"github.com/Toorop/govh/ip"
	"strings"
)

func ipHandler(cmd *Cmd) (resp string, err error) {
	// New govh client
	client := govh.NewClient(OVH_APP_KEY, OVH_APP_SECRET, ck)
	// New ip ressource
	ip, err := ip.New(client)

	switch cmd.Action {
	// List
	case "list":
		ipType := "all"
		if len(cmd.Args) > 2 {
			ipType = cmd.Args[2]
		}
		ips, err := ip.List(ipType)
		if err != nil {
			dieError(err)
		}
		for _, i := range ips {
			resp = fmt.Sprintf("%s%s\r\n", resp, i.IP)
		}
		if len(resp) > 2 {
			resp = resp[0 : len(resp)-2]
		}
		break
	case "lb":
		if len(cmd.Args) < 3 {
			dieError("\"ip lb\" needs an argument see doc at https://github.com/Toorop/govh/blob/master/cli/README.md")
		}
		resp, err = ip.LbList()
		break
	default:
		err = errors.New(fmt.Sprintf("This action : '%s' is not valid or not implemented yet !", strings.Join(cmd.Args, " ")))
	}
	return

}
