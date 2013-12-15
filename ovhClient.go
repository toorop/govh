package govh

import (
	"crypto/sha1"
	"fmt"
	//"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	API_BASE = "https://api.ovh.com/1.0"
)

type ovhClient struct {
	ak     string
	as     string
	ck     string
	client *http.Client
}

func NewClient(ak string, as string, ck string) (c *ovhClient) {
	return &ovhClient{ak, as, ck, &http.Client{}}

}

func (c *ovhClient) Do(method string, ressource string, payload string) (response string, err error) {
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
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		/*fmt.Println("The calculated length is:", len(string(contents)))
		fmt.Println("   ", resp.StatusCode)
		hdr := resp.Header
		for key, value := range hdr {
			fmt.Println("   ", key, ":", value)
		}*/
		response = string(contents)
	}
	return

}
