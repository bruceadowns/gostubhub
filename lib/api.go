package lib

import "time"

// StubhubTime ...
type StubhubTime struct {
	time.Time
}

// UnmarshalJSON ...
func (stubhubTime *StubhubTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)

	var t time.Time
	t, err = time.Parse(`"2006-01-02T15:04:05-0400"`, s)
	if err != nil {
		t, err = time.Parse(`"2006-01-02T15:04:05+0400"`, s)
		if err != nil {
			return
		}
	}

	stubhubTime.Time = t
	return
}
