package core

import (
	"github.com/araddon/dateparse"
	"time"
)

var loc, _ = time.LoadLocation("Asia/Shanghai")
var normalFormatLen = len("2006-01-02 15:04:05")

func (m Timestamp) Format() string {
	return m.ToTime().Format(time.RFC3339)
}

func ParseTimestamp(value string) (*Timestamp, error) {
	ts := &Timestamp{}
	if err := ts.Parse(value); err != nil {
		return nil, err
	}
	return ts, nil
}

func (m *Timestamp) Parse(value string) error {
	if m != nil {
		// ts has timezone info, like "2006-01-02 15:04:05+0800"
		// since '+' will be replaced by space in url, we restore it to '+' if possible
		if len(value) > normalFormatLen && value[normalFormatLen] == ' ' {
			value = value[:normalFormatLen] + "+" + value[normalFormatLen+1:]
		}

		t, err := dateparse.ParseIn(value, loc)
		if err != nil {
			return err
		}

		m.FromTime(t)
	}
	return nil
}