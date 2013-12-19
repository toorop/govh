package govh

import (
	"crypto/sha1"
	"fmt"
	//"io"
	"io/ioutil"
	//"log"
	"errors"
	"net/http"
	"strings"
	"time"
)

const (
	API_BASE = "https://api.ovh.com/1.0"
)

type OvhClient struct {
	ak     string
	as     string
	ck     string
	client *http.Client
}

func NewClient(ak string, as string, ck string) (c *OvhClient) {
	return &OvhClient{ak, as, ck, &http.Client{}}

}

func (c *OvhClient) Do(method string, ressource string, payload string) (response string, err error) {
	query := fmt.Sprintf("%s/%s", API_BASE, ressource)
	req, err := http.NewRequest(method, query, strings.NewReader(payload))
	if err != nil {
		return
	}

	timestamp := fmt.Sprintf("%d", int32(time.Now().Unix()))
	req.Header.Add("X-Ovh-Timestamp", timestamp)
	req.Header.Add("X-Ovh-Application", c.ak)
	req.Header.Add("X-Ovh-Consumer", c.ck)
	// Signature
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%s+%s+%s+%s+%s+%s", c.as, c.ck, "GET", query, payload, timestamp)))
	req.Header.Add("X-Ovh-Signature", fmt.Sprintf("$1$%x", h.Sum(nil)))

	resp, err := c.client.Do(req)
	if err != nil {
		return
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return response, err
		}
		if resp.StatusCode != 200 {
			err = errors.New(resp.Status)
		}
		response = string(contents)
		return response, err
	}
	return

}
