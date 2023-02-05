package internal

type TimeSlot struct {
	TimeStart  TimeOfDay `json:"timeStart"`
	TimeEnd    TimeOfDay `json:"timeEnd"`
	CourseName string    `json:"courseName"`
}

func (ts *TimeSlot) IsFree() bool {
	return ts.CourseName == ""
}
