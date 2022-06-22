package utilgo

import "reflect"

func ArrayIn(needle interface{}, array interface{}) bool {
	exists := false

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) == true {
				exists = true
				return exists
			}
		}
	}

	return exists
}

func ArrayReverse(array interface{}) {
	n := reflect.ValueOf(array).Len()
	swap := reflect.Swapper(array)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
