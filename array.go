package goutil

func ArrayIn(target interface{}, array interface{}) bool {
	for _, v := range array.([]interface{}) {
		if v == target {
			return true
		}
	}
	return false
}
