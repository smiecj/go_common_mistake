package go_common_mistake

import (
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
)

func add(a, b int) int {
	return a + b
}

func TestMockFunc(t *testing.T) {
	Convey("TestApplyFunc", t, func() {

		Convey("one func for succ", func() {
			patches := ApplyFunc(add, func(_, _ int) int {
				return 0
			})
			defer patches.Reset()
			addRet := add(1, 2)
			So(addRet, ShouldNotEqual, 3)
		})
	})
}
