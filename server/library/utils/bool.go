package utils

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func IntToBool(i int) bool {
	if i == 1 {
		return true
	}
	return false
}
