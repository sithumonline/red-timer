package time

import (
	"fmt"
	"time"
)

type Time struct {
	Duration string `json:"duration,omitempty"`
	Color    string `json:"color,omitempty"`
}

var mainTime *time.Time

func (obj *Time) GetTime() Time {
	if mainTime == nil {
		pointTime := time.Now()
		mainTime = &pointTime
	}
	currentTime := time.Now()
	difference := currentTime.Sub(*mainTime)
	d := fmt.Sprintf("%v:%v:%v", 0-int(difference.Hours()), 0-int(difference.Minutes()), 60-int(difference.Seconds()))
	c := "black"
	if int(difference.Seconds()) >= 50 {
		c = "red"
	}
	if int(difference.Seconds()) >= 60 {
		d = "0:0:0"
	}
	return Time{
		Duration: d,
		Color:    c,
	}
}

func (obj *Time) ResetTime() {
	mainTime = nil
}
