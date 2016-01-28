package sms

import (
	"fmt"
	"strings"
)

// Job represesnts an sms.Job OVH
type Job struct {
	ID               int    `json:"id,omitempty"`
	Receiver         string `json:"receiver"`
	DeliveryReceipt  int    `json:"deliveryReceipt,omitempty"`
	MessageLength    int    `json:"messageLength,omitempty"`
	DifferedDelivery int    `json:"differedDelivery,omitempty"`
	Credits          int    `json:"credits,omitempty"`
	Message          string `json:"message,omitempty"`
	Ptt              int    `json:"ptt,omitempty"`
	Sender           string `json:"sender,omitempty"`
	CreationDatetime string `json:"creationDatetime,omitempty"`
	NumberOfSms      int    `json:"numberOfSms,omitempty"`
}

// SendJob represents
// Used when adding a new job
type SendJob struct {
	NoStopClause         bool     `json:"noStopClause,omitempty"`         // Do not display STOP clause in the message, this requires that this is not an advertising message
	Priority             string   `json:"priority,omitempty"`             // The priority of the message
	ValidityPeriod       int      `json:"validityPeriod,omitempty"`       // The maximum time -in minute(s)- before the message is dropped. default 2880
	SenderForResponse    bool     `json:"senderForResponse,omitempty"`    // Set the flag to send a special sms which can be reply by the receiver (smsResponse).
	Receivers            []string `json:"receivers,omitempty"`            // The receivers list
	Charset              string   `json:"charset,omitempty"`              // The sms encoding (default UTF-8)
	Coding               string   `json:"coding,omitempty"`               // Sms encodig (default 7bit)
	Message              string   `json:"message,omitempty"`              // The sms message
	DifferedPeriod       int      `json:"differedPeriod,omitempty"`       // The time -in minute(s)- to wait before sending the message
	ReceiversSlotID      string   `json:"receiversSlotId,omitempty"`      // The receivers document slot id
	Sender               string   `json:"sender,omitempty"`               // The sender
	Tag                  string   `json:"tag,omitempty"`                  // The identifier group tag
	ReceiversDocumentURL string   `json:"receiversDocumentUrl,omitempty"` // The receivers document url link in csv format
	PhoneDisplay         string   `json:"phoneDisplay,omitempty"`         // The sms class
	Class                string   `json:"class,omitempty"`                // The sms class
	ServiceName          string   `json:"serviceName,omitempty"`          // The internal name of your SMS offer
}

// SendingReport represents OVH sms.SmsSendingreport type
type SendingReport struct {
	TotalCreditsRemoved int      `json:"totalCreditsRemoved"` // Credit removed
	InvalidReceivers    []string `json:"invalidReceivers"`    // List of invalids receivers
	JobIDs              []int    `json:"ids"`                 // Id of SMS/jobs
	ValidReceivers      []string `json:"validReceivers"`      // List of valids receivers
}

// String is the stringer for SendingReport
func (s SendingReport) String() string {
	out := ""
	for _, id := range s.JobIDs {
		out += fmt.Sprintf("Job ID:%d\n", id)
	}
	out += fmt.Sprintf("Invalid receivers: %s\n", strings.Join(s.InvalidReceivers, ", "))
	out += fmt.Sprintf("Valid receivers:%s\n", strings.Join(s.ValidReceivers, ", "))
	out += fmt.Sprintf("Credits removed: %d\n", s.TotalCreditsRemoved)
	return out
}
