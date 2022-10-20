package internal

type Room struct {
	Name      string     `json:"name"`
	Sunday    []TimeSlot `json:"sunday"`
	Monday    []TimeSlot `json:"monday"`
	Tuesday   []TimeSlot `json:"tuesday"`
	Wednesday []TimeSlot `json:"wednesday"`
	Thursday  []TimeSlot `json:"thursday"`
}

func (r *Room) SetName(n string) {
	r.Name = n
}

func (r *Room) AppendTimeSlotToSunday(ts TimeSlot) {
	r.Sunday = append(r.Sunday, ts)
}

func (r *Room) AppendTimeSlotToMonday(ts TimeSlot) {
	r.Monday = append(r.Monday, ts)
}

func (r *Room) AppendTimeSlotToTuesday(ts TimeSlot) {
	r.Tuesday = append(r.Tuesday, ts)
}

func (r *Room) AppendTimeSlotToWednesday(ts TimeSlot) {
	r.Wednesday = append(r.Wednesday, ts)
}

func (r *Room) AppendTimeSlotToThursday(ts TimeSlot) {
	r.Thursday = append(r.Thursday, ts)
}
