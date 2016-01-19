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
	*http.Client
	ak       string
	as       string
	ck       string
	endpoint string
}

// New returns an OVH API Client
func New(ak string, as string, ck string, region string) (c *OVHClient) {
	endpoint := API_ENDPOINT_EU
	if strings.ToLower(region) == "ca" {
		endpoint = API_ENDPOINT_CA
	}
	return &OVHClient{&http.Client{}, ak, as, ck, endpoint}
}

// ovhResponseErr represents a response from OVH in case of error
type responseERR struct {
	ErrorCode string `json:"errorCode"`
	HTTPCode  string `json:"httpCode"`
	Message   string `json:"message"`
}

// APIResponse represents a response from OVH API
type APIResponse struct {
	StatusCode int
	Status     string
	Body       []byte
}

// HandleErr return error on unexpected HTTP code
func (r *APIResponse) HandleErr(err error, expectedHTTPCode []int) error {
	if err != nil {
		return err
	}
	for _, code := range expectedHTTPCode {
		if r.StatusCode == code {
			return nil
		}
	}
	// Try to get OVH response about the error
	if r.Body != nil {
		var ovhResponse responseERR
		err := json.Unmarshal(r.Body, &ovhResponse)
		if err == nil {
			return errors.New(ovhResponse.HTTPCode + ovhResponse.ErrorCode + ovhResponse.Message)
		}
	}
	return fmt.Errorf("%d - %s", r.StatusCode, r.Status)
}

// GET do a GET query
func (c *OVHClient) GET(ressource string, expectedHTTPCode ...int) (response APIResponse, err error) {
	return c.Query("GET", ressource, "", expectedHTTPCode)
}

// POST do a POST query
func (c *OVHClient) POST(ressource, payload string, expectedHTTPCode ...int) (APIresponse APIResponse, err error) {
	return c.Query("POST", ressource, payload, expectedHTTPCode)
}

// PUT do a PUT query
func (c *OVHClient) PUT(ressource, payload string, expectedHTTPCode ...int) (response APIResponse, err error) {
	return c.Query("PUT", ressource, payload, expectedHTTPCode)
}

// DELETE do a GET query
func (c *OVHClient) DELETE(ressource string, expectedHTTPCode ...int) (response APIResponse, err error) {
	return c.Query("DELETE", ressource, "", expectedHTTPCode)
}

// Query process the request & return a response (or error)
func (c *OVHClient) Query(method string, ressource string, payload string, expectedHTTPCode []int) (response APIResponse, err error) {
	if len(expectedHTTPCode) == 0 {
		expectedHTTPCode = []int{200}
	}
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

	r, err := c.doTimeoutRequest(time.NewTimer(30*time.Second), req)
	if err != nil {
		return
	}
	defer r.Body.Close()
	response.StatusCode = r.StatusCode
	response.Status = r.Status
	response.Body, err = ioutil.ReadAll(r.Body)
	for _, code := range expectedHTTPCode {
		if response.StatusCode == code {
			return
		}
	}
	// Try to get OVH response about the error
	if response.Body != nil {
		var ovhResponse responseERR
		err = json.Unmarshal(response.Body, &ovhResponse)
		if err == nil {
			return response, errors.New(ovhResponse.HTTPCode + ovhResponse.ErrorCode + ovhResponse.Message)
		}
	}
	return response, fmt.Errorf("%d - %s", response.StatusCode, response.Status)
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
		resp, err := c.Do(req)
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
