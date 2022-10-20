package internal

import (
	"fmt"
	"strconv"
	"strings"
)

type TimeOfDay struct {
	hour   int
	minute int
}

func NewTimeOfDay(hour int, minute int) (*TimeOfDay, error) {
	if hour < 0 || hour > 23 || minute < 0 || minute > 59 {
		return &TimeOfDay{}, nil
	}

	return &TimeOfDay{
		hour:   hour,
		minute: minute,
	}, nil
}

func ParseTimeOfDay(time string) (*TimeOfDay, error) {
	// "12:30 AM"
	timePartsWithAMandPM := strings.Split(time, " ")
	if len(timePartsWithAMandPM) != 2 {
		return &TimeOfDay{}, fmt.Errorf("invalid time format, expected: 12:30 AM, got: %s", time)
	}
	isAM := timePartsWithAMandPM[1] == "AM"
	isPM := timePartsWithAMandPM[1] == "PM"

	// (isAM && isPM) || (!isAM && !isPM)
	// that can't be both true or both false, so it's invalid
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

	minute, err := strconv.ParseInt(timeParts[1], 10, 64)
	if err != nil {
		return &TimeOfDay{}, fmt.Errorf("invalid minute format, expected: 12:30 AM, but got: %s", time)
	}

	return NewTimeOfDay(int(hour), int(minute))
}
