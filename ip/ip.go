package ip

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Toorop/govh"
	"net/url"
	//"os"
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
func (r *IpRessource) List(ipType string) (ips []IpBlock, err error) {
	var uri string
	if ipType == "all" {
		uri = "ip"
		ipType = ""
	} else {
		uri = fmt.Sprintf("ip?type=%s", ipType)
	}
	t, err := r.client.Do("GET", uri, "")
	var ipl = []string{}
	if err = json.Unmarshal(t, &ipl); err == nil {
		for _, i := range ipl {
			ips = append(ips, IpBlock{i, ipType})
		}
	}
	return
}

//// LOADBALANCING

// List IP load balancing
func (r *IpRessource) LbList() (resp []byte, err error) {
	resp, err = r.client.Do("GET", "ip/loadBalancing", "")
	return
}

// FIREWALL

// List IP of block IP under firewall protection
func (r *IpRessource) FwListIpOfBlock(block IpBlock) (ips []string, err error) {
	uri := fmt.Sprintf("ip/%s/firewall", url.QueryEscape(block.IP))
	t, err := r.client.Do("GET", uri, "")
	if err != nil {
		return
	}
	err = json.Unmarshal(t, &ips)
	return
}

// Add IP to firewall
func (r *IpRessource) FwAddIp(block IpBlock, ipv4 string) error {
	uri := fmt.Sprintf("ip/%s/firewall", url.QueryEscape(block.IP))
	type p struct {
		IpOnFirewall string `json:"ipOnFirewall"`
	}
	payload, err := json.Marshal(p{ipv4})
	if err != nil {
		return err
	}
	_, err = r.client.Do("POST", uri, string(payload))
	return err
}

// Remove IP from firewall
func (r *IpRessource) FwRemoveIp(block IpBlock, ipv4 string) error {
	uri := fmt.Sprintf("ip/%s/firewall/%s", url.QueryEscape(block.IP), ipv4)
	_, err := r.client.Do("DELETE", uri, "")
	return err
}

// Enable / disable firewall for IP ipv4 of IpBlock block
func (r *IpRessource) FwSetFirewallEnable(block IpBlock, ipv4 string, enabled bool) error {
	uri := fmt.Sprintf("ip/%s/firewall/%s", url.QueryEscape(block.IP), ipv4)
	type p struct {
		Enabled bool `json:"enabled"`
	}
	payload, err := json.Marshal(p{enabled})
	if err != nil {
		return err
	}
	_, err = r.client.Do("PUT", uri, string(payload))
	return err
}

// Get properties about an IP firewalled
func (r *IpRessource) FwGetIpProperties(block IpBlock, ipv4 string) (i IpFirewallIp, err error) {
	uri := fmt.Sprintf("ip/%s/firewall/%s", url.QueryEscape(block.IP), ipv4)
	t, err := r.client.Do("GET", uri, "")
	if err != nil {
		return i, err
	}
	err = json.Unmarshal(t, &i)
	return
}

// Add a rule
func (r *IpRessource) FwAddRule(block IpBlock, ipv4 string, rule FirewallRule2Add) error {
	uri := fmt.Sprintf("ip/%s/firewall/%s/rules", url.QueryEscape(block.IP), ipv4)
	if rule.DestinationPort.From == rule.DestinationPort.To && rule.DestinationPort.To == 0 {
		rule.DestinationPort.From = 0
		rule.DestinationPort.To = 65535
	}
	if rule.SourcePort.From == rule.SourcePort.To && rule.SourcePort.To == 0 {
		rule.SourcePort.From = 0
		rule.SourcePort.To = 65535
	}

	payload, err := json.Marshal(rule)
	_, err = r.client.Do("POST", uri, string(payload))
	return err
}

// List firewall rules sequences
func (r *IpRessource) FwGetRulesSequences(block IpBlock, ipv4 string, state string) (sequences []int, err error) {
	uri := fmt.Sprintf("ip/%s/firewall/%s/rules", url.QueryEscape(block.IP), ipv4)
	if state == "creationPending" || state == "ok" || state == "removalPending" {
		uri = fmt.Sprintf("%s?state=%s", state)
	}
	t, err := r.client.Do("GET", uri, "")
	if err != nil {
		return sequences, err
	}
	err = json.Unmarshal(t, &sequences)
	return
}

// Get rule
func (r *IpRessource) FwGetRule(block IpBlock, ipv4 string, sequence int) (rule FirewallRule, err error) {
	uri := fmt.Sprintf("ip/%s/firewall/%s/rules/%d", url.QueryEscape(block.IP), ipv4, sequence)
	t, err := r.client.Do("GET", uri, "")
	if err != nil {
		return rule, err
	}
	err = json.Unmarshal(t, &rule)
	return
}

// Remove a rule
func (r *IpRessource) FwRemoveRule(block IpBlock, ipv4 string, sequence int) error {
	uri := fmt.Sprintf("ip/%s/firewall/%s/rules/%d", url.QueryEscape(block.IP), ipv4, sequence)
	_, err := r.client.Do("DELETE", uri, "")
	return err
}
