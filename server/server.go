package server

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"

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
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &servers)
	return
}

// GetProperties returns server properties
func (c *Client) GetProperties(serverName string) (properties Properties, err error) {
	r, err := c.GET("dedicated/server/" + url.QueryEscape(serverName))
	if err = r.HandleErr(err, []int{200}); err != nil {
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
func (c *Client) GetTasks(serverName, function, status string) (taskID []uint32, err error) {
	uri := "dedicated/server/" + url.QueryEscape(serverName) + "/task"
	if len(function) != 0 || len(status) != 0 {
		uri += "?"
	}
	if function != "" {
		uri += "function=" + function
	}
	if status != "" {
		if function != "" {
			uri += "&"
		}
		uri += "status=" + status
	}
	r, err := c.GET(uri)
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &taskID)
	return
}

// GetTaskProperties returns a server task
func (c *Client) GetTaskProperties(serverName string, taskID uint64) (task Task, err error) {
	r, err := c.GET("dedicated/server/" + url.QueryEscape(serverName) + "/task/" + url.QueryEscape(strconv.FormatUint(taskID, 10)))
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &task)
	return
}

// CancelTask cancels a task
func (c *Client) CancelTask(serverName string, taskID uint64) (err error) {
	r, err := c.POST("dedicated/server/"+url.QueryEscape(serverName)+"/task/"+url.QueryEscape(strconv.FormatUint(taskID, 10))+"/cancel", "")
	err = r.HandleErr(err, []int{200})
	return
}

// Reboot reboot server
func (c *Client) Reboot(serverName string) (task Task, err error) {
	r, err := c.POST("dedicated/server/"+url.QueryEscape(serverName)+"/reboot", "")
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &task)
	return
}
