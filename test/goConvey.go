package test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// Все тесты пройдут успешно
// Для запуска сервера с результатами теста требуется ввести $GOPATH/bin/goconvey
func TestExampleCleanup(t *testing.T) {
	x := 0
	Convey("A", t, func() {
		x++
		Convey("A-B", func() {
			x++
			Convey("A-B-C1", func() {
				So(x, ShouldEqual, 2)
			})
			Convey("A-B-C2", func() {
				So(x, ShouldEqual, 4)
			})
			Convey("A-B-C3", func() {
				So(x, ShouldEqual, 6)
			})
		})
	})

	Convey("Somvething wirkds properly", t, func() {
		So(1, ShouldEqual, 1)
		So(2*2, ShouldEqual, 4)
		Convey("More test", func() {
			So(1, ShouldEqual, 1)
			So(2*2, ShouldEqual, 4)
		})
	})
}
