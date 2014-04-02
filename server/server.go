package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Toorop/govh"
)

type ServerRessource struct {
	client *govh.OvhClient
}

func New(client *govh.OvhClient) (*ServerRessource, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return &ServerRessource{client: client}, nil
}

// List servers
func (r *ServerRessource) List() (servers []string, err error) {
	t, err := r.client.Do("GET", "dedicated/server", "")
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(t, &servers); err != nil {
		return nil, err
	}
	return
}

// Get server properties
func (r *ServerRessource) GetProperties(serverName string) (properties Properties, err error) {
	t, err := r.client.Do("GET", "dedicated/server/"+serverName, "")
	if err != nil {
		return properties, err
	}
	if err = json.Unmarshal(t, &properties); err != nil {
		return properties, err
	}
	return
}

// Get server tasks
// function & status are filter, uses "all" as wildcard
func (r *ServerRessource) GetTasks(serverName, function, status string) (taskId []uint32, err error) {
	uri := "dedicated/server/" + serverName + "/task"
	if function != "all" || status != "all" {
		uri += "?"
	}
	if function != "all" {
		uri += "function=" + function
	}
	if status != "all" {
		if function != "all" {
			uri += "&"
		}
		uri += "status=" + status
	}
	t, err := r.client.Do("GET", uri, "")
	if err != nil {
		return taskId, err
	}
	if err = json.Unmarshal(t, &taskId); err != nil {
		return taskId, err
	}
	return
}

// Get task properties
func (r *ServerRessource) GetTaskProperties(serverName string, taskId uint64) (task Task, err error) {
	uri := fmt.Sprintf("dedicated/server/%s/task/%d", serverName, taskId)
	t, err := r.client.Do("GET", uri, "")
	if err != nil {
		return task, err
	}
	if err = json.Unmarshal(t, &task); err != nil {
		return task, err
	}
	return
}

// Cancel a task
func (r *ServerRessource) CancelTask(serverName string, taskId uint64) (err error) {
	uri := fmt.Sprintf("dedicated/server/%s/task/%d/cancel", serverName, taskId)
	_, err = r.client.Do("POST", uri, "")
	if err != nil {
		return err
	}
	return
}

// Reboot server
func (r *ServerRessource) Reboot(serverName string) (task Task, err error) {
	t, err := r.client.Do("POST", "dedicated/server/"+serverName+"/reboot", "")
	if err != nil {
		return task, err
	}
	if err = json.Unmarshal(t, &task); err != nil {
		return task, err
	}
	return
}
