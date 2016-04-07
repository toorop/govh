package cloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// projectStatus represents the status of a project
type projectStatus int

const (
	ProjectStatusCreating projectStatus = 1 + iota
	ProjectStatusDeleting
	ProjectStatusOK
	ProjectStatusSuspended
)

func (status projectStatus) String() string {
	switch status {
	case ProjectStatusCreating:
		return "creating"
	case ProjectStatusDeleting:
		return "deleting"
	case ProjectStatusOK:
		return "ok"
	case ProjectStatusSuspended:
		return "suspended"
	}
	return "unknow"
}

// Project represents a cloud project
type Project struct {
	ID           string
	CreationDate time.Time
	Status       projectStatus
	Description  string
}

// project stringer
func (p Project) String() string {
	out := fmt.Sprintf("ID: %s\n", p.ID)
	out += fmt.Sprintf("Status: %s\n", p.Status)
	out += fmt.Sprintf("Creation date: %s\n", p.CreationDate.Format(time.RFC3339))
	out += fmt.Sprintf("Description: %s\n", p.Description)
	return out

}

// UnmarshalJSON is an unmarshaller
// TODO: refactor
func (p *Project) UnmarshalJSON(data []byte) (err error) {
	type resp struct {
		ID           string `json:"project_id"`
		CreationDate string `json:"creationDate"`
		Status       string `json:"status"`
		Description  string `json:"description"`
	}

	rp := new(resp)
	if err := json.Unmarshal(data, &rp); err != nil {
		return err
	}
	p.ID = rp.ID
	// 2015-05-06T20:20:26+02:00
	p.CreationDate, err = time.Parse(time.RFC3339, rp.CreationDate)
	if err != nil {
		return err
	}

	switch rp.Status {
	case "creating":
		p.Status = ProjectStatusCreating
	case "deleting":
		p.Status = ProjectStatusDeleting
	case "ok":
		p.Status = ProjectStatusOK
	case "suspended":
		p.Status = ProjectStatusSuspended
	default:
		return errors.New("unknow project status: " + rp.Status)
	}
	p.Description = rp.Description
	return nil
}
