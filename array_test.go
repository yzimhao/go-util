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
