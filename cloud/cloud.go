package cloud

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/toorop/govh"
)

type CloudRessource struct {
	client *govh.OvhClient
}

// New return a new CloudRessource
func New(client *govh.OvhClient) (*CloudRessource, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return &CloudRessource{client: client}, nil
}

// GetPassports returns cloud passports
func (c *CloudRessource) GetPassports() (passports []string, err error) {
	resp, err := c.client.Do("GET", "cloud", "")
	if err = resp.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(resp.Body, &passports)
	return
}

// GetProjects returns clouds projects
func (c *CloudRessource) GetProjectsId() (projectid []string, err error) {
	resp, err := c.client.Do("GET", "cloud/project", "")
	if err = resp.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(resp.Body, &projectid)
	return
}

// GetProject return a project
func (c *CloudRessource) GetProject(id string) (p project, err error) {
	resp, err := c.client.Do("GET", "cloud/project/"+url.QueryEscape(id), "")
	if err = resp.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(resp.Body, &p)
	return
}
