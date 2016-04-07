package cloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// instanceStatus represents the status of an instance
type instanceStatus int

const (
	InstanceStatusCREATING instanceStatus = 1 + iota
	InstanceStatusACTIVE
	InstanceStatusBUILDING
	InstanceStatusDELETED
	InstanceStatusERROR
	InstanceStatusHARDREBOOT
	InstanceStatusPASSWORD
	InstanceStatusPAUSED
	InstanceStatusREBOOT
	InstanceStatusREBUILD
	InstanceStatusRESCUED
	InstanceStatusRESIZED
	InstanceStatusREVERTRESIZE
	InstanceStatusSOFTDELETED
	InstanceStatusSTOPPED
	InstanceStatusSUSPENDED
	InstanceStatusUNKNOWN
	InstanceStatusVERIFYRESIZE
	InstanceStatusMIGRATING
	InstanceStatusRESIZE
	InstanceStatusBUILD
	InstanceStatusSHUTOFF
	InstanceStatusRESCUE
	InstanceStatusSHELVED
	InstanceStatusSHELVEDOFFLOADED
	InstanceStatusRESCUING
	InstanceStatusUNRESCUING
	InstanceStatusSNAPSHOTTING
)

func (status instanceStatus) String() string {
	switch status {
	case InstanceStatusACTIVE:
		return "active"
	case InstanceStatusBUILDING:
		return "building"
	case InstanceStatusDELETED:
		return "deleted"
	case InstanceStatusERROR:
		return "error"
	case InstanceStatusHARDREBOOT:
		return "hardreboot"
	case InstanceStatusPASSWORD:
		return "passwaord"
	case InstanceStatusPAUSED:
		return "paused"
	case InstanceStatusREBOOT:
		return "reboot"
	case InstanceStatusREBUILD:
		return "rebuild"
	case InstanceStatusRESCUED:
		return "rescued"
	case InstanceStatusRESIZED:
		return "resized"
	case InstanceStatusREVERTRESIZE:
		return "revertresize"
	case InstanceStatusSOFTDELETED:
		return "softdeleted"
	case InstanceStatusSTOPPED:
		return "stopped"
	case InstanceStatusSUSPENDED:
		return "suspended"
	case InstanceStatusUNKNOWN:
		return " unknown"
	case InstanceStatusVERIFYRESIZE:
		return "verifsize"
	case InstanceStatusMIGRATING:
		return "migrating"
	case InstanceStatusRESIZE:
		return "resize"
	case InstanceStatusBUILD:
		return "build"
	case InstanceStatusSHUTOFF:
		return "shutoff"
	case InstanceStatusRESCUE:
		return "rescue"
	case InstanceStatusSHELVED:
		return "shelved"
	case InstanceStatusSHELVEDOFFLOADED:
		return "shelvedoffloaded"
	case InstanceStatusRESCUING:
		return "rescuing"
	case InstanceStatusUNRESCUING:
		return "unrescuing"
	case InstanceStatusSNAPSHOTTING:
		return "snapshotting"
	}
	return "unknow"
}

// Instance represents an instance
type Instance struct {
	ID             string
	Status         instanceStatus
	Name           string
	Region         string
	ImageID        string
	Created        time.Time
	FlavorID       string
	SSHKeyID       string
	MonthlyBilling InstanceMonthlyBilling
	IPAddresses    []InstanceIPAddress
}

// String is a stringer for Instance
func (i Instance) String() string {
	s := "ID: " + i.ID
	s += "\nName: " + i.Name
	s += fmt.Sprintf("\nCreated: %s", i.Created)
	s += fmt.Sprintf("\nStatus: %s", i.Status)
	s += "\nRegion: " + i.Region
	s += "\nImage ID: " + i.ImageID
	s += "\nFlavor: " + i.FlavorID
	s += "\nSSH Key ID: " + i.SSHKeyID
	s += fmt.Sprintf("\nMonthly billing: %s", i.MonthlyBilling)
	s += "\nIP addresse(s):"
	for _, IP := range i.IPAddresses {
		s += fmt.Sprintf("\n\t%s", IP)
	}
	s += "\n"
	return s
}

// UnmarshalJSON is an unmarshaller
func (i *Instance) UnmarshalJSON(data []byte) (err error) {
	type resp struct {
		ID             string                 `json:"id"`
		Status         string                 `json:"status"`
		Name           string                 `json:"name"`
		Region         string                 `json:"region"`
		ImageID        string                 `json:"imageId"`
		Created        string                 `json:"created"`
		FlavorID       string                 `json:"flavorId"`
		SSHKeyID       string                 `json:"sshkeyId"`
		MonthlyBilling InstanceMonthlyBilling `json:"monthlyBilling"`
		IPAddresses    []InstanceIPAddress    `json:"ipAddresses"`
	}

	rp := new(resp)
	if err := json.Unmarshal(data, &rp); err != nil {
		return err
	}

	// ID
	i.ID = rp.ID

	// status
	switch rp.Status {
	case "ACTIVE":
		i.Status = InstanceStatusACTIVE
	case "BUILDING":
		i.Status = InstanceStatusBUILDING
	case "DELETED":
		i.Status = InstanceStatusDELETED
	case "ERROR":
		i.Status = InstanceStatusERROR
	case "HARD_REBOOT":
		i.Status = InstanceStatusHARDREBOOT
	case "PASSWORD":
		i.Status = InstanceStatusPASSWORD
	case "PAUSED":
		i.Status = InstanceStatusPAUSED
	case "REBOOT":
		i.Status = InstanceStatusREBOOT
	case "REBUILD":
		i.Status = InstanceStatusREBUILD
	case "RESCUED":
		i.Status = InstanceStatusRESCUE
	case "RESIZED":
		i.Status = InstanceStatusRESIZED
	case "REVERT_RESIZE":
		i.Status = InstanceStatusREVERTRESIZE
	case "SOFT_DELETED":
		i.Status = InstanceStatusSOFTDELETED
	case "STOPPED":
		i.Status = InstanceStatusSTOPPED
	case "SUSPENDED":
		i.Status = InstanceStatusSUSPENDED
	case "UNKNOWN":
		i.Status = InstanceStatusUNKNOWN
	case "VERIFYRESIZE":
		i.Status = InstanceStatusVERIFYRESIZE
	case "MIGRATING":
		i.Status = InstanceStatusMIGRATING
	case "RESIZE":
		i.Status = InstanceStatusRESIZE
	case "BUILD":
		i.Status = InstanceStatusBUILD
	case "SHUTOFF":
		i.Status = InstanceStatusSHUTOFF
	case "RESCUE":
		i.Status = InstanceStatusRESCUE
	case "SHELVED":
		i.Status = InstanceStatusSHELVED
	case "SHELVED_OFFLOADED":
		i.Status = InstanceStatusSHELVEDOFFLOADED
	default:
		return errors.New("unknow instance status: " + rp.Status)
	}

	// Name
	i.Name = rp.Name

	// Region
	i.Region = rp.Region

	// ImageID
	i.ImageID = rp.ImageID

	// Created

	// FlavorID
	i.FlavorID = rp.FlavorID

	// sshkeyId
	i.SSHKeyID = rp.SSHKeyID

	// monthlyBilling
	type mb struct {
	}

	i.MonthlyBilling = rp.MonthlyBilling

	// ipAddresses
	i.IPAddresses = rp.IPAddresses

	return nil
}

// InstanceMonthlyBilling Instance monthly billing status
type InstanceMonthlyBilling struct {
	Since  time.Time
	Status InstanceMonthlyBillingStatusEnum
}

// String is the stringer for InstanceMonthlyBilling
func (i InstanceMonthlyBilling) String() string {
	return fmt.Sprintf("Since: %s Status: %s", i.Since, i.Status)
}

// UnmarshalJSON unmarshaller
func (mb *InstanceMonthlyBilling) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	type imbs struct {
		Since  string `json:"since"`
		Status string `json:"status"`
	}
	imb := new(imbs)
	err := json.Unmarshal(data, &imb)
	// Since
	mb.Since, err = time.Parse(time.RFC3339, imb.Since)
	if err != nil {
		return err
	}
	// Status
	switch imb.Status {
	case "activationPending":
		mb.Status = InstanceMonthlyBillingStatusActivationPending
	case "ok":
		mb.Status = InstanceMonthlyBillingStatusOK
	default:
		return errors.New("unknow billing status: " + imb.Status)
	}
	return nil
}

// InstanceMonthlyBillingStatusEnum Monthly billing status
type InstanceMonthlyBillingStatusEnum int

const (
	InstanceMonthlyBillingStatusActivationPending InstanceMonthlyBillingStatusEnum = 1 + iota
	InstanceMonthlyBillingStatusOK
)

// InstanceMonthlyBillingStatusEnum stringer
func (status InstanceMonthlyBillingStatusEnum) String() string {
	switch status {
	case InstanceMonthlyBillingStatusActivationPending:
		return "Activation pending"
	case InstanceMonthlyBillingStatusOK:
		return "ok"
	}
	return "unknow"
}

// InstanceIPAddress IP address of instance
type InstanceIPAddress struct {
	IP   string `json:"ip"`
	Type string `json:"type"`
}

// String is the stringer for InstanceIPAddress
func (i InstanceIPAddress) String() string {
	return fmt.Sprintf("IP: %s Type: %s", i.IP, i.Type)
}
