package sms

// Job
type Job struct {
	Receiver string `json:"receiver"`

	DeliveryReceipt  int    `json:"deliveryReceipt,omitempty"`
	MessageLength    int    `json:"messageLength,omitempty"`
	DifferedDelivery int    `json:"differedDelivery,omitempty"`
	Credits          int    `json:"credits,omitempty"`
	Message          string `json:"message,omitempty"`
	Ptt              int    `json:"ptt,omitempty"`
	Sender           string `json:"sender,omitempty"`
	CreationDatetime string `json:"creationDatetime,omitempty"`
	numberOfSms      int    `json:"numberOfSms,omitempty"`
	Id               int    `json:"id,omitempty"`
}

// NewJob
// Used when addinf a new job
type NewJob struct {
	NoStopClause         bool     `json:"noStopClause,omitempty"`         // Do not display STOP clause in the message, this requires that this is not an advertising message
	Priority             string   `json:"priority,omitempty"`             // The priority of the message
	ValidityPeriod       int      `json:"validityPeriod,omitempty"`       // The maximum time -in minute(s)- before the message is dropped. default 2880
	SenderForResponse    bool     `json:"senderForResponse,omitempty"`    // Set the flag to send a special sms which can be reply by the receiver (smsResponse).
	Receivers            []string `json:"receivers,omitempty"`            // The receivers list
	Charset              string   `json:"charset,omitempty"`              // The sms encoding (default UTF-8)
	Coding               string   `json:"coding,omitempty"`               // Sms encodig (default 7bit)
	Message              string   `json:"message,omitempty"`              // The sms message
	DifferedPeriod       int      `json:"differedPeriod,omitempty"`       // The time -in minute(s)- to wait before sending the message
	ReceiversSlotId      string   `json:"receiversSlotId,omitempty"`      // The receivers document slot id
	Sender               string   `json:"sender,omitempty"`               // The sender
	Tag                  string   `json:"tag,omitempty"`                  // The identifier group tag
	ReceiversDocumentUrl string   `json:"receiversDocumentUrl,omitempty"` // The receivers document url link in csv format
	PhoneDisplay         string   `json:"phoneDisplay,omitempty"`         // The sms class
	Class                string   `json:"class,omitempty"`                // The sms class
	serviceName          string   `json:"serviceName,omitempty"`          // The internal name of your SMS offer
}

// Sending Report
type SendingReport struct {
	TotalCreditsRemoved int   `json:"totalCreditsRemoved,omitempty"` // Credit removed
	Ids                 []int `json:"ids,omitempty"`                 // Id of SMS/jobs
}
