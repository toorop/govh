package govh

type Ip struct {
	OvhClient *ovhClient
}

func (i *Ip) List() (resp string, err error) {
	resp, err = i.OvhClient.Do("GET", "ip", "")
	return
}
