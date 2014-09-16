package ip

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Toorop/govh"
	"net/url"
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

// List return a slice of IpBlock
func (i *IpRessource) List(filterDesc, filterIp, filterRoutedTo, filterType string) (ips []IpBlock, err error) {
	uri := "ip"
	args := []string{}

	if len(filterDesc) > 0 {
		args = append(args, "description="+url.QueryEscape(filterDesc))
	}
	if len(filterIp) > 0 {
		args = append(args, "ip="+url.QueryEscape(filterIp))
	}
	if len(filterRoutedTo) > 0 {
		args = append(args, "routedTo.serviceName="+url.QueryEscape(filterRoutedTo))
	}
	if len(filterType) > 0 {
		args = append(args, "type="+url.QueryEscape(filterType))
	}

	if len(args) > 0 {
		uri = uri + "?" + strings.Join(args, "&")
	}
	r, err := i.client.Do("GET", uri, "")
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	var ipl = []string{}
	if err = json.Unmarshal(r.Body, &ipl); err == nil {
		for _, i := range ipl {
			ips = append(ips, IpBlock{i, filterType})
		}
	}
	return
}

// GetIpProperties return properties of an IP
func (i *IpRessource) GetIpProperties(ip string) (properties IpProperties, err error) {
	r, err := i.client.Do("GET", "ip/"+url.QueryEscape(ip), "")
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &properties)
	return
}

// UpdateProperties update IP properties
func (i *IpRessource) UpdateProperties(ip, desc string) error {
	payload, err := json.Marshal(IpUpdatableProperties{
		Description: desc,
	})
	if err != nil {
		return err
	}
	r, err := i.client.Do("PUT", "ip/"+url.QueryEscape(ip), string(payload))
	err = r.HandleErr(err, []int{200})
	return err
}

/*
//
//// LOADBALANCING
//

// List IP load balancing
func (r *IpRessource) LbList() (resp []byte, err error) {
	resp, err = r.client.Do("GET", "ip/loadBalancing", "")
	return
}*/

//
//// FIREWALL
//

// List IP of block IP under firewall protection
func (i *IpRessource) FwListIpOfBlock(block IpBlock) (ips []string, err error) {
	r, err := i.client.Do("GET", "ip/"+url.QueryEscape(block.IP)+"/firewall", "")
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &ips)
	return
}

// Add IP to firewall
func (i *IpRessource) FwAddIp(block IpBlock, ipv4 string) error {
	type p struct {
		IpOnFirewall string `json:"ipOnFirewall"`
	}
	payload, err := json.Marshal(p{ipv4})
	if err != nil {
		return err
	}
	r, err := i.client.Do("POST", "ip/"+url.QueryEscape(block.IP)+"/firewall", string(payload))
	return r.HandleErr(err, []int{200})
}

// Remove IP from firewall
func (i *IpRessource) FwRemoveIp(block IpBlock, ipv4 string) error {
	r, err := i.client.Do("DELETE", "ip/"+url.QueryEscape(block.IP)+"/firewall/"+url.QueryEscape(ipv4), "")
	return r.HandleErr(err, []int{200})
}

// Get properties about an IP firewalled
func (i *IpRessource) FwGetIpProperties(block IpBlock, ipv4 string) (ip FirewalledIp, err error) {
	r, err := i.client.Do("GET", "ip/"+url.QueryEscape(block.IP)+"/firewall/"+url.QueryEscape(ipv4), "")
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &ip)
	return
}

// FwUpdateIp: update properties of an IP on firwall
func (i *IpRessource) FwUpdateIp(block IpBlock, ipv4 string, enabled bool) error {
	var err error
	var payload []byte
	type p struct {
		Enabled bool `json:"enabled"`
	}
	if payload, err = json.Marshal(p{enabled}); err != nil {
		return err
	}
	r, err := i.client.Do("PUT", "ip/"+url.QueryEscape(block.IP)+"/firewall/"+url.QueryEscape(ipv4), string(payload))
	return r.HandleErr(err, []int{200})
}

// FwAddRule adds a firewall rule
func (i *IpRessource) FwAddRule(block IpBlock, ipv4 string, rule FwRule2Add) error {
	payload, err := json.Marshal(rule)
	r, err := i.client.Do("POST", "ip/"+url.QueryEscape(block.IP)+"/firewall/"+url.QueryEscape(ipv4)+"/rule", string(payload))
	return r.HandleErr(err, []int{200})
}

// FwGetRulesSequences return rules sequences
func (i *IpRessource) FwListRules(block IpBlock, ipv4 string, state ...string) (sequences []int, err error) {
	uri := fmt.Sprintf("ip/%s/firewall/%s/rule", url.QueryEscape(block.IP), url.QueryEscape(ipv4))
	if len(state) > 0 {
		if !(state[0] == "creationPending" || state[0] == "ok" || state[0] == "removalPending") {
			err = errors.New("Bad parameter state (creationPending|ok|removalPending)")
			return
		}
		uri = uri + "?state=" + url.QueryEscape(state[0])
	}
	r, err := i.client.Do("GET", uri, "")
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &sequences)
	return
}

// FwGetRuleProperties returns rule Properties
func (i *IpRessource) FwGetRuleProperties(block IpBlock, ipv4 string, sequence int) (rule FirewallRule, err error) {
	r, err := i.client.Do("GET", "ip/"+url.QueryEscape(block.IP)+"/firewall/"+url.QueryEscape(ipv4)+"/rule/"+url.QueryEscape(fmt.Sprintf("%d", sequence)), "")
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &rule)
	return
}

// Remove a rule
func (i *IpRessource) FwRemoveRule(block IpBlock, ipv4 string, sequence int) error {
	r, err := i.client.Do("DELETE", "ip/"+url.QueryEscape(block.IP)+"/firewall/"+url.QueryEscape(ipv4)+"/rule/"+url.QueryEscape(fmt.Sprintf("%d", sequence)), "")
	return r.HandleErr(err, []int{200})
}

//
// SPAM
//

// SpamGetSpammingIps returns Spamming IP
// state :
// 		* blockedForSpam : currently blocked
// 		* unblocking : in the way to be unblocked (or not)
// 		* unblocked : blocked in the past
//
func (i *IpRessource) SpamGetSpammingIps(block IpBlock, state string) (ips []string, err error) {
	uri := fmt.Sprintf("ip/%s/spam", url.QueryEscape(block.IP))
	if state == "blockedForSpam" || state == "unblocking" || state == "unblocked" {
		uri = fmt.Sprintf("%s?state=%s", uri, state)
	}
	r, err := i.client.Do("GET", uri, "")
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &ips)
	return
}

// SpamGetIp returns detailed info about a spamming IP
func (i *IpRessource) SpamGetIp(block IpBlock, ipv4 string) (spamIp *SpamIp, err error) {
	r, err := i.client.Do("GET", "ip/"+url.QueryEscape(block.IP)+"/spam/"+url.QueryEscape(ipv4), "")
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &spamIp)
	return
}

// SpamGetIpStats returns stats about a spamming IP
func (i *IpRessource) SpamGetIpStats(block IpBlock, ipv4 string, from time.Time, to time.Time) (spamStats *SpamStats, err error) {
	uri := fmt.Sprintf("ip/%s/spam/%s/stats?from=%s&to=%s", url.QueryEscape(block.IP), ipv4, url.QueryEscape(from.Format(time.RFC3339)), url.QueryEscape(to.Format(time.RFC3339)))
	r, err := i.client.Do("GET", uri, "")
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	if len(r.Body) > 2 {
		err = json.Unmarshal(r.Body[1:len(r.Body)-1], &spamStats)
	}
	return
}

//SpamUnblockSpamIp Unblocks a spamming IP
func (i *IpRessource) SpamUnblockSpamIp(block IpBlock, ipv4 string) error {
	r, err := i.client.Do("POST", "ip/"+url.QueryEscape(block.IP)+"/spam/"+url.QueryEscape(ipv4)+"/unblock", "")
	return r.HandleErr(err, []int{200})
}

// GetBlockedForSpam returns IPs which are currently blocked for spam
func (i *IpRessource) GetBlockedForSpam() (ips []string, err error) {
	ipBlocks, err := i.List("", "", "", "")
	if err != nil {
		return
	}
	for _, ipb := range ipBlocks {
		// remove IPv6
		if len(strings.Split(ipb.IP, ":")) > 1 {
			continue
		}
		ipsBlocked, err := i.SpamGetSpammingIps(ipb, "blockedForSpam")
		if err != nil {
			// Not all IP are concerned by spamming status, if not found continue
			if strings.HasPrefix(err.Error(), "404 This service does not exist") {
				continue
			} else if strings.HasPrefix(err.Error(), "460 This Service is expired") {
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
