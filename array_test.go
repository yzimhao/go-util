package utilgo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestArrayIn(t *testing.T) {

	Convey("int in array", t, func() {
		array := []interface{}{1, 2, 3, 4, 5}
		So(ArrayIn(1, array), ShouldBeTrue)
	})

	Convey("string in array", t, func() {
		array := []interface{}{"a", "b", "c", "d", "e", "1"}
		So(ArrayIn("a", array), ShouldBeTrue)
		So(ArrayIn("1", array), ShouldBeTrue)
		So(ArrayIn(1, array), ShouldBeFalse)
	})
}

func TestArrayReverse(t *testing.T) {
	Convey("arrayReverse", t, func() {
		array := []int64{1, 4, 8, 9}
		ArrayReverse(array)
		So(array, ShouldResemble, []int64{9, 8, 4, 1})

		arr := []string{"a", "b", "c"}
		ArrayReverse(arr)
		So(arr, ShouldResemble, []string{"c", "b", "a"})

		arr1 := []string{"a"}
		ArrayReverse(arr1)
		So(arr1, ShouldResemble, []string{"a"})

	})
}
