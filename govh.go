package govh

import (
	//"net/http"
	"strings"
)

const (
	//API_HOST        = "api.ovh.com"
	API_ENDPOINT_EU = "https://api.ovh.com"
	API_ENDPOINT_CA = "https:///ca.api.ovh.com"
	API_ENDPOINT_SYS_EU = "https://eu.api.soyoustart.com"
	API_ENDPOINT_SYS_CA = "https://ca.api.soyoustart.com"
	API_VERSION     = "1.0"
)

func RegionEndpoint(region string) string {
	endpoint := API_ENDPOINT_EU
	switch strings.ToLower(region) {
		case "sys_ca":
			endpoint = API_ENDPOINT_SYS_CA
			break
		case "sys_eu":
			endpoint = API_ENDPOINT_SYS_EU
			break
		case "ca":
			endpoint = API_ENDPOINT_CA
			break
		case "eu":
			break
	}
	return endpoint
}
