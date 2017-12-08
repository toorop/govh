package govh

import (
	"encoding/json"
	//"fmt"
	"time"
)

// Datetimes
// DateTime represents date as returned by OVH
type DateTime struct {
	time.Time
}

func (dt *DateTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == "\"\"" {
		dt.Time = time.Time{}
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	//2014-09-16T06:50:09+02:00 RFC3339
	//t, err := time.Parse("2006-01-02T15:04:05+02:00", s)
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	dt.Time = t
	return nil
}

func (dt DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal((*time.Time)(&dt.Time).Format("2006-01-02T15:04:05+02:00"))
}

// DateTime2 represents an other date as returned by OVH
// don't ask me why but OVH use differents TZ for datetime...
type DateTime2 struct {
	time.Time
}

func (dt *DateTime2) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == "\"\"" {
		dt.Time = time.Time{}
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	// 2014-09-15T13:41:35+02:00
	t, err := time.Parse("2006-01-02T15:04:05+01:00", s)
	if err != nil {
		return err
	}
	dt.Time = t
	return nil
}

func (dt DateTime2) MarshalJSON() ([]byte, error) {
	return json.Marshal((*time.Time)(&dt.Time).Format("2006-01-02T15:04:05+01:00"))
}
