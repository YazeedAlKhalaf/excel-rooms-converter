package internal

type TimeSlot struct {
	TimeStart  TimeOfDay `json:"time_start"`
	TimeEnd    TimeOfDay `json:"time_end"`
	CourseName string
}

func (ts *TimeSlot) IsFree() bool {
	return ts.CourseName == ""
}
