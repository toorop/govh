package ip

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Toorop/govh"
	"net/url"
	//"os"
	"strings"
	"time"
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
	if err != nil {
		err = errors.New(fmt.Sprintf("%s - %s", err, t))
		return
	}

	var ipl = []string{}
	if err = json.Unmarshal(t, &ipl); err == nil {
		for _, i := range ipl {
			ips = append(ips, IpBlock{i, ipType})
		}
	}
	return
}

//
//// LOADBALANCING
//

// List IP load balancing
func (r *IpRessource) LbList() (resp []byte, err error) {
	resp, err = r.client.Do("GET", "ip/loadBalancing", "")
	return
}

//
//// FIREWALL
//

// List IP of block IP under firewall protection
func (r *IpRessource) FwListIpOfBlock(block IpBlock) (ips []string, err error) {
	uri := fmt.Sprintf("ip/%s/firewall", url.QueryEscape(block.IP))
	t, err := r.client.Do("GET", uri, "")
	if err != nil {
		err = errors.New(fmt.Sprintf("%s - %s", err, t))
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
	t, err := r.client.Do("POST", uri, string(payload))
	if err != nil {
		err = errors.New(fmt.Sprintf("%s - %s", err, t))
	}
	return err
}

// Remove IP from firewall
func (r *IpRessource) FwRemoveIp(block IpBlock, ipv4 string) error {
	uri := fmt.Sprintf("ip/%s/firewall/%s", url.QueryEscape(block.IP), ipv4)
	t, err := r.client.Do("DELETE", uri, "")
	if err != nil {
		err = errors.New(fmt.Sprintf("%s - %s", err, t))
	}
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
	t, err := r.client.Do("PUT", uri, string(payload))
	if err != nil {
		err = errors.New(fmt.Sprintf("%s - %s", err, t))
	}
	return err
}

// Get properties about an IP firewalled
func (r *IpRessource) FwGetIpProperties(block IpBlock, ipv4 string) (i IpFirewallIp, err error) {
	uri := fmt.Sprintf("ip/%s/firewall/%s", url.QueryEscape(block.IP), ipv4)
	t, err := r.client.Do("GET", uri, "")
	if err != nil {
		return i, errors.New(fmt.Sprintf("%s - %s", err, t))
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
	t, err := r.client.Do("POST", uri, string(payload))
	if err != nil {
		err = errors.New(fmt.Sprintf("%s - %s", err, t))
	}
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
		return sequences, errors.New(fmt.Sprintf("%s - %s", err, t))
	}
	err = json.Unmarshal(t, &sequences)
	return
}

// Get rule
func (r *IpRessource) FwGetRule(block IpBlock, ipv4 string, sequence int) (rule FirewallRule, err error) {
	uri := fmt.Sprintf("ip/%s/firewall/%s/rules/%d", url.QueryEscape(block.IP), ipv4, sequence)
	t, err := r.client.Do("GET", uri, "")
	if err != nil {
		return rule, errors.New(fmt.Sprintf("%s - %s", err, t))
	}
	err = json.Unmarshal(t, &rule)
	return
}

// Remove a rule
func (r *IpRessource) FwRemoveRule(block IpBlock, ipv4 string, sequence int) error {
	uri := fmt.Sprintf("ip/%s/firewall/%s/rules/%d", url.QueryEscape(block.IP), ipv4, sequence)
	t, err := r.client.Do("DELETE", uri, "")
	if err != nil {
		err = errors.New(fmt.Sprintf("%s - %s ", err, t))
	}
	return err
}

//
// SPAM
//

// Get Spamming IP
// state :
// 		* blockedForSpam : currently blocked
// 		* unblocking : in the way to be unblocked (or not)
// 		* unblocked : blocked in the past
//
func (r *IpRessource) SpamGetSpammingIps(block IpBlock, state string) (ips []string, err error) {
	uri := fmt.Sprintf("ip/%s/spam", url.QueryEscape(block.IP))
	if state == "blockedForSpam" || state == "unblocking" || state == "unblocked" {
		uri = fmt.Sprintf("%s?state=%s", uri, state)
	}
	t, err := r.client.Do("GET", uri, "")
	if err != nil {
		return ips, errors.New(fmt.Sprintf("%s - %s", err, t))
	}
	err = json.Unmarshal(t, &ips)
	return
}

// Get detailed info about a spamming IP
func (r *IpRessource) SpamGetSpamIp(block IpBlock, ipv4 string) (spamIp *SpamIp, err error) {
	uri := fmt.Sprintf("ip/%s/spam/%s", url.QueryEscape(block.IP), ipv4)
	t, err := r.client.Do("GET", uri, "")
	if err != nil {
		return spamIp, errors.New(fmt.Sprintf("%s - %s ", err, t))
	}
	err = json.Unmarshal(t, &spamIp)
	return
}

// Get stats about a spamming IP
func (r *IpRessource) SpamGetIpStats(block IpBlock, ipv4 string, from time.Time, to time.Time) (spamStats *SpamStats, err error) {
	uri := fmt.Sprintf("ip/%s/spam/%s/stats?from=%s&to=%s", url.QueryEscape(block.IP), ipv4, url.QueryEscape(from.Format(time.RFC3339)), url.QueryEscape(to.Format(time.RFC3339)))
	t, err := r.client.Do("GET", uri, "")
	if err != nil {
		return spamStats, errors.New(fmt.Sprintf("%s - %s ", err, t))
	}
	if len(t) > 2 {
		err = json.Unmarshal(t[1:len(t)-1], &spamStats)
	}
	return
}

// Unblock a spamming IP
func (r *IpRessource) SpamUnblockSpamIp(block IpBlock, ipv4 string) error {
	uri := fmt.Sprintf("ip/%s/spam/%s/unblock", url.QueryEscape(block.IP), ipv4)
	t, err := r.client.Do("POST", uri, "")
	if err != nil {
		err = errors.New(fmt.Sprintf("%s - %s", err, t))
	}
	return err
}

// Helper
// Return IPs which are currently blocked for spam
func (r *IpRessource) GetBlockedForSpam() (ips []string, err error) {
	ipBlocks, err := r.List("all")
	if err != nil {
		return
	}
	for _, ipb := range ipBlocks {
		if len(strings.Split(ipb.IP, ":")) > 1 {
			continue
		}
		ipsBlocked, err := r.SpamGetSpammingIps(ipb, "blockedForSpam")
		if err != nil {
			// Not all IP are concerned by spamming status, if not found continue
			if strings.HasPrefix(err.Error(), "404 This service does not exist") {
				continue
			}
			return ips, err
		}
		if len(ipsBlocked) > 0 {
			for _, i := range ipsBlocked {
				ips = append(ips, i)
			}
		}
	}
	return
}
