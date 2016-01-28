package sms

import (
	"encoding/json"
	"errors"
	//"fmt"
	"net/url"

	"github.com/toorop/govh"
	//"os"
)

// Client is a REST client for sms API
type Client struct {
	*govh.OVHClient
}

// New return a new SmsRessource
func New(client *govh.OVHClient) (*Client, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return &Client{client}, nil
}

// GetServices returns availables SMS services
func (c *Client) GetServices() (services []string, err error) {
	resp, err := c.GET("sms")
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Body, &services)
	return
}

// AddJob send a new SMS
func (c *Client) AddJob(serviceName string, job *SendJob) (report SendingReport, err error) {
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

	resp, err := c.POST("sms/"+url.QueryEscape(serviceName)+"/jobs", string(payload))
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Body, &report)
	return
}
