package server

import (
	"fmt"

	"github.com/toorop/govh"
)

// Dedicated represents dedicated.server.Dedicater OVH type
type Dedicated struct {
	ID              int    `json:"serverId"`        // Server ID
	Name            string `json:"Name"`            // Server name (OVH name)
	IP              string `json:"ip"`              // Main server IPv4
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
	BootID          uint32 `json:"bootid"`          // Boot id
}

// String represents Dedicated as string
func (d Dedicated) String() string {
	out := fmt.Sprintf("ID: %d\n", d.ID)
	out += fmt.Sprintf("Name: %s\n", d.Name)
	out += fmt.Sprintf("IP: %s\n", d.IP)
	out += fmt.Sprintf("Datacenter: %s\n", d.Datacenter)
	out += fmt.Sprintf("Support Level: %s\n", d.SupportLevel)
	out += fmt.Sprintf("Professional Use: %t\n", d.ProfessionalUse)
	out += fmt.Sprintf("Commercial Range: %s\n", d.CommercialRange)
	out += fmt.Sprintf("OS: %s\n", d.Os)
	out += fmt.Sprintf("State: %s\n", d.State)
	out += fmt.Sprintf("Reverse: %s\n", d.Reverse)
	out += fmt.Sprintf("Monitored: %t\n", d.Monitored)
	out += fmt.Sprintf("Rack: %s\n", d.Rack)
	out += fmt.Sprintf("Root Device: %s\n", d.RootDevice)
	out += fmt.Sprintf("Link Speed: %d\n", d.LinkSpeed)
	out += fmt.Sprintf("Boot ID: %d", d.BootID)
	return out

}

// Task represents OVH dedictaded.server.Task
type Task struct {
	ID         int           `json:"taskId"`     // ID of the task
	Function   string        `json:"function"`   // Function name
	LastUpdate govh.DateTime `json:"lastUpdate"` // Last update
	Comment    string        `json:"comment"`    // Details of this task
	Status     string        `json:"status"`     // Task status
	StartDate  govh.DateTime `json:"startDate"`  // Task Creation date
	DoneDate   govh.DateTime `json:"doneDate"`   // "Completion date"
}

// Stringer
func (t Task) String() string {
	return fmt.Sprintf("ID: %d\nLast update: %s\nFunction: %s\nStatus: %s\nComments: %s\nStart at: %s\nDone at: %s", t.ID, t.LastUpdate, t.Function, t.Status, t.Comment, t.StartDate, t.DoneDate)
}
