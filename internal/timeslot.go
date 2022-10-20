package internal

type TimeSlot struct {
	TimeStart  TimeOfDay `json:"time_start"`
	TimeEnd    TimeOfDay `json:"time_end"`
	CourseName string    `json:"course_name"`
}

func (ts *TimeSlot) IsFree() bool {
	return ts.CourseName == ""
}
