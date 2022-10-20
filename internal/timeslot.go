package internal

type TimeSlot struct {
	TimeStart  TimeOfDay
	TimeEnd    TimeOfDay
	CourseName string
}

func (ts *TimeSlot) IsFree() bool {
	return ts.CourseName == ""
}
