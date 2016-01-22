package domain

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"

	"github.com/toorop/govh"
)

// Client is an OVH API client
type Client struct {
	*govh.OVHClient
}

// New return a new Client
func New(client *govh.OVHClient) (*Client, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return &Client{client}, nil
}

// List return a list of domains
func (c *Client) List(whoisOwner ...string) (domains []string, err error) {
	uri := "domain"
	if len(whoisOwner) != 0 {
		uri += "?whoisOwner=" + url.QueryEscape(strings.Join(whoisOwner, ""))
	}
	r, err := c.GET(uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &domains)
	return
}

// GetRecordIDs return record ID for the zone zone
func (c *Client) GetRecordIDs(zone string) (IDs int, err error) {
	return
}

// GetRecordByID return a ZoneRecord by its ID
func (c *Client) GetRecordByID(zone string, ID int) (record ZoneRecord, err error) {
	return
}
