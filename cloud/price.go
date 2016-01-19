package cloud

import (
	"encoding/json"

	"github.com/toorop/govh/order"
)

// VolumePrice is Prices for volumes
type VolumePrice struct {
	VolumeName   string      `json:"volumeName"`
	Region       string      `json:"region"`
	Price        order.Price `json:"price"`
	MonthlyPrice order.Price `json:"monthlyPrice"`
}

// FlavourPrice is price for flavour
type FlavourPrice struct {
	FlavorName   string      `json:"flavorName"`
	FlavorID     string      `json:"flavorId"`
	Region       string      `json:"region"`
	MonthlyPrice order.Price `json:"monthlyPrice"`
	Price        order.Price `json:"price"`
}

// SnapshotPrice is prices for snapshot
type SnapshotPrice struct {
	Region       string      `json:"region"`
	MonthlyPrice order.Price `json:"monthlyPrice"`
	Price        order.Price `json:"price"`
}

// GetPriceResponse represents cloud.Price as returned by OVH
type GetPriceResponse struct {
	Volumes         []VolumePrice   `json:"volumes"`
	ProjectCreation order.Price     `json"projectCreation"`
	Instances       []FlavourPrice  `json:"instances"`
	Snapshots       []SnapshotPrice `json:"snapshots"`
}

// String return string representation of GetPriceResponse
func (r *GetPriceResponse) String() string {
	s := ""
	// Volumes
	s += "Volumes (per GB)"
	for _, v := range r.Volumes {
		s += "\n\t" + v.VolumeName + " @ " + v.Region + ":\n\t\t" + v.Price.Text + " " + v.Price.CurrencyCode + " hourly\n\t\t" + v.MonthlyPrice.Text + " " + v.MonthlyPrice.CurrencyCode + " monthly"
	}

	// Project Creation
	s += "\nProject Creation\n\t" + r.ProjectCreation.Text + " " + r.ProjectCreation.CurrencyCode

	// Instances
	s += "\nInstances"
	for _, i := range r.Instances {
		s += "\n\t" + i.FlavorName + " @ " + i.Region + ":\n\t\t" + i.Price.Text + " " + i.Price.CurrencyCode + " hourly\n\t\t" + i.MonthlyPrice.Text + " " + i.MonthlyPrice.CurrencyCode + " monthly"
	}

	//snapshots
	s += "\nSnapshots (per GB)"
	for _, r := range r.Snapshots {
		s += "\n\t @ " + r.Region + ":\n\t\t" + r.Price.Text + " " + r.Price.CurrencyCode + " hourly\n\t\t" + r.MonthlyPrice.Text + " " + r.MonthlyPrice.CurrencyCode + " monthly"

	}

	return s
}

// JSON return a JSON representation of GetPriceResponse
func (r *GetPriceResponse) JSON() string {
	s, _ := json.Marshal(r)
	return string(s)
}
