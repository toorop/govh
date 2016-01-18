package ip

import (
	"github.com/toorop/govh"
)

// Type OF IP
var IPType = []string{"cdn", "dedicated", "hosted_ssl", "loadBalancing", "mail", "pcc", "pci", "vpn", "vps", "xdsl"}

/*
// DateTime represents date as returned by OVH
type DateTime struct {
	time.Time
}

func (dt *DateTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02T15:04:05+02:00", s)
	if err != nil {
		return err
	}
	dt.Time = t
	return nil
}

func (dt DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal((*time.Time)(&dt.Time).Format("2006-01-02T15:04:05+02:00"))
}

// DateTime2 represents an other date as returned by OVH
// don't ask me why but OVH use differents TZ for their datetime
type DateTime2 struct {
	time.Time
}

func (dt *DateTime2) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02T15:04:05+01:00", s)
	if err != nil {
		return err
	}
	dt.Time = t
	return nil
}

func (dt DateTime2) MarshalJSON() ([]byte, error) {
	return json.Marshal((*time.Time)(&dt.Time).Format("2006-01-02T15:04:05+01:00"))
}
*/

// IP
type IPBlock struct {
	IP   string
	Type string // IpType
}

// IpProperties represents properties of an IP
type IPProperties struct {
	Ip          string `json:"ip,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	RoutedTo    struct {
		ServiceName string `json:"serviceName,omitempty"`
	} `json:"routedTo,omitempty"`
}

// IpUpdatableProperties represents updatable properties of an IP
type IpUpdatableProperties struct {
	Description string `json:"description,omitempty"`
}

//FirewalledIp represents an IP on the Firewall
type FirewalledIp struct {
	Ip      string `json:"ipOnFirewall"`
	Enabled bool   `json:"enabled"`
	State   string `json:"state"`
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

// fwTcpOption represents TCP option for a firewall rule
type FwTcpOption struct {
	Fragments bool   `json:"fragments,omitempty"`
	Option    string `json:"option,omitempty"`
}

// FwFirewallRule2Add
type FwRule2Add struct {
	Action    string       `json:"action"`
	ToPort    int          `json:"destinationPort,omitempty"`
	Protocol  string       `json:"protocol"`
	Sequence  int          `json:"sequence"`
	FromIp    string       `json:"source,omitempty"`
	FromPort  int          `json:"sourcePort,omitempty"`
	TcpOption *FwTcpOption `json:"tcpOption,omitempty"`
}

// Reply
type FirewallRule struct {
	Protocol     string        `json:"protocol"`
	FromIp       string        `json:"source"`
	ToPort       string        `json:"destinationPort"`
	Sequence     int           `json:"sequence"`
	TcpOption    string        `json:"tcpOption"`
	ToIp         string        `json:"destination"`
	Rule         string        `json:"rule"`
	FromPort     string        `json:"sourcePort"`
	State        string        `json:"state"`
	CreationDate govh.DateTime `json:"creationDate"`
	Action       string        `json:"action"`
	Fragments    bool          `json:"fragments"`
}

//
//// SPAM
//

// SpamIp
type SpamIP struct {
	Time       uint32        `json:"time"`       // Time (in seconds) while the IP will be blocked
	Date       govh.DateTime `json:"date"`       // Last date the ip was blocked
	IpSpamming string        `json:"ipSpamming"` // IP address which is sending spam
	State      string        `json:"state"`      // Current state of the ip. blockedForSpam | unblocked | unblocking
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
