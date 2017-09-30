// Package model implements time and date funcs for a Message
package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Message holds a returned Slack message
type Message struct {
	Text string `json:"text"`
	Ts   string `json:"ts"`
}

// GetTimeStamp parses the time stamp into a float64
func (m Message) GetTimeStamp() float64 {
	f, _ := strconv.ParseFloat(m.Ts, 64)

	return f
}

// GetTime returns the Time of a message
func (m Message) GetTime() time.Time {
	intPart, _ := strconv.Atoi(strings.Split(m.Ts, ".")[0])

	return time.Unix(int64(intPart), 0)
}

// GetDateString returns the date as a string
func (m Message) GetDateString() string {
	weekDay := m.GetTime().Weekday().String()
	year, month, day := m.GetTime().Date()
	return fmt.Sprintf("%v %4d-%2d-%2d", weekDay, year, month, day)
}

// GetTimeString returns the time as a string
func (m Message) GetTimeString() string {
	time := m.GetTime()

	return fmt.Sprintf("%02d:%02d", time.Hour(), time.Minute())
}
