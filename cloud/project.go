package cloud

import (
	"encoding/json"
	"errors"
	//"fmt"
	"time"

	//"github.com/toorop/govh"
)

// projectStatus represents the status of a project
type projectStatus int

const (
	CREATING projectStatus = 1 + iota
	DELETING
	OK
	SUSPENDED
)

func (status projectStatus) String() string {
	switch status {
	case CREATING:
		return "creating"
	case DELETING:
		return "deleting"
	case OK:
		return "ok"
	case SUSPENDED:
		return "suspended"
	}
	return "unknow"
}

// project represents a cloud project
type project struct {
	Id           string
	CreationDate time.Time
	Status       projectStatus
	Description  string
}

func (p *project) UnmarshalJSON(data []byte) (err error) {
	type resp struct {
		Id           string `json:"project_id"`
		CreationDate string `json:"creationDate"`
		Status       string `json:"status"`
		Description  string `json:"description"`
	}

	rp := new(resp)
	if err := json.Unmarshal(data, &rp); err != nil {
		return err
	}
	p.Id = rp.Id
	// 2015-05-06T20:20:26+02:00
	p.CreationDate, err = time.Parse(time.RFC3339, rp.CreationDate)
	if err != nil {
		return err
	}

	switch rp.Status {
	case "creating":
		p.Status = CREATING
	case "deleting":
		p.Status = DELETING
	case "ok":
		p.Status = OK
	case "suspended":
		p.Status = SUSPENDED
	default:
		return errors.New("unknow project status: " + rp.Status)
	}
	p.Description = rp.Description
	return nil
}
