package server

import (
	"github.com/Toorop/govh"
)

// Server properties
type Properties struct {
	Id              uint64 `json:"serverId"`        // Server ID
	Name            string `json:"Name"`            // Server name (for OVH)
	Ip              string `json:"ip"`              // Main server IPv4
	Datacenter      string `json:"datacenter"`      // Datacenter localisation
	ProfessionalUse bool   `json:"professionalUse"` // Does this server have professional use option"
	CommercialRange string `json:"commercialRange"` // Commercial range
	SupportLevel    string `json:"supportLevel"`    // Support level
	Os              string `json:"os"`              // Operating system
	State           string `json:"state"`           // State of the server (ok|error|hacked)
	Reverse         string `json:"reverse"`         // Main IP reverse
	Monitored       bool   `json:"monitoring"`      // Icmp monitoring state
	Rack            string `json:"rack"`            // Rack
	RootDevice      string `json:"rootDevice"`      // Root device
	LinkSpeed       uint64 `json:"linkSpeed"`       // Inteface speed
	BootId          uint32 `json:"bootid"`          // Boot id
}

// Server task
type Task struct {
	Id         uint64        `json:"taskId"`     // ID of the task
	Function   string        `json:"function"`   // Function name
	LastUpdate govh.DateTime `json:"lastUpdate"` // Last update
	Comment    string        `json:"comment"`    // Details of this task
	Status     string        `json:"status"`     // Task status
	StartDate  govh.DateTime `json:"startDate"`  // Task Creation date
	DoneDate   govh.DateTime `json:"doneDate"`   // "Completion date"
}
