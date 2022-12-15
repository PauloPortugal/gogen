package slices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFilter(t *testing.T) {
	Convey("Given I have an array of numbers from 1 to 10", t, func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Convey("When I call 'Filter' with a function to filter multiples of 5", func() {
			expected := Filter(input, func(e int) bool {
				return e%5 == 0
			})
			Convey("Then the result should be an array of elements 5 and 10", func() {
				actual := []int{5, 10}
				So(actual, ShouldResemble, expected)
			})
		})
	})

	Convey("Given I have an empty array", t, func() {
		var input []int
		Convey("When I call 'Filter' with a function to filter multiples of 5", func() {
			expected := Filter(input, func(e int) bool {
				return e%5 == 0
			})
			Convey("Then the result should be an empty array", func() {
				var actual []int
				So(actual, ShouldResemble, expected)
			})
		})
	})
}

func TestMap(t *testing.T) {
	Convey("Given I have an array of numbers from 1 to 10", t, func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Convey("When I call 'Map' with a function to double the element", func() {
			expected := Map(input, func(e int) int {
				return e * 2
			})
			Convey("Then the result should be an array of elements 5 and 10", func() {
				actual := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
				So(actual, ShouldResemble, expected)
			})
		})
	})

	Convey("Given I have an empty array", t, func() {
		var input []int
		Convey("When I call 'Map' with a function to double the element", func() {
			expected := Map(input, func(e int) int {
				return e * 2
			})
			Convey("Then the result should be an empty array", func() {
				So(0, ShouldEqual, len(expected))
			})
		})
	})
}

func TestCount(t *testing.T) {
	Convey("Given I have an array of numbers from 1 to 10", t, func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Convey("When I call 'Count'", func() {
			expected := Count(input)
			Convey("Then the result should be 10", func() {
				So(10, ShouldEqual, expected)
			})
		})
	})

	Convey("Given I have an empty array", t, func() {
		var input []int
		Convey("When I call 'Count'", func() {
			expected := Count(input)
			Convey("Then the result should be 0", func() {
				So(0, ShouldEqual, expected)
			})
		})
	})
}

func TestContains(t *testing.T) {
	Convey("Given I have an array of numbers from 1 to 10", t, func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Convey("When I call 'Contains' to find '2'", func() {
			expected := Contains(input, func(e int) bool {
				return e == 2
			})
			Convey("Then the result should be 'true'", func() {
				So(true, ShouldEqual, expected)
			})
		})
	})

	Convey("Given I have an array of numbers from 1 to 10", t, func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Convey("When I call 'Contains' to find '12'", func() {
			expected := Contains(input, func(e int) bool {
				return e == 12
			})
			Convey("Then the result should be 'false'", func() {
				So(false, ShouldEqual, expected)
			})
		})
	})

	Convey("Given I have an empty array", t, func() {
		var input []int
		Convey("When I call 'Contains' to find '2'", func() {
			expected := Contains(input, func(e int) bool {
				return e == 2
			})
			Convey("Then the result should be 'false'", func() {
				So(false, ShouldEqual, expected)
			})
		})
	})
}
