package goutil

func ArrayIn(target interface{}, array []interface{}) bool {
	for _, v := range array {
		if v == target {
			return true
		}
	}
	return false
}
