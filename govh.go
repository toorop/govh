package govh

import (
	"net/http"
)

type govh struct {
	client *OvhClient
}

func New(ak string, as string, ck string) (c *govh) {
	return &govh{&OvhClient{ak, as, ck, &http.Client{}}}
}

/*type ipList struct {
	ips string
}*/
