package ip

import (
//"net"
)

// Type OF IP
var IpType = []string{"cdn", "dedicated", "hosted_ssl", "loadBalancing", "mail", "pcc", "pci", "vpn", "vps", "xdsl"}

// IP
type Ip struct {
	// IP block (eg 91.121.78.23/32")
	IP   string
	Type string // IpType
}
