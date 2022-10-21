package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDayName_TrueIfSunday(t *testing.T) {
	r := IsDayName("Sunday")
	assert.True(t, r)
}

func TestIsDayName_TrueIfMonday(t *testing.T) {
	r := IsDayName("Monday")
	assert.True(t, r)
}

func TestIsDayName_TrueIfTuesday(t *testing.T) {
	r := IsDayName("Tuesday")
	assert.True(t, r)
}

func TestIsDayName_TrueIfWednesday(t *testing.T) {
	r := IsDayName("Wednesday")
	assert.True(t, r)
}

func TestIsDayName_TrueIfThursday(t *testing.T) {
	r := IsDayName("Thursday")
	assert.True(t, r)
}

func TestIsDayName_FalseIfFriday(t *testing.T) {
	r := IsDayName("Friday")
	assert.False(t, r)
}

func TestIsDayName_FalseIfSaturday(t *testing.T) {
	r := IsDayName("Saturday")
	assert.False(t, r)
}
