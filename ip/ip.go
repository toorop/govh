package ip

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/toorop/govh"
)

// Client is a REST client for server API
type Client struct {
	*govh.OVHClient
}

// New return a new server Client
func New(client *govh.OVHClient) (*Client, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return &Client{client}, nil
}

// IPBlock

// List return a slice of IpBlock
func (c *Client) List(filterDesc, filterIP, filterRoutedTo, filterType string) (ips []IPBlock, err error) {
	uri := "ip"
	args := []string{}

	if len(filterDesc) > 0 {
		args = append(args, "description="+url.QueryEscape(filterDesc))
	}
	if len(filterIP) > 0 {
		args = append(args, "ip="+url.QueryEscape(filterIP))
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
	r, err := c.GET(uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &ips)
	return
}

// GetBlockProperties return properties of an IP
func (c *Client) GetBlockProperties(block IPBlock) (ip IP, err error) {
	ip = IP{}
	r, err := c.GET("ip/" + url.QueryEscape(string(block)))
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &ip)
	return
}

// UpdateBlockProperties update IP properties
func (c *Client) UpdateBlockProperties(block IPBlock, desc string) error {
	payload, err := json.Marshal(IpUpdatableProperties{
		Description: desc,
	})
	if err != nil {
		return err
	}
	r, err := c.PUT("ip/"+url.QueryEscape(string(block)), string(payload))
	err = r.HandleErr(err, []int{200})
	return err
}

// Change IPBlock destination
func (c *Client) Move(block IPBlock, toServiceName string) (*IpTask,error) {
	iptask := IpTask{}
	payload, err := json.Marshal(MoveTo{
		To: toServiceName,
	})

	if err != nil {
		return nil,err
	}
	r, err := c.POST("ip/"+url.QueryEscape(string(block))+"/move", string(payload))
	err = r.HandleErr(err, []int{200})

	if err != nil {
		return nil,err
	}

	err = json.Unmarshal(r.Body, &iptask)

	if err != nil {
		return nil,err
	}
	return &iptask,err
}

// Get IpTask
func (c *Client) Task(block IPBlock, taskid IPTaskId) (*IpTask,error) {
	iptask := IpTask{}
	r, err := c.GET("ip/"+url.QueryEscape(string(block))+fmt.Sprintf("/task/%d",taskid))

	if err != nil {
		return nil,err
	}

	err = json.Unmarshal(r.Body, &iptask)

	if err != nil {
		return nil,err
	}

	return &iptask,nil
}

// IPs
/*
//
//// LOADBALANCING
//

// List IP load balancing
func (r *Client) LbList() (resp []byte, err error) {
	resp, err = r.client.Do("GET", "ip/loadBalancing", "")
	return
}

//
//// FIREWALL
//

// FwListIPOfBlock List IP of block IP under firewall protection
func (c *Client) FwListIPOfBlock(block IPBlock) (IPs []string, err error) {
	r, err := c.GET("ip/" + url.QueryEscape(block.IP) + "/firewall")
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &IPs)
	return
}

// FwAddIP Add IP to firewall
func (c *Client) FwAddIP(block IPBlock, IPv4 string) error {
	type p struct {
		IPOnFirewall string `json:"ipOnFirewall"`
	}
	payload, err := json.Marshal(p{IPv4})
	if err != nil {
		return err
	}
	r, err := c.POST("ip/"+url.QueryEscape(block.IP)+"/firewall", string(payload))
	return r.HandleErr(err, []int{200})
}

// FwRemoveIP Remove IP from firewall
func (c *Client) FwRemoveIP(block IPBlock, IPv4 string) error {
	r, err := c.DELETE("ip/" + url.QueryEscape(block.IP) + "/firewall/" + url.QueryEscape(IPv4))
	return r.HandleErr(err, []int{200})
}

// FwGetIPProperties Get properties about an IP firewalled
func (c *Client) FwGetIPProperties(block IPBlock, IPv4 string) (IP FirewalledIp, err error) {
	r, err := c.GET("ip/" + url.QueryEscape(block.IP) + "/firewall/" + url.QueryEscape(IPv4))
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &IP)
	return
}

// FwUpdateIP  update properties of an IP on firewall
func (c *Client) FwUpdateIP(block IPBlock, IPv4 string, enabled bool) error {
	var err error
	var payload []byte
	type p struct {
		Enabled bool `json:"enabled"`
	}
	if payload, err = json.Marshal(p{enabled}); err != nil {
		return err
	}
	r, err := c.PUT("ip/"+url.QueryEscape(block.IP)+"/firewall/"+url.QueryEscape(IPv4), string(payload))
	return r.HandleErr(err, []int{200})
}

// FwAddRule adds a firewall rule
func (c *Client) FwAddRule(block IPBlock, IPv4 string, rule FwRule2Add) error {
	payload, err := json.Marshal(rule)
	r, err := c.POST("ip/"+url.QueryEscape(block.IP)+"/firewall/"+url.QueryEscape(IPv4)+"/rule", string(payload))
	return r.HandleErr(err, []int{200})
}

// FwListRules return rules sequences
func (c *Client) FwListRules(block IPBlock, IPv4 string, state ...string) (sequences []int, err error) {
	uri := fmt.Sprintf("ip/%s/firewall/%s/rule", url.QueryEscape(block.IP), url.QueryEscape(IPv4))
	if len(state) > 0 {
		if !(state[0] == "creationPending" || state[0] == "ok" || state[0] == "removalPending") {
			err = errors.New("Bad parameter state (creationPending|ok|removalPending)")
			return
		}
		uri = uri + "?state=" + url.QueryEscape(state[0])
	}
	r, err := c.GET(uri)
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &sequences)
	return
}

// FwGetRuleProperties returns rule Properties
func (c *Client) FwGetRuleProperties(block IPBlock, IPv4 string, sequence int) (rule FirewallRule, err error) {
	r, err := c.GET("ip/" + url.QueryEscape(block.IP) + "/firewall/" + url.QueryEscape(IPv4) + "/rule/" + url.QueryEscape(fmt.Sprintf("%d", sequence)))
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &rule)
	return
}

//FwRemoveRule removes a rule
func (c *Client) FwRemoveRule(block IPBlock, IPv4 string, sequence int) error {
	r, err := c.DELETE("ip/" + url.QueryEscape(block.IP) + "/firewall/" + url.QueryEscape(IPv4) + "/rule/" + url.QueryEscape(fmt.Sprintf("%d", sequence)))
	return r.HandleErr(err, []int{200})
}
*/
// REVERSE

// GetReverse return reverse of IP ip
func (c *Client) GetReverse(IP string) (string, error) {
	RIP := ReverseIP{}
	r, err := c.GET("ip/" + url.QueryEscape(IP+"/32") + "/reverse/" + url.QueryEscape(IP))
	if err != nil {
		return "", err
	}
	err = json.Unmarshal([]byte(r.Body), &RIP)
	return RIP.Reverse, err
}

// SetReverse set the reverse of an IP
func (c *Client) SetReverse(IP, reverse string) error {
	payload, err := json.Marshal(ReverseIP{
		IPReverse: IP,
		Reverse:   reverse,
	})
	if err != nil {
		return err
	}
	r, err := c.POST("ip/"+url.QueryEscape(IP+"/32")+"/reverse", string(payload))
	if err != nil {
		return err
	}
	RIP := ReverseIP{}
	if err = json.Unmarshal([]byte(r.Body), &RIP); err != nil {
		return err
	}
	if RIP.IPReverse != IP || RIP.Reverse != reverse+"." {
		return fmt.Errorf("returned reverseIP doesn't match reverseIP expected. Expected: %s %s got %s", IP, reverse, RIP.String())
	}
	return nil
}
