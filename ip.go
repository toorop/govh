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
	ipr, err := ip.New(client)

	//debug(cmd.Action)

	switch cmd.Action {
	// List
	case "list":
		ipType := "all"
		if len(cmd.Args) > 2 {
			ipType = cmd.Args[2]
		}
		ips, err := ipr.List(ipType)
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
		var t []byte
		t, err = ipr.LbList()
		resp = string(t)
		break

	case "fw":
		// ip fw ipBlock.IP list
		// ip fw x.x.x.x/y
		// Return IP V4 list of this block which is under firewall
		if len(cmd.Args) == 4 && cmd.Args[3] == "list" {
			block := ip.IpBlock{cmd.Args[2], ""}
			ips, err := ipr.FwListIpOfBlock(block)
			if err != nil {
				dieError(err)
			}
			for _, i := range ips {
				resp = fmt.Sprintf("%s%s\r\n", resp, i)
			}
			if len(resp) > 2 {
				resp = resp[0 : len(resp)-2]
			}
			break
		}

		// Add IP to firewall
		// cmd : ip fw ibBlock.IP ipV4 add
		if len(cmd.Args) == 5 && cmd.Args[4] == "add" {
			block := ip.IpBlock{cmd.Args[2], ""}
			if err = ipr.FwAddIp(block, cmd.Args[3]); err != nil {
				dieError(err)
			}
			dieOk(fmt.Sprintf("%s added to firewall", cmd.Args[3]))
		}

		// ip fw ipBlock.IP ipV4
		//
		//

		// ip fw ipVlock ipV4 enable
		//
		//

		// ip fw ipBlock.IP ipV4 disable
		//
		//

		// Remove IPv4 from firewall
		// cmd : ip fw ipBlock.IP ipV4 remove
		if len(cmd.Args) == 5 && cmd.Args[4] == "remove" {
			block := ip.IpBlock{cmd.Args[2], ""}
			if err = ipr.FwRemoveIp(block, cmd.Args[3]); err != nil {
				dieError(err)
			}
			dieOk(fmt.Sprintf("%s removed from firewall", cmd.Args[3]))
		}

		// ip fw ipBlock.IP ipV4 listRules all

		// ip fw ipBlock.IP ipV4 addRules rule (as Json)

		// ip fw ipBlock.IP ipV4 getRule sequence

		// ip fw ipBlock.IP ipV4 delRule sequence

		break
	default:
		err = errors.New(fmt.Sprintf("This action : '%s' is not valid or not implemented yet !", strings.Join(cmd.Args, " ")))
	}
	return

}
