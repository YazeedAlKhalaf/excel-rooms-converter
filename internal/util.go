package internal

func IsDayName(s string) bool {
	if s == "Sunday" || s == "Monday" || s == "Tuesday" || s == "Wednesday" || s == "Thursday" {
		return true
	}

	return false
}
