package domain

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
	ID        string `json:"id"`
	Zone      string `json:"zone"`
	TTL       int    `json:"ttl"`
	Target    string `json:"target"`
	FieldType string `json:"fielType"`
	SubDomain string `json:"subDomain"`
}
