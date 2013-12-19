package ip

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Toorop/govh"
)

type IpRessource struct {
	client *govh.OvhClient
}

func New(client *govh.OvhClient) (*IpRessource, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return &IpRessource{client: client}, nil
}

// List IP
func (i *IpRessource) List(ipType string) (ips []Ip, err error) {
	var uri, r string
	if ipType == "all" {
		uri = "ip"
		ipType = ""
	} else {
		uri = fmt.Sprintf("ip?type=%s", ipType)
	}
	r, err = i.client.Do("GET", uri, "")
	var ipl = []string{}
	err = json.Unmarshal([]byte(r), &ipl)
	if err == nil {
		for _, t := range ipl {
			ips = append(ips, Ip{t, ipType})
		}
	}
	return
}

// List IP load balancing
func (i *IpRessource) LbList() (resp string, err error) {
	resp, err = i.client.Do("GET", "ip/loadBalancing", "")
	return
}
