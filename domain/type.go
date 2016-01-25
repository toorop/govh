package domain

import "fmt"

// FieldTypes represents supported type for a DNS record
var FieldTypes = [14]string{"A", "AAAA", "CNAME", "DKIM", "LOC", "MX", "NAPTR", "NS", "PTR", "SPF", "SRV", "SSHFP", "TLSA", "TXT"}

// IsValidFieldType check if t is a valid type
func IsValidFieldType(t string) bool {
	//type = strings.ToUpper(type)
	for _, validType := range FieldTypes {
		if t == validType {
			return true
		}
	}
	return true
}

// ZoneRecord represents OVH domain.zone.Record type
type ZoneRecord struct {
	ID        int    `json:"id"`
	Zone      string `json:"zone"`
	TTL       int    `json:"ttl"`
	Target    string `json:"target"`
	FieldType string `json:"fieldType"`
	SubDomain string `json:"subDomain"`
}

// String return ZoneRecord as string
func (r *ZoneRecord) String() string {
	//123345 toorop.fr A 300 Target
	out := fmt.Sprintf("%d ", r.ID)
	if r.SubDomain != "" {
		out += r.SubDomain + "."
	}
	out += fmt.Sprintf("%s %s %d %s", r.Zone, r.FieldType, r.TTL, r.Target)
	return out
}
