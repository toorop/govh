package ip

import (
//"time"
)

// Type OF IP
var IpType = []string{"cdn", "dedicated", "hosted_ssl", "loadBalancing", "mail", "pcc", "pci", "vpn", "vps", "xdsl"}

// IP
type IpBlock struct {
	// IP (eg 91.121.78.23/32")
	IP   string
	Type string // IpType
}

// ip.FirewallIp
type IpFirewallIp struct {
	IpOnFirewall string `json:"ipOnFirewall"`
	Enabled      bool   `json:"enabled"`
	State        string `json:"state"`
}

// Firewall rules

// destinationPort
type DestinationPort struct {
	From int `json:"from"`
	To   int `json:"to"`
}

// sourcePort
type SourcePort struct {
	From int `json:"from"`
	To   int `json:"to"`
}

// tcpOption
type TcpOption struct {
	Urg         bool `json:"urg"`
	Psh         bool `json:"psh"`
	Ack         bool `json:"ack"`
	Established bool `json:"established"`
	Syn         bool `json:"syn"`
	Fin         bool `json:"fin"`
	Rst         bool `json:"rst"`
}

// udpOption
type UdpOption struct {
	Fragments bool `json:"fragment"`
}

// Post
type FirewallRule2Add struct {
	Action          string          `json:"action"`
	DestinationPort DestinationPort `json:"destinationPort,omitempty"`
	Protocol        string          `json:"protocol"`
	Sequence        string          `json:"sequence"`
	Source          string          `json:"source,omitempty"`
	SourcePort      SourcePort      `json:"sourcePort,omitempty"`
	TcpOption       TcpOption       `json:"tcpOption,omitempty"`
	UdpOption       UdpOption       `json:"udpOption,omitempty"`
}

// Reply
type FirewallRule struct {
	Protocol        string   `json:"protocol"`
	Source          string   `json:"source"`
	DestinationPort string   `json:"destinationPort"`
	Sequence        int      `json:"sequence"`
	Options         []string `json:"options"`
	Destination     string   `json:"destination"`
	Rule            string   `json:"rule"`
	SourcePort      string   `json:"sourcePort"`
	State           string   `json:"state"`
	CreationDate    string   `json:"creationDate"`
	Action          string   `json:"action"`
}

//
//// SPAM
//

// SpamIp
type SpamIp struct {
	Time       uint32 `json:"time"`       // Time (in seconds) while the IP will be blocked
	Date       string `json:"date"`       // Last date the ip was blocked
	IpSpamming string `json:"ipSpamming"` // IP address which is sending spam
	State      string `json:"state"`      // Current state of the ip. blockedForSpam | unblocked | unblocking
}

// SpamTarget
// Spam's target information
type SpamTarget struct {
	DestinationIp string `json:"destinationIp"` // IP address of the target
	MessageId     string `json:"messageId"`     // The message-id of the email
	Date          int64  `json:"date"`          // Timestamp when the email was sent
	//Spamcause     string `json:"spamcause"`     // Detailled spam cause
	Spamscore uint `json:"spamscore"` // Spam score of the email
}

// SpamStats
// Spam statistics about an IP address
type SpamStats struct {
	Timestamp        int64 `json:"timestamp"` // Time when the IP address was blocked
	DetectedSpams    []SpamTarget
	AverageSpamScore int `json:"averageSpamscore"` // Average spam score.
	Total            int `json:"total"`            // Number of emails sent
	NumberOfSpams    int `json:"numberOfSpams"`    //Number of spams sent
}
