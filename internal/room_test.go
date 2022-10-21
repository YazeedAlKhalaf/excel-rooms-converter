package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetName(t *testing.T) {
	room := &Room{}
	assert.Equal(t, "", room.Name)

	room.SetName("F105")
	assert.Equal(t, "F105", room.Name)

}

func TestAppendTimeSlotToSunday(t *testing.T) {
	room := &Room{}
	assert.Equal(t, 0, len(room.Sunday))

	ts := TimeSlot{}
	room.AppendTimeSlotToSunday(ts)
	assert.Equal(t, 1, len(room.Sunday))

	assert.Equal(t, ts, room.Sunday[0])
}

func TestAppendTimeSlotToMonday(t *testing.T) {
	room := &Room{}
	assert.Equal(t, 0, len(room.Monday))

	ts := TimeSlot{}
	room.AppendTimeSlotToMonday(ts)
	assert.Equal(t, 1, len(room.Monday))

	assert.Equal(t, ts, room.Monday[0])
}

func TestAppendTimeSlotToTuesday(t *testing.T) {
	room := &Room{}
	assert.Equal(t, 0, len(room.Tuesday))

	ts := TimeSlot{}
	room.AppendTimeSlotToTuesday(ts)
	assert.Equal(t, 1, len(room.Tuesday))

	assert.Equal(t, ts, room.Tuesday[0])
}

func TestAppendTimeSlotToWednesday(t *testing.T) {
	room := &Room{}
	assert.Equal(t, 0, len(room.Wednesday))

	ts := TimeSlot{}
	room.AppendTimeSlotToWednesday(ts)
	assert.Equal(t, 1, len(room.Wednesday))

	assert.Equal(t, ts, room.Wednesday[0])
}

func TestAppendTimeSlotToThursday(t *testing.T) {
	room := &Room{}
	assert.Equal(t, 0, len(room.Thursday))

	ts := TimeSlot{}
	room.AppendTimeSlotToThursday(ts)
	assert.Equal(t, 1, len(room.Thursday))

	assert.Equal(t, ts, room.Thursday[0])
}
