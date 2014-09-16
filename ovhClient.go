package govh

import (
	"crypto/sha1"
	"fmt"
	//"io"
	"io/ioutil"
	//"log"
	"errors"
	"net/http"
	//"net/url"
	//"os"
	"encoding/json"
	"strings"
	"time"
)

type OvhClient struct {
	ak           string
	as           string
	ck           string
	api_endpoint string
	client       *http.Client
}

// ovhResponseErr represents an unmarshalled reponse from OVH in case od error
type ovhResponseErr struct {
	ErrorCode string `json:"errorCode"`
	HttpCode  string `json:"httpCode"`
	Message   string `json:"message"`
}

func NewClient(ak string, as string, ck string, region string) (c *OvhClient) {
	endpoint := API_ENDPOINT_EU
	if strings.ToLower(region) == "ca" {
		endpoint = API_ENDPOINT_CA
	}
	return &OvhClient{ak, as, ck, endpoint, &http.Client{}}

}

// Response represents an response from OVH API
type response struct {
	StatusCode int
	Status     string
	Body       []byte
}

// handleCommon return error on unexpected HTTP code
func (r *response) HandleErr(err error, expectedHttpCode []int) error {
	if err != nil {
		return err
	}
	for _, code := range expectedHttpCode {
		if r.StatusCode == code {
			return nil
		}
	}
	// Try to get OVH response
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
	return errors.New(fmt.Sprintf("%d - %s", r.StatusCode, r.Status))
}

// Do process the request & return a reponse (or error)
func (c *OvhClient) Do(method string, ressource string, payload string) (response response, err error) {
	query := fmt.Sprintf("%s/%s/%s", c.api_endpoint, API_VERSION, ressource)
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

	r, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer r.Body.Close()

	response.StatusCode = r.StatusCode
	response.Status = r.Status
	response.Body, err = ioutil.ReadAll(r.Body)
	return
}
