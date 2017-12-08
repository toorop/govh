package me

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"time"

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

// GetMe return User from GET /me
func (c *Client) GetMe() (user User, err error) {
	r, err := c.GET("me")
	if err != nil {
		return
	}
	log.Println(string(r.Body))
	err = json.Unmarshal(r.Body, &user)
	return
}

// GetBillIDs return IDs of bill from dateFrom to dateTo
func (c *Client) GetBillIDs(dateFrom, dateTo time.Time) (billIDs []string, err error) {
	// format date RFC ISO 8601 - time.RFC3339
	from := dateFrom.Format(time.RFC3339)
	to := dateTo.Format(time.RFC3339)
	r, err := c.GET("me/bill?date.from=" + url.QueryEscape(from) + "&date.to=" + url.QueryEscape(to))
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &billIDs)
	return
}

// GetBillByID return Bill by its ID
func (c *Client) GetBillByID(ID string) (bill Bill, err error) {
	r, err := c.GET("me/bill/" + url.QueryEscape(ID))
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &bill)
	return
}
