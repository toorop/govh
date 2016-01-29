package dedicatedcloud

import (
	"fmt"

	"github.com/toorop/govh"
)

// DedicatedCloud represents dedicatedCloud.dedicatedCloud OVH type
type DedicatedCloud struct {
	Name                       string `json:"serviceName"`                // Dedicated Cloud name
	State                      string `json:"state"`                      // State
	Description                string `json:"description"`                // Icmp monitoring state
	CommercialRange            string `json:"commercialRange"`            // Commercial range
	Type                       string `json:"managementInterface"`        // Management interface
	BillingType                string `json:"billingType"`                // Billing type (demo|monthly)
	Location                   string `json:"location"`                   // Location
	SslV3                      bool   `json:"sslV3"`                      // Is SSL v3 support enabled
	Spla                       bool   `json:"spla"`                       // Is Spla licensing enabled
	UserAccessPolicy           string `json:"userAccessPolicy"`           // Access policy (open|filtered)
	UserLogoutPolicy           string `json:"userLogoutPolicy"`           // Logout policy (first|last)
	UserLimitConcurrentSession int    `json:"userLimitConcurrentSession"` // User concurrent sessions limit
	UserSessionTimeout         int    `json:"userSessionTimeout"`         // User session timeout
	Bandwidth                  string `json:"bandwidth"`                  // Bandwidth
	Vscope                     string `json:"vScopeUrl"`                  // vScope url
}

// String represents DedicatedCloud as string
func (dc DedicatedCloud) String() string {
	out := fmt.Sprintf("Name: %s\n", dc.Name)
	out += fmt.Sprintf("State: %s\n", dc.State)
	out += fmt.Sprintf("Description: %s\n", dc.Description)
	out += fmt.Sprintf("Commercial Range: %s\n", dc.CommercialRange)
	out += fmt.Sprintf("Management Interface : %s\n", dc.Type)
	out += fmt.Sprintf("Billing Type: %s\n", dc.BillingType)
	out += fmt.Sprintf("Location: %s\n", dc.Location)
	out += fmt.Sprintf("SslV3: %t\n", dc.SslV3)
	out += fmt.Sprintf("Spla: %t\n", dc.Spla)
	out += fmt.Sprintf("User Access Policy: %s\n", dc.UserAccessPolicy)
	out += fmt.Sprintf("User Logout Policy: %s\n", dc.UserLogoutPolicy)
	out += fmt.Sprintf("User Limit Concurrent Session: %d\n", dc.UserLimitConcurrentSession)
	out += fmt.Sprintf("User Session Timeout: %d\n", dc.UserSessionTimeout)
	out += fmt.Sprintf("Bandwidth: %s\n", dc.Bandwidth)
	out += fmt.Sprintf("Vscope Url: %s", dc.Vscope)
	return out
}

// AllowedNetwork represents OVH dedicatedCloud.AllowedNetwork
type AllowedNetwork struct {
	ID      int    `json:"networkAccessId"` // ID of the network
	Network string `json:"network"`         // Network name
	State   string `json:"state"`           // Network status
}

// Stringer
func (n AllowedNetwork) String() string {
	return fmt.Sprintf("ID: %d\nNetwork: %s\nState: %s", n.ID, n.Network, n.State)
}

// User represents OVH dedicatedCloud.User
type User struct {
	ID                   int    `json:"userId"`               // ID of the user
	Name                 string `json:"name"`                 // User name
	Email                string `json:"email"`                // Email
	State                string `json:"state"`                // State
	ActivationState      string `json:"activationState"`      // Activation State
	IsEnableManageable   bool   `json:"isEnableManageable"`   // Is Enable Manageable
	CanManageNetwork     bool   `json:"canManageNetwork"`     // Can Manage Network
	CanManageIpFailOvers bool   `json:"canManageIpFailOvers"` // Can Manage Ip FailOvers
	FullAdminRo          bool   `json:"fullAdminRo"`          // Full Admin RO
}

// String represents User as string
func (u User) String() string {
	out := fmt.Sprintf("ID: %d\n", u.ID)
	out += fmt.Sprintf("Name: %s\n", u.Name)
	out += fmt.Sprintf("Email: %s\n", u.Email)
	out += fmt.Sprintf("State: %s\n", u.State)
	out += fmt.Sprintf("Activation State: %s\n", u.ActivationState)
	out += fmt.Sprintf("Is Enable Manageable: %t\n", u.IsEnableManageable)
	out += fmt.Sprintf("Can Manage Network: %t\n", u.CanManageNetwork)
	out += fmt.Sprintf("Can Manage Ip FailOvers: %t\n", u.CanManageIpFailOvers)
	out += fmt.Sprintf("Full Admin RO: %t\n", u.FullAdminRo)
	return out
}

// Datacenter represents OVH dedicatedCloud.Datacenter
type Datacenter struct {
	ID                  int    `json:"datacenterId"`        // ID of the datacenter
	Name                string `json:"name"`                // Datacenter name
	Version             string `json:"version"`             // Version
	Description         string `json:"description"`         // Description
	CommercialRangeName string `json:"commercialRangeName"` // Commercial Range Name
	IsRemovable         bool   `json:"isRemovable"`         // Is Removable
}

// String represents Datacenter as string
func (d Datacenter) String() string {
	out := fmt.Sprintf("ID: %d\n", d.ID)
	out += fmt.Sprintf("Name: %s\n", d.Name)
	out += fmt.Sprintf("Version: %s\n", d.Version)
	out += fmt.Sprintf("Description: %s\n", d.Description)
	out += fmt.Sprintf("Commercial Range Name: %s\n", d.CommercialRangeName)
	out += fmt.Sprintf("Is Removable: %t\n", d.IsRemovable)
	return out
}

// Task represents OVH dedicatedCloud.Task
type Task struct {
	ID                   int           `json:"taskId"`               // ID of the task
	Name                 string        `json:"name"`                 // Task name
	Type                 string        `json:"type"`                 // Type (maintenance|generic)
	State                string        `json:"state"`                // State
	Progress             int           `json:"progress"`             // Progress
	Description          string        `json:"description"`          // Description
	LastModificationDate govh.DateTime `json:"lastModificationDate"` // Last modification date
	ExecutionDate        govh.DateTime `json:"executionDate"`        // Execution date
	EndDate              govh.DateTime `json:"endDate"`              // End date
	MaintenanceDateFrom  govh.DateTime `json:"maintenanceDateFrom"`  // Maintenance date from
	MaintenanceDateTo    govh.DateTime `json:"maintenanceDateTo"`    // Maintenance date to
}

// String represents Task as string
func (t Task) String() string {
	endDate := ""
	if !t.EndDate.IsZero() {
		endDate = fmt.Sprintf("%s", t.EndDate)
	}
	maintenanceDateFrom := ""
	if !t.MaintenanceDateFrom.IsZero() {
		maintenanceDateFrom = fmt.Sprintf("%s", t.MaintenanceDateFrom)
	}
	maintenanceDateTo := ""
	if !t.MaintenanceDateTo.IsZero() {
		maintenanceDateTo = fmt.Sprintf("%s", t.MaintenanceDateTo)
	}
	out := fmt.Sprintf("ID: %d\n", t.ID)
	out += fmt.Sprintf("Name: %s\n", t.Name)
	out += fmt.Sprintf("Type: %s\n", t.Type)
	out += fmt.Sprintf("State: %s\n", t.State)
	out += fmt.Sprintf("Progress: %d\n", t.Progress)
	out += fmt.Sprintf("Description: %s\n", t.Description)
	out += fmt.Sprintf("Last Modification Date: %s\n", t.LastModificationDate)
	out += fmt.Sprintf("Execution Date: %s\n", t.ExecutionDate)
	out += fmt.Sprintf("End Date: %s\n", endDate)
	out += fmt.Sprintf("Maintenance Date From : %s\n", maintenanceDateFrom)
	out += fmt.Sprintf("Maintenance Date To : %s", maintenanceDateTo)
	return out
}
