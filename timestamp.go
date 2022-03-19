package main

import (
	"strconv"
	"time"
)

type timestamp struct {
	time.Time
}

func (t timestamp) String() string {
	const layout = "2006-01-02 15:04"
	return t.Format(layout)
}

func (t timestamp) MarshalJSON() (data []byte, _ error) {
	return strconv.AppendInt(data, t.Unix(), 10), nil
}

func (t *timestamp) UnmarshalJSON(data []byte) error {
	*t = toTimestamp(string(data))
	return nil
}

func toTimestamp(v interface{}) (ts timestamp) {
	var t int

	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	}

	ts.Time = time.Unix(int64(t), 0)
	return ts
}
