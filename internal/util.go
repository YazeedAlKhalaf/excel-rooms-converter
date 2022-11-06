package internal

func IsDayName(s string) bool {
	if s == "Sunday" || s == "Monday" || s == "Tuesday" || s == "Wednesday" || s == "Thursday" {
		return true
	}

	return false
}

func InsertToSlice[SliceType comparable](org []SliceType, objectToInsert SliceType, insertionIndex int) []SliceType {
	return append(
		org[:insertionIndex],
		append(
			[]SliceType{objectToInsert},
			org[insertionIndex:]...,
		)...,
	)
}
