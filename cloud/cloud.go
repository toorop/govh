package cloud

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/toorop/govh"
)

// Client is a REST client for cloud API
type Client struct {
	*govh.OVHClient
}

// New return a new Cloud API Client
func New(client *govh.OVHClient) (*Client, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return &Client{client}, nil
}

// GetPassports returns cloud passports
func (c *Client) GetPassports() (passports []string, err error) {
	r, err := c.GET("cloud")
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &passports)
	return
}

// GetPrices return cloud prices
func (c *Client) GetPrices() (prices GetPriceResponse, err error) {
	var r govh.APIResponse
	r, err = c.GET("cloud/price")
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &prices)
	return
}

// GetProjects returns clouds projects
func (c *Client) GetProjectsId() (projectid []string, err error) {
	r, err := c.GET("cloud/project")
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &projectid)
	return
}

// GetProject return a project
func (c *Client) GetProject(id string) (p project, err error) {
	r, err := c.GET("cloud/project/" + url.QueryEscape(id))
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &p)
	return
}
