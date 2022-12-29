package slices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAll(t *testing.T) {
	Convey("Given I have an array of even numbers from 2 to 20", t, func() {
		input := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
		Convey("When I call 'All' to test even numbers", func() {
			actual := All(input, func(e int) bool {
				return e%2 == 0
			})
			Convey("Then the result should be true", func() {
				So(actual, ShouldEqual, true)
			})
		})
	})

	Convey("Given I have an array of numbers from 2 to 21", t, func() {
		input := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 21}
		Convey("When I call 'All' to test even numbers", func() {
			actual := All(input, func(e int) bool {
				return e%2 == 0
			})
			Convey("Then the result should be false", func() {
				So(actual, ShouldEqual, false)
			})
		})
	})
}

func TestAny(t *testing.T) {
	Convey("Given I have an array of numbers from 1 to 10", t, func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Convey("When I call 'Any' for values greater than 7", func() {
			actual := Any(input, func(e int) bool {
				return e > 7
			})
			Convey("Then the result should be true", func() {
				So(actual, ShouldEqual, true)
			})
		})
	})

	Convey("Given I have an array of numbers from 1 to 10", t, func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Convey("When I call 'Any' for values greater than 10", func() {
			actual := Any(input, func(e int) bool {
				return e > 10
			})
			Convey("Then the result should be false", func() {
				So(actual, ShouldEqual, false)
			})
		})
	})
}

func TestCount(t *testing.T) {
	Convey("Given I have an array of numbers from 1 to 10", t, func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Convey("When I call 'Count'", func() {
			actual := Count(input)
			Convey("Then the result should be 10", func() {
				So(actual, ShouldEqual, 10)
			})
		})
	})

	Convey("Given I have an empty array", t, func() {
		var input []int
		Convey("When I call 'Count'", func() {
			actual := Count(input)
			Convey("Then the result should be 0", func() {
				So(actual, ShouldEqual, 0)
			})
		})
	})
}

func TestContains(t *testing.T) {
	Convey("Given I have an array of numbers from 1 to 10", t, func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Convey("When I call 'Contains' to find '2'", func() {
			actual := Contains(input, func(e int) bool {
				return e == 2
			})
			Convey("Then the result should be 'true'", func() {
				So(actual, ShouldEqual, true)
			})
		})
	})

	Convey("Given I have an array of numbers from 1 to 10", t, func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Convey("When I call 'Contains' to find '12'", func() {
			actual := Contains(input, func(e int) bool {
				return e == 12
			})
			Convey("Then the result should be 'false'", func() {
				So(actual, ShouldEqual, false)
			})
		})
	})

	Convey("Given I have an empty array", t, func() {
		var input []int
		Convey("When I call 'Contains' to find '2'", func() {
			actual := Contains(input, func(e int) bool {
				return e == 2
			})
			Convey("Then the result should be 'false'", func() {
				So(actual, ShouldEqual, false)
			})
		})
	})
}

func TestDistinct(t *testing.T) {
	Convey("Given I have an array of duplicate numbers from 1 to 3", t, func() {
		input := []int{1, 2, 3, 3, 2, 1}
		Convey("When I call 'Count'", func() {
			actual := Distinct(input)
			Convey("Then the result should be [1, 2, 3]", func() {
				So(len(actual), ShouldEqual, 3)
				So(actual, ShouldResemble, []int{1, 2, 3})
			})
		})
	})

	Convey("Given I have an empty array", t, func() {
		var input []int
		Convey("When I call 'Distinct'", func() {
			actual := Distinct(input)
			Convey("Then the result should be an empty array", func() {
				So(len(actual), ShouldEqual, 0)
			})
		})
	})
}

func TestDistinctBy(t *testing.T) {
	Convey("Given I have an array of duplicate numbers from 1 to 3", t, func() {
		input := []int{1, 2, 3, 3, 2, 1}
		Convey("When I call 'Count'", func() {
			actual := DistinctBy(input, func(e int) int {
				return e
			})
			Convey("Then the result should be [1, 2, 3]", func() {
				So(len(actual), ShouldEqual, 3)
				So(actual, ShouldResemble, []int{1, 2, 3})
			})
		})
	})

	Convey("Given I have an empty array", t, func() {
		var input []int
		Convey("When I call 'Distinct'", func() {
			actual := DistinctBy(input, func(e int) int {
				return e
			})
			Convey("Then the result should be an empty array", func() {
				So(len(actual), ShouldEqual, 0)
			})
		})
	})
}

func TestFilter(t *testing.T) {
	Convey("Given I have an array of numbers from 1 to 10", t, func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Convey("When I call 'Filter' with a function to filter multiples of 5", func() {
			actual := Filter(input, func(e int) bool {
				return e%5 == 0
			})
			Convey("Then the result should be an array of elements 5 and 10", func() {
				expected := []int{5, 10}
				So(actual, ShouldResemble, expected)
			})
		})
	})

	Convey("Given I have an empty array", t, func() {
		var input []int
		Convey("When I call 'Filter' with a function to filter multiples of 5", func() {
			actual := Filter(input, func(e int) bool {
				return e%5 == 0
			})
			Convey("Then the result should be an empty array", func() {
				var expected []int
				So(actual, ShouldResemble, expected)
			})
		})
	})
}

func TestMap(t *testing.T) {
	Convey("Given I have an array of numbers from 1 to 10", t, func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Convey("When I call 'Map' with a function to double the element", func() {
			actual := Map(input, func(e int) int {
				return e * 2
			})
			Convey("Then the result should be an array of elements 5 and 10", func() {
				expected := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
				So(actual, ShouldResemble, expected)
			})
		})
	})

	Convey("Given I have an empty array", t, func() {
		var input []int
		Convey("When I call 'Map' with a function to double the element", func() {
			actual := Map(input, func(e int) int {
				return e * 2
			})
			Convey("Then the result should be an empty array", func() {
				So(len(actual), ShouldEqual, 0)
			})
		})
	})
}

func TestNoneMatch(t *testing.T) {
	Convey("Given I have an array of even numbers from 2 to 20", t, func() {
		input := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
		Convey("When I call 'None Match' to test odd numbers", func() {
			actual := NoneMatch(input, func(e int) bool {
				return e%2 != 0
			})
			Convey("Then the result should be true", func() {
				So(actual, ShouldEqual, true)
			})
		})
	})

	Convey("Given I have an array of numbers from 2 to 21", t, func() {
		input := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 21}
		Convey("When I call 'None Match' to test odd numbers", func() {
			actual := NoneMatch(input, func(e int) bool {
				return e%2 != 0
			})
			Convey("Then the result should be false", func() {
				So(actual, ShouldEqual, false)
			})
		})
	})
}
