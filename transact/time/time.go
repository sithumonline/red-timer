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
		pointTime := time.Now().Add(time.Second * 60)
		mainTime = &pointTime
	}
	currentTime := time.Now()
	difference := mainTime.Sub(currentTime)

	d := fmt.Sprintf("%v:%v:%v", int(difference.Hours()), int(difference.Minutes()), int(difference.Seconds()))
	c := "white"
	if int(difference.Seconds()) <= 10 {
		c = "red"
	}
	if int(difference.Seconds()) <= 0 {
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
