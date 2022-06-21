package utilgo

import "reflect"

func ArrayIn(target interface{}, array interface{}) bool {
	for _, v := range array.([]interface{}) {
		if v == target {
			return true
		}
	}
	return false
}

func ArrayReverse(array interface{}) {
	n := reflect.ValueOf(array).Len()
	swap := reflect.Swapper(array)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
