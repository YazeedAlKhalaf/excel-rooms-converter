package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFreeCourseNameEmpty(t *testing.T) {
	ts := &TimeSlot{
		TimeStart:  TimeOfDay{},
		TimeEnd:    TimeOfDay{},
		CourseName: "",
	}

	assert.True(t, ts.IsFree())
}

func TestIsFreeCourseNameNotEmpty(t *testing.T) {
	ts := &TimeSlot{
		TimeStart:  TimeOfDay{},
		TimeEnd:    TimeOfDay{},
		CourseName: "ENG 101",
	}

	assert.False(t, ts.IsFree())
}
