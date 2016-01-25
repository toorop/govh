package server

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

// New returns a new Server ressource
func New(client *govh.OVHClient) (*Client, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return &Client{client}, nil
}

// List returns a list (slice) of servers name
func (c *Client) List() (servers []string, err error) {
	r, err := c.GET("dedicated/server")
	err = json.Unmarshal(r.Body, &servers)
	return
}

// GetProperties returns server properties
func (c *Client) GetProperties(serverName string) (properties Dedicated, err error) {
	r, err := c.GET("dedicated/server/" + url.QueryEscape(serverName))
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &properties)
	return
}

/*
// Setproperties update server properties
// TODO

// GetNetboots return compatible netboots (as int ID) for server serverName
func (r *Client) GetNetboots(serverName string, bootType ...string) (netbootIds []int, err error) {
	if len(bootType) > 1 {
		err = errors.New("Bad method call")
		return
	}
	ressource := "dedicated/server/" + serverName + "/boot"
	if len(bootType) == 1 {
		ressource = ressource + "?bootType=" + bootType[0]
	}
	t, err := r.Query("GET", ressource, "")
	if err != nil {
		return
	}
	err = json.Unmarshal(t, &netbootIds)
	return
}
*/

// GetTasks returns tasks for server serverName
func (c *Client) GetTasks(serverName, function, status string) (taskID []int, err error) {
	uri := "dedicated/server/" + url.QueryEscape(serverName) + "/task"
	v := url.Values{}
	if function != "" {
		v.Add("function", function)
	}
	if status != "" {
		v.Add("status", status)
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

// GetTask returns a server task
func (c *Client) GetTask(serverName string, taskID int) (task Task, err error) {
	r, err := c.GET("dedicated/server/" + url.QueryEscape(serverName) + "/task/" + fmt.Sprintf("%d", taskID))
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &task)
	return
}

// CancelTask cancels a task
func (c *Client) CancelTask(serverName string, taskID int) (err error) {
	_, err = c.POST("dedicated/server/"+url.QueryEscape(serverName)+"/task/"+fmt.Sprintf("%d", taskID)+"/cancel", "")
	return
}

// Reboot reboot server serverName
func (c *Client) Reboot(serverName string) (task Task, err error) {
	r, err := c.POST("dedicated/server/"+url.QueryEscape(serverName)+"/reboot", "")
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &task)
	return
}
