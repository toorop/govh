package sms

import (
	"encoding/json"
	"errors"
	//"fmt"
	"github.com/toorop/govh"
	"net/url"
	//"os"
)

type SmsRessource struct {
	client *govh.OvhClient
}

// New return a new SmsRessource
func New(client *govh.OvhClient) (*SmsRessource, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return &SmsRessource{client: client}, nil
}

// ListServices returns availables SMS services
func (r *SmsRessource) ListServices() (services []string, err error) {
	resp, err := r.client.Do("GET", "sms", "")
	if err = resp.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(resp.Body, &services)
	return
}

// AddJob send a new SMS
func (r *SmsRessource) AddJob(serviceName string, job *SendJob) (report SendingReport, err error) {
	// Default value
	if job.ValidityPeriod == 0 {
		job.ValidityPeriod = 2880
	}
	if job.Class == "" {
		job.Class = "sim"
	}

	payload, err := json.Marshal(job)
	if err != nil {
		return
	}

	resp, err := r.client.Do("POST", "sms/"+url.QueryEscape(serviceName)+"/jobs", string(payload))
	if err = resp.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(resp.Body, &report)
	return
}
