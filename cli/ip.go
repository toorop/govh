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
		resp, err = ip.List()
		break
	default:
		err = errors.New(fmt.Sprintf("This action : '%s' is not valid or not implemented yet !", strings.Join(cmd.Args, " ")))
	}
	return

}
