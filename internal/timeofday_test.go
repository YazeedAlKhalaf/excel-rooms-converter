package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTimeOfDay_HourLessThan0(t *testing.T) {
	_, err := NewTimeOfDay(-1, 0)
	assert.Error(t, err)
}

func TestNewTimeOfDay_HourMoreThan23(t *testing.T) {
	_, err := NewTimeOfDay(24, 0)
	assert.Error(t, err)
}

func TestNewTimeOfDay_MinuteLessThan0(t *testing.T) {
	_, err := NewTimeOfDay(0, -1)
	assert.Error(t, err)
}

func TestNewTimeOfDay_MinuteMoreThan59(t *testing.T) {
	_, err := NewTimeOfDay(0, 60)
	assert.Error(t, err)
}

func TestNewTimeOfDay_HourAndMinuteValid(t *testing.T) {
	_, err := NewTimeOfDay(0, 0)
	assert.NoError(t, err)
}

func TestParseTimeOfDay_NoSpaceToSplit(t *testing.T) {
	_, err := ParseTimeOfDay("12:30AM")
	assert.Error(t, err)
}

func TestParseTimeOfDay_TwoSpacesToSplit(t *testing.T) {
	_, err := ParseTimeOfDay("12:30 AM PM")
	assert.Error(t, err)
}

func TestParseTimeOfDay_NoAMorPM(t *testing.T) {
	_, err := ParseTimeOfDay("12:30 GG")
	assert.Error(t, err)
}

func TestParseTimeOfDay_AM(t *testing.T) {
	tod, err := ParseTimeOfDay("12:30 AM")
	assert.NoError(t, err)

	assert.Equal(t, 12, tod.Hour)
	assert.Equal(t, 30, tod.Minute)
}

func TestParseTimeOfDay_AMAfter12(t *testing.T) {
	tod, err := ParseTimeOfDay("1:30 AM")
	assert.NoError(t, err)

	assert.Equal(t, 1, tod.Hour)
	assert.Equal(t, 30, tod.Minute)
}

func TestParseTimeOfDay_AMAfter12WithLeadingZero(t *testing.T) {
	tod, err := ParseTimeOfDay("01:30 AM")
	assert.NoError(t, err)

	assert.Equal(t, 1, tod.Hour)
	assert.Equal(t, 30, tod.Minute)
}

func TestParseTimeOfDay_PMBefore12(t *testing.T) {
	tod, err := ParseTimeOfDay("11:30 PM")
	assert.NoError(t, err)

	assert.Equal(t, 23, tod.Hour)
	assert.Equal(t, 30, tod.Minute)
}

func TestParseTimeOfDay_PMEqual12(t *testing.T) {
	tod, err := ParseTimeOfDay("12:30 PM")
	assert.NoError(t, err)

	assert.Equal(t, 0, tod.Hour)
	assert.Equal(t, 30, tod.Minute)
}

func TestParseTimeOfDay_PMAfter12(t *testing.T) {
	tod, err := ParseTimeOfDay("1:30 PM")
	assert.NoError(t, err)

	assert.Equal(t, 13, tod.Hour)
	assert.Equal(t, 30, tod.Minute)
}

func TestParseTimeOfDay_PMAfter12WithLeadingZero(t *testing.T) {
	tod, err := ParseTimeOfDay("01:30 PM")
	assert.NoError(t, err)

	assert.Equal(t, 13, tod.Hour)
	assert.Equal(t, 30, tod.Minute)
}

func TestParseTimeOfDay_TimePartLessThan2Parts(t *testing.T) {
	_, err := ParseTimeOfDay("12 PM")
	assert.Error(t, err)
}

func TestParseTimeOfDay_TimePartMoreThan2Parts(t *testing.T) {
	_, err := ParseTimeOfDay("12:30:00 PM")
	assert.Error(t, err)
}

func TestParseTimeOfDay_TimePartHourNotInteger(t *testing.T) {
	_, err := ParseTimeOfDay("12.5:30 PM")
	assert.Error(t, err)
}

func TestParseTimeOfDay_TimePartMinuteNotInteger(t *testing.T) {
	_, err := ParseTimeOfDay("12:30.5 PM")
	assert.Error(t, err)
}
