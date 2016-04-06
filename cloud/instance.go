package cloud

import "time"

// instanceStatus represents the status of an instance
type instanceStatus int

const (
	CREATING instanceStatus = 1 + iota
	ACTIVE
	BUILDING
	DELETED
	ERROR
	HARDREBOOT
	PASSWORD
	PAUSED
	REBOOT
	REBUILD
	RESCUED
	RESIZED
	REVERTRESIZE
	SOFTDELETED
	STOPPED
	SUSPENDED
	UNKNOWN
	VERIFYRESIZE
	MIGRATING
	RESIZE
	BUILD
	SHUTOFF
	RESCUE
	SHELVED
	SHELVEDOFFLOADED
	RESCUING
	UNRESCUING
	SNAPSHOTTING
)

func (status instanceStatus) String() string {
	switch status {
	case ACTIVE:
		return "active"
	case BUILDING:
		return "building"
	case DELETED:
		return "deleted"
	case ERROR:
		return "error"
	case HARDREBOOT:
		return "hardreboot"
	case PASSWORD:
		return "passwaord"
	case PAUSED:
		return "paused"
	case REBOOT:
		return "reboot"
	case REBUILD:
		return "rebuild"
	case RESCUED:
		return "rescued"
	case RESIZED:
		return "resized"
	case REVERTRESIZE:
		return "revertresize"
	case SOFTDELETED:
		return "softdeleted"
	case STOPPED:
		return "stopped"
	case SUSPENDED:
		return "suspended"
	case UNKNOWN:
		return " unknown"
	case VERIFYRESIZE:
		return "verifsize"
	case MIGRATING:
		return "migrating"
	case RESIZE:
		return "resize"
	case BUILD:
		return "build"
	case SHUTOFF:
		return "shutoff"
	case RESCUE:
		return "rescue"
	case SHELVED:
		return "shelved"
	case SHELVEDOFFLOADED:
		return "shelvedoffloaded"
	case RESCUING:
		return "rescuing"
	case UNRESCUING:
		return "unrescuing"
	case SNAPSHOTTING:
		return "snapshotting"
	}
	return "unknow"
}

// Instance represents an instance
type Instance struct {
	Status   instanceStatus
	Name     string
	Region   string
	ImageID  string
	Created  time.Time
	FlavorID string
	SSHKeyID string
}

// MonthlyBilling Instance monthly billing status
type MonthlyBilling struct {
	Since  time.Time
	Status MonthlyBillingStatusEnum
}

// MonthlyBillingStatusEnum Monthly billing status
type MonthlyBillingStatusEnum int

const (
	ActivationPending MonthlyBillingStatusEnum = 1 + iota
	OK
)

func (status MonthlyBillingStatusEnum) String() string {
	switch status {
	case ActivationPending:
		return "Activation pending"
	case OK:
		return "ok"
	}
	return "unknow"
}
