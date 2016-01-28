package ip

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

//
// SPAM
//

// SpamGetIPs returns Spamming IP
// state :
// 		* blockedForSpam : currently blocked
// 		* unblocking : in the way to be unblocked (or not)
// 		* unblocked : blocked in the past
//
func (c *Client) SpamGetIPs(block IPBlock, state string) (IPs []string, err error) {
	uri := "ip/" + url.QueryEscape(string(block)) + "/spam"
	if state != "" {
		uri += "?state=" + state
	}
	r, err := c.GET(uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &IPs)
	return
}

// SpamGetIP returns detailed info about a spamming IP
func (c *Client) SpamGetIP(block IPBlock, IP string) (spamIP SpamIP, err error) {
	r, err := c.GET("ip/" + url.QueryEscape(string(block)) + "/spam/" + url.QueryEscape(IP))
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &spamIP)
	return
}

// SpamGetIPStats returns stats about a spamming IP
func (c *Client) SpamGetIPStats(block IPBlock, IP string, from time.Time, to time.Time) (spamStats []SpamStats, err error) {
	r, err := c.GET(fmt.Sprintf("ip/%s/spam/%s/stats?from=%s&to=%s", url.QueryEscape(string(block)), url.QueryEscape(IP), url.QueryEscape(from.Format(time.RFC3339)), url.QueryEscape(to.Format(time.RFC3339))))
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &spamStats)
	return
}

//SpamUnblockSpamIP Unblocks a spamming IP
func (c *Client) SpamUnblockSpamIP(block IPBlock, IP string) error {
	_, err := c.POST("ip/"+url.QueryEscape(string(block))+"/spam/"+url.QueryEscape(IP)+"/unblock", "")
	return err
}

// GetBlockedForSpam returns IPs which are currently blocked for spam
func (c *Client) GetBlockedForSpam() (IPs []string, err error) {
	IPBlocks, err := c.List("", "", "", "")
	if err != nil {
		return
	}
	for _, block := range IPBlocks {
		// remove IPv6
		if len(strings.Split(string(block), ":")) > 1 {
			continue
		}
		IPsBlocked, err := c.SpamGetIPs(block, "blockedForSpam")
		if err != nil {
			// Not all IP are concerned by spamming status, if not found continue
			if strings.HasPrefix(err.Error(), "404 This service does not exist") {
				continue
			} else if strings.HasPrefix(err.Error(), "460 This Service is expired") {
				continue
			}
			return IPs, err
		}
		for _, IP := range IPsBlocked {
			IPs = append(IPs, IP)
		}
	}
	return
}
