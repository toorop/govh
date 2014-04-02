package govh

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func AuthGetConsumerKey(ak string) (ck string, link string, err error) {
	type response struct {
		ValidationUrl string
		ConsumerKey   string
		State         string
	}

	var jresp response

	client := &http.Client{}

	body := "{\"accessRules\":[{\"method\":\"GET\",\"path\":\"/*\"},{\"method\":\"POST\",\"path\":\"/*\"},{\"method\":\"DELETE\",\"path\":\"/*\"},{\"method\":\"PUT\",\"path\":\"/*\"},{\"method\":\"DELETE\",\"path\":\"/*\"} ]}"
	url := fmt.Sprintf("%s/%s/auth/credential", API_BASE, API_VERSION)

	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	req.Header.Add("User-Agent", "Govh (https://github.com/Toorop/govh)")
	req.Header.Add("X-Ovh-Application", ak)
	req.Header.Add("Content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	// Bad HTTP status
	if resp.StatusCode > 399 {
		err = errors.New(resp.Status)
		return
	}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(r, &jresp)
	if err != nil {
		return
	}
	resp.Body.Close()
	ck = jresp.ConsumerKey
	link = jresp.ValidationUrl
	return
}
