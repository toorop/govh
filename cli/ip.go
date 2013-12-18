package main

import (
	"errors"
	"fmt"
	"github.com/Toorop/govh"
	"strings"
)

func ipHandler(cmd *Cmd) (resp string, err error) {
	ip := govh.Ip{govh.NewClient(OVH_APP_KEY, OVH_APP_SECRET, ck)}

	switch cmd.Action {
	// List
	case "list":
		ipType := "all"
		if len(cmd.Args) > 2 {
			ipType = cmd.Args[2]
		}
		resp, err = ip.List(ipType)
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
