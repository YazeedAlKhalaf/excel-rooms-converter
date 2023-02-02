package internal

import (
	"fmt"
	"strconv"
	"strings"
)

type TimeOfDay struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

func NewTimeOfDay(hour int, minute int) (*TimeOfDay, error) {
	if hour < 0 || hour > 23 || minute < 0 || minute > 59 {
		return &TimeOfDay{}, fmt.Errorf("invalid hour or minute, expected: 0 <= hour <= 23, 0 <= minute <= 59, but got: %d:%d", hour, minute)
	}

	return &TimeOfDay{
		Hour:   hour,
		Minute: minute,
	}, nil
}

func ParseTimeOfDay(time string) (*TimeOfDay, error) {
	timePartsWithAMandPM := strings.Split(time, " ")
	if len(timePartsWithAMandPM) != 2 {
		return &TimeOfDay{}, fmt.Errorf("invalid time format, expected: 12:30 AM, got: %s", time)
	}

	// (isAM && isPM) || (!isAM && !isPM)
	// that can't be both true or both false, so it's invalid
	isAM := timePartsWithAMandPM[1] == "AM" || timePartsWithAMandPM[1] == "am"
	isPM := timePartsWithAMandPM[1] == "PM" || timePartsWithAMandPM[1] == "pm"
	if isAM == isPM {
		return &TimeOfDay{}, fmt.Errorf("invalid am or pm format, expected: 12:30 AM, but got: %s", time)
	}

	timeParts := strings.Split(timePartsWithAMandPM[0], ":")
	if len(timeParts) != 2 {
		return &TimeOfDay{}, fmt.Errorf("invalid time format, expected: 12:30 AM, but got: %s", time)
	}

	hour, err := strconv.ParseInt(timeParts[0], 10, 64)
	if err != nil {
		return &TimeOfDay{}, fmt.Errorf("invalid hour format, expected: 12:30 AM, but got: %s", time)
	}
	if isPM {
		hour += 12
		if hour == 24 {
			hour = 0
		}
	}

	minute, err := strconv.ParseInt(timeParts[1], 10, 64)
	if err != nil {
		return &TimeOfDay{}, fmt.Errorf("invalid minute format, expected: 12:30 AM, but got: %s", time)
	}

	return NewTimeOfDay(int(hour), int(minute))
}

func (t *TimeOfDay) IsBefore(t2 *TimeOfDay) bool {
	if t.Hour < t2.Hour {
		return true
	}

	if t.Hour > t2.Hour {
		return false
	}

	// if we reach this, it means the hours are equal
	return t.Minute < t2.Minute
}

func (t *TimeOfDay) IsAfter(t2 *TimeOfDay) bool {
	if t.Hour > t2.Hour {
		return true
	}

	if t.Hour < t2.Hour {
		return false
	}

	// if we reach this, it means the hours are equal
	return t.Minute > t2.Minute
}

func (t *TimeOfDay) IsEqual(ts *TimeOfDay) bool {
	return t.Hour == ts.Hour && t.Minute == ts.Minute
}

func (t *TimeOfDay) AddMinutes(minutes int) (*TimeOfDay, error) {
	if minutes < 0 {
		return &TimeOfDay{}, fmt.Errorf("invalid minutes, expected: minutes >= 0, but got: %d", minutes)
	}

	minutes += t.Minute
	hours := t.Hour + minutes/60
	minutes = minutes % 60

	return NewTimeOfDay(hours, minutes)
}
