package server

import (
	"encoding/json"
	"errors"
	//"fmt"
	"github.com/Toorop/govh"
	"net/url"
	"strconv"
)

type ServerRessource struct {
	client *govh.OvhClient
}

// New returns a new Server ressource
func New(client *govh.OvhClient) (*ServerRessource, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return &ServerRessource{client: client}, nil
}

// List returns a list (slice) of servers name
func (s *ServerRessource) List() (servers []string, err error) {
	r, err := s.client.Do("GET", "dedicated/server", "")
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &servers)
	return
}

// GetProperties returns server properties
func (s *ServerRessource) GetProperties(serverName string) (properties Properties, err error) {
	r, err := s.client.Do("GET", "dedicated/server/"+url.QueryEscape(serverName), "")
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
func (r *ServerRessource) GetNetboots(serverName string, bootType ...string) (netbootIds []int, err error) {
	if len(bootType) > 1 {
		err = errors.New("Bad method call")
		return
	}
	ressource := "dedicated/server/" + serverName + "/boot"
	if len(bootType) == 1 {
		ressource = ressource + "?bootType=" + bootType[0]
	}
	t, err := r.client.Do("GET", ressource, "")
	if err != nil {
		return
	}
	err = json.Unmarshal(t, &netbootIds)
	return
}
*/
// GetTasks returns tasks for server serverName
func (s *ServerRessource) GetTasks(serverName, function, status string) (taskId []uint32, err error) {
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
	r, err := s.client.Do("GET", uri, "")
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &taskId)
	return
}

// GetTaskProperties returns a server task
func (s *ServerRessource) GetTaskProperties(serverName string, taskId uint64) (task Task, err error) {
	r, err := s.client.Do("GET", "dedicated/server/"+url.QueryEscape(serverName)+"/task/"+url.QueryEscape(strconv.FormatUint(taskId, 10)), "")
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &task)
	return
}

// Cancel a task
func (s *ServerRessource) CancelTask(serverName string, taskId uint64) (err error) {
	r, err := s.client.Do("POST", "dedicated/server/"+url.QueryEscape(serverName)+"/task/"+url.QueryEscape(strconv.FormatUint(taskId, 10))+"/cancel", "")
	err = r.HandleErr(err, []int{200})
	return
}

// reboot creates a new server task
func (s *ServerRessource) Reboot(serverName string) (task Task, err error) {
	r, err := s.client.Do("POST", "dedicated/server/"+url.QueryEscape(serverName)+"/reboot", "")
	if err = r.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &task)
	return
}
