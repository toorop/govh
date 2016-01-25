package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"sync"

	"github.com/toorop/govh"
)

// Client is an OVH API client
type Client struct {
	*govh.OVHClient
}

// New return a new Client
func New(client *govh.OVHClient) (*Client, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return &Client{client}, nil
}

// List return a list of domains
func (c *Client) List(whoisOwner ...string) (domains []string, err error) {
	uri := "domain"
	if len(whoisOwner) != 0 {
		uri += "?whoisOwner=" + url.QueryEscape(strings.Join(whoisOwner, ""))
	}
	r, err := c.GET(uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &domains)
	return
}

// GetRecordsOptions options for Client.GetRecordIDs
type GetRecordsOptions struct {
	FieldType string
	SubDomain string
}

// GetRecordIDs return record ID for the zone zone
func (c *Client) GetRecordIDs(zone string, options GetRecordsOptions) (IDs []int, err error) {
	uri := "domain/zone/" + url.QueryEscape(strings.ToLower(zone)) + "/record"
	v := url.Values{}
	if options.FieldType != "" {
		options.FieldType = strings.ToUpper(options.FieldType)
		if !IsValidFieldType(options.FieldType) {
			return IDs, fmt.Errorf("%s is not a valid type", options.FieldType)
		}
		v.Add("fieldType", options.FieldType)
	}
	if options.SubDomain != "" {
		v.Add("subDomain", strings.ToLower(options.SubDomain))
	}
	params := v.Encode()
	if params != "" {
		uri += "?" + params
	}
	r, err := c.GET(uri)
	if err != nil {
		return
	}
	if string(r.Body) != "" {
		err = json.Unmarshal(r.Body, &IDs)
	}
	return
}

// GetRecordByID return a ZoneRecord by its ID
func (c *Client) GetRecordByID(zone string, ID int) (record ZoneRecord, err error) {
	record = ZoneRecord{}
	r, err := c.GET("domain/zone/" + url.QueryEscape(strings.ToLower(zone)) + "/record/" + fmt.Sprintf("%d", ID))
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Body, &record)
	return
}

// GetRecords returns record(s) for zone filtered by filedType
func (c *Client) GetRecords(zone string, options GetRecordsOptions) (records []ZoneRecord, err error) {
	IDs, err := c.GetRecordIDs(zone, options)
	if err != nil {
		return
	}
	var wg sync.WaitGroup
	errChan := make(chan error, 1)
	doneChan := make(chan int)

	for _, ID := range IDs {
		//log.Println("range", ID)
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			record, err := c.GetRecordByID(zone, id)
			if err != nil {
				errChan <- err
				return
			}
			records = append(records, record)
		}(ID)
	}

	go func() {
		wg.Wait()
		doneChan <- 1
	}()

	select {
	case err = <-errChan:
		return []ZoneRecord{}, err

	case <-doneChan:
		break
	}
	return
}
