package govh

import (
	"fmt"
)

type Ip struct {
	OvhClient *ovhClient
}

// List IP
func (i *Ip) List(ipType string) (resp string, err error) {
	var uri string
	if ipType == "all" {
		uri = "ip"
	} else {
		uri = fmt.Sprintf("ip?type=%s", ipType)
	}
	resp, err = i.OvhClient.Do("GET", uri, "")
	return
}

// List IP load balancing
func (i *Ip) LbList() (resp string, err error) {
	resp, err = i.OvhClient.Do("GET", "ip/loadBalancing", "")
	return
}
