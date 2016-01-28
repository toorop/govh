package dedicatedcloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/toorop/govh"
)

// Client is a REST client for cloud API
type Client struct {
	*govh.OVHClient
}

// New returns a new Dedicated Cloud ressource
func New(client *govh.OVHClient) (*Client, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return &Client{client}, nil
}

// List returns a list (slice) of Dedicated Cloud name
func (c *Client) List() (dedicatedClouds []string, err error) {
	r, err := c.GET("dedicatedCloud")
	err = json.Unmarshal(r.Body, &dedicatedClouds)
	return
}

// GetProperties returns Dedicated Cloud properties
func (c *Client) GetProperties(dedicatedCloudName string) (properties DedicatedCloud, err error) {
	r, err := c.GET("dedicatedCloud/" + url.QueryEscape(dedicatedCloudName))
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &properties)
	return
}

// GetAllowedNetworks returns Allowed Networks for Dedicated Cloud dedicatedCloudName
func (c *Client) GetAllowedNetworks(dedicatedCloudName string) (allowedNetworkID []int, err error) {
	uri := "dedicatedCloud/" + url.QueryEscape(dedicatedCloudName) + "/allowedNetwork"
	v := url.Values{}

	if t := v.Encode(); t != "" {
		uri += "?" + t
	}
	r, err := c.GET(uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &allowedNetworkID)
	return
}

// GetAllowedNetwork returns a Dedicated Cloud AllowedNetwork
func (c *Client) GetAllowedNetwork(dedicatedCloudName string, allowedNetworkID int) (allowedNetwork AllowedNetwork, err error) {
	r, err := c.GET("dedicatedCloud/" + url.QueryEscape(dedicatedCloudName) + "/allowedNetwork/" + fmt.Sprintf("%d", allowedNetworkID))
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &allowedNetwork)
	return
}

// GetUsers returns users for Dedicated Cloud dedicatedCloudName
func (c *Client) GetUsers(dedicatedCloudName, name string) (userID []int, err error) {
	uri := "dedicatedCloud/" + url.QueryEscape(dedicatedCloudName) + "/user"
	v := url.Values{}
	if name != "" {
		v.Add("name", name)
	}

	if t := v.Encode(); t != "" {
		uri += "?" + t
	}
	r, err := c.GET(uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &userID)
	return
}

// GetUser returns a Dedicated Cloud user
func (c *Client) GetUser(dedicatedCloudName string, userID int) (user User, err error) {
	r, err := c.GET("dedicatedCloud/" + url.QueryEscape(dedicatedCloudName) + "/user/" + fmt.Sprintf("%d", userID))
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &user)
	return
}

// GetDatacenters returns datacenters for Dedicated Cloud dedicatedCloudName
func (c *Client) GetDatacenters(dedicatedCloudName string) (datacenterID []int, err error) {
	uri := "dedicatedCloud/" + url.QueryEscape(dedicatedCloudName) + "/datacenter"
	v := url.Values{}

	if t := v.Encode(); t != "" {
		uri += "?" + t
	}
	r, err := c.GET(uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &datacenterID)
	return
}

// GetDatacenter returns a Dedicated Cloud datacenter
func (c *Client) GetDatacenter(dedicatedCloudName string, datacenterID int) (datacenter Datacenter, err error) {
	r, err := c.GET("dedicatedCloud/" + url.QueryEscape(dedicatedCloudName) + "/datacenter/" + fmt.Sprintf("%d", datacenterID))
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &datacenter)
	return
}

// GetTasks returns tasks for Dedicated Cloud dedicatedCloudName
func (c *Client) GetTasks(dedicatedCloudName, state string) (taskID []int, err error) {
	uri := "dedicatedCloud/" + url.QueryEscape(dedicatedCloudName) + "/task"
	v := url.Values{}
	if state != "" {
		v.Add("state", state)
	}

	if t := v.Encode(); t != "" {
		uri += "?" + t
	}
	r, err := c.GET(uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &taskID)
	return
}

// GetTask returns a Dedicated Cloud task
func (c *Client) GetTask(dedicatedCloudName string, taskID int) (task Task, err error) {
	r, err := c.GET("dedicatedCloud/" + url.QueryEscape(dedicatedCloudName) + "/task/" + fmt.Sprintf("%d", taskID))
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &task)
	return
}
