package ovh

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Keyring struct {
	AppKey      string
	AppSecret   string
	ConsumerKey string
}

func AuthGetConsumerKey(k Keyring) (ck string, link string, err error) {
	type response struct {
		ValidationUrl string
		ConsumerKey   string
		State         string
	}

	var jresp response

	if len(k.AppSecret) == 0 {
		err = errors.New("Application Secret not found in your keyring")
		return
	}
	client := &http.Client{}

	body := "{\"accessRules\":[{\"method\":\"GET\",\"path\":\"/*\"},{\"method\":\"GET\",\"path\":\"/*\"},{\"method\":\"DELETE\",\"path\":\"/*\"},{\"method\":\"PUT\",\"path\":\"/*\"},{\"method\":\"GET\",\"path\":\"/*\"} ]}"
	url := fmt.Sprintf("%s/auth/credential", API_URI)

	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	req.Header.Add("User-Agent", "Govh (https://github.com/Toorop/govh)")
	req.Header.Add("X-Ovh-Application", k.AppKey)
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
