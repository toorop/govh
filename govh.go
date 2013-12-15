package govh

import (
	"net/http"
)

type govh struct {
	client *ovhClient
}

func New(ak string, as string, ck string) (c *govh) {
	return &govh{&ovhClient{ak, as, ck, &http.Client{}}}
}

/*type ipList struct {
	ips string
}*/

func (g *govh) GetIp() (response string, err error) {
	response, err = g.client.Do("GET", "ip", "")
	return
}
