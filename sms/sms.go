package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Toorop/govh"
	//"net/url"
	//"os"
)

type SmsRessource struct {
	client *govh.OvhClient
}

func New(client *govh.OvhClient) (*SmsRessource, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return &SmsRessource{client: client}, nil
}

// List SMS services availables
func (r *SmsRessource) ListServices() (services []string, err error) {
	t, err := r.client.Do("GET", "sms", "")
	if err != nil {
		return
	}
	if err = json.Unmarshal(t, &services); err != nil {
		return nil, err
	}
	return
}

// Add a new SMS job
// ie send a SMS
func (r *SmsRessource) AddJob(serviceName string, job *NewJob) (report SendingReport, err error) {
	// Default value
	if job.ValidityPeriod == 0 {
		job.ValidityPeriod = 2880
	}
	if job.Class == "" {
		job.Class = "sim"
	}
	uri := fmt.Sprintf("sms/%s/jobs", serviceName)
	payload, err := json.Marshal(job)
	if err != nil {
		return
	}

	t, err := r.client.Do("POST", uri, string(payload))
	if err != nil {
		err = errors.New(fmt.Sprintf("%s - %s", err, t))
	}
	err = json.Unmarshal(t, &report)
	if err != nil {
		err = errors.New(fmt.Sprintf("%s - %s", err, t))
	}
	return
}
