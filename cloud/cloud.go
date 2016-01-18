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
	resp, err := c.GET("cloud")
	if err = resp.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(resp.Body, &passports)
	return
}

// GetProjects returns clouds projects
func (c *Client) GetProjectsId() (projectid []string, err error) {
	resp, err := c.GET("cloud/project")
	if err = resp.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(resp.Body, &projectid)
	return
}

// GetProject return a project
func (c *Client) GetProject(id string) (p project, err error) {
	resp, err := c.GET("cloud/project/" + url.QueryEscape(id))
	if err = resp.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(resp.Body, &p)
	return
}
