package domain

import (
	"fmt"

	"github.com/toorop/govh"
)

// FieldTypes represents supported type for a DNS record
var FieldTypes = [14]string{"A", "AAAA", "CNAME", "DKIM", "LOC", "MX", "NAPTR", "NS", "PTR", "SPF", "SRV", "SSHFP", "TLSA", "TXT"}

// IsValidFieldType check if t is a valid type
func IsValidFieldType(t string) bool {
	//type = strings.ToUpper(type)
	for _, validType := range FieldTypes {
		if t == validType {
			return true
		}
	}
	return true
}

// ZoneRecord represents OVH domain.zone.Record type
type ZoneRecord struct {
	ID        int    `json:"id"`
	Zone      string `json:"zone"`
	TTL       int    `json:"ttl"`
	Target    string `json:"target"`
	FieldType string `json:"fieldType"`
	SubDomain string `json:"subDomain"`
}

// String return ZoneRecord as string
func (r ZoneRecord) String() string {
	//123345 toorop.fr A 300 Target
	out := fmt.Sprintf("%d ", r.ID)
	if r.SubDomain != "" {
		out += r.SubDomain + "."
	}
	out += fmt.Sprintf("%s %s %d %s", r.Zone, r.FieldType, r.TTL, r.Target)
	return out
}

// ZoneTask represents an OVH domain.zone.Task type
type ZoneTask struct {
	ID            int           `json:"id"`
	Function      string        `json:"function"`
	Status        string        `json:"status"`
	CanAccelerate bool          `json:"canAccelerate"`
	LastUpdate    govh.DateTime `json:"lastUpdate"`
	CreationDate  govh.DateTime `json:"creationDate"`
	DoneDate      govh.DateTime `json:"doneDate"`
	TodoDate      govh.DateTime `json:"todoDate"`
	Comment       string        `json:"comment"`
	CanCancel     bool          `json:"canCancel"`
	CanRelaunch   bool          `json:"canRelaunch"`
}

// Stringer
func (z ZoneTask) String() string {
	out := fmt.Sprintf("ID: %d\n", z.ID)
	out += "Function: " + z.Function + "\n"
	out += "Status: " + z.Status + "\n"
	out += "Status: " + z.Comment + "\n"
	out += fmt.Sprintf("Created at: %s\n", z.CreationDate)
	out += fmt.Sprintf("Updated at: %s\n", z.LastUpdate)
	out += fmt.Sprintf("Todo at: %s\n", z.TodoDate)
	out += fmt.Sprintf("Done at: %s\n", z.DoneDate)
	out += fmt.Sprintf("Can cancel: %t\n", z.CanCancel)
	out += fmt.Sprintf("Can accelerate: %t\n", z.CanAccelerate)
	out += fmt.Sprintf("Can relaunch: %t\n", z.CanRelaunch)
	return out
}
