package govh

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// OVHClient API client
type OVHClient struct {
	ak string
	as string
	ck string
	//api_endpoint string
	endpoint string
	client   *http.Client
}

// ovhResponseErr represents an unmarshalled reponse from OVH in case of error
type ovhResponseErr struct {
	ErrorCode      string `json:"errorCode"`
	HTTPStatusCode string `json:"httpCode"`
	Message        string `json:"message"`
}

// NewClient returns an OVH API Client
func NewClient(ak string, as string, ck string, region string) (c *OVHClient) {
	endpoint := API_ENDPOINT_EU
	if strings.ToLower(region) == "ca" {
		endpoint = API_ENDPOINT_CA
	}
	return &OVHClient{ak, as, ck, endpoint, &http.Client{}}

}

// Response represents a response from OVH API
type Response struct {
	StatusCode int
	Status     string
	Body       []byte
}

// HandleErr return error on unexpected HTTP code
func (r *Response) HandleErr(err error, expectedHTTPCode []int) error {
	if err != nil {
		return err
	}
	for _, code := range expectedHTTPCode {
		if r.StatusCode == code {
			return nil
		}
	}
	// Try to get OVH returning info about the error
	if r.Body != nil {
		var ovhResponse ovhResponseErr
		err := json.Unmarshal(r.Body, &ovhResponse)
		if err == nil {
			if len(ovhResponse.ErrorCode) != 0 {
				return errors.New(ovhResponse.ErrorCode)
			} else {
				return errors.New(ovhResponse.Message)
			}
		}
	}
	return fmt.Errorf("%d - %s", r.StatusCode, r.Status)
}

// Do process the request & return a reponse (or error)
func (c *OVHClient) Do(method string, ressource string, payload string) (response Response, err error) {
	query := fmt.Sprintf("%s/%s/%s", c.endpoint, API_VERSION, ressource)
	req, err := http.NewRequest(method, query, strings.NewReader(payload))
	if err != nil {
		return
	}
	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/json;charset=utf-8")
	}
	req.Header.Add("Accept", "application/json")

	timestamp := fmt.Sprintf("%d", int32(time.Now().Unix()))
	req.Header.Add("X-Ovh-Timestamp", timestamp)
	req.Header.Add("X-Ovh-Application", c.ak)
	req.Header.Add("X-Ovh-Consumer", c.ck)
	p := strings.Split(ressource, "?")
	req.URL.Opaque = fmt.Sprintf("/%s/%s", API_VERSION, p[0])

	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%s+%s+%s+%s+%s+%s", c.as, c.ck, method, query, payload, timestamp)))
	req.Header.Add("X-Ovh-Signature", fmt.Sprintf("$1$%x", h.Sum(nil)))

	//r, err := c.client.Do(req)
	r, err := c.doTimeoutRequest(time.NewTimer(30*time.Second), req)
	if err != nil {
		return
	}
	defer r.Body.Close()

	response.StatusCode = r.StatusCode
	response.Status = r.Status
	response.Body, err = ioutil.ReadAll(r.Body)
	return
}

// doTimeoutRequest do a HTTP request with timeout
func (c *OVHClient) doTimeoutRequest(timer *time.Timer, req *http.Request) (*http.Response, error) {
	// Do the request in the background so we can check the timeout
	type result struct {
		resp *http.Response
		err  error
	}
	done := make(chan result, 1)
	go func() {
		resp, err := c.client.Do(req)
		done <- result{resp, err}
	}()
	// Wait for the read or the timeout
	select {
	case r := <-done:
		return r.resp, r.err
	case <-timer.C:
		return nil, errors.New("timeout on reading data from OVH API")
	}
}
