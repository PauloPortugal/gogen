package slices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAll(t *testing.T) {
	Convey("Given I have a slice of even numbers from 2 to 20", t, func() {
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

	Convey("Given I have a slice of numbers from 2 to 21", t, func() {
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
	Convey("Given I have a slice of numbers from 1 to 10", t, func() {
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

	Convey("Given I have a slice of numbers from 1 to 10", t, func() {
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
	Convey("Given I have a slice of numbers from 1 to 10", t, func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Convey("When I call 'Count'", func() {
			actual := Count(input)
			Convey("Then the result should be 10", func() {
				So(actual, ShouldEqual, 10)
			})
		})
	})

	Convey("Given I have an empty slice", t, func() {
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
	Convey("Given I have a slice of numbers from 1 to 10", t, func() {
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

	Convey("Given I have a slice of numbers from 1 to 10", t, func() {
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

	Convey("Given I have an empty slice", t, func() {
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
	Convey("Given I have a slice of duplicate numbers from 1 to 3", t, func() {
		input := []int{1, 2, 3, 3, 2, 1}
		Convey("When I call 'Count'", func() {
			actual := Distinct(input)
			Convey("Then the result should be [1, 2, 3]", func() {
				So(len(actual), ShouldEqual, 3)
				So(actual, ShouldResemble, []int{1, 2, 3})
			})
		})
	})

	Convey("Given I have an empty slice", t, func() {
		var input []int
		Convey("When I call 'Distinct'", func() {
			actual := Distinct(input)
			Convey("Then the result should be an empty slice", func() {
				So(len(actual), ShouldEqual, 0)
			})
		})
	})
}

func TestDistinctBy(t *testing.T) {
	Convey("Given I have a slice of duplicate numbers from 1 to 3", t, func() {
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

	Convey("Given I have an empty slice", t, func() {
		var input []int
		Convey("When I call 'Distinct'", func() {
			actual := DistinctBy(input, func(e int) int {
				return e
			})
			Convey("Then the result should be an empty slice", func() {
				So(len(actual), ShouldEqual, 0)
			})
		})
	})
}

func TestFilter(t *testing.T) {
	Convey("Given I have a slice of numbers from 1 to 10", t, func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Convey("When I call 'Filter' with a function to filter multiples of 5", func() {
			actual := Filter(input, func(e int) bool {
				return e%5 == 0
			})
			Convey("Then the result should be a slice of elements 5 and 10", func() {
				expected := []int{5, 10}
				So(actual, ShouldResemble, expected)
			})
		})
	})

	Convey("Given I have an empty slice", t, func() {
		var input []int
		Convey("When I call 'Filter' with a function to filter multiples of 5", func() {
			actual := Filter(input, func(e int) bool {
				return e%5 == 0
			})
			Convey("Then the result should be an empty slice", func() {
				var expected []int
				So(actual, ShouldResemble, expected)
			})
		})
	})
}

func TestMap(t *testing.T) {
	Convey("Given I have a slice of numbers from 1 to 10", t, func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Convey("When I call 'Map' with a function to double the element", func() {
			actual := Map(input, func(e int) int {
				return e * 2
			})
			Convey("Then the result should be a slice of elements 5 and 10", func() {
				expected := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
				So(actual, ShouldResemble, expected)
			})
		})
	})

	Convey("Given I have an empty slice", t, func() {
		var input []int
		Convey("When I call 'Map' with a function to double the element", func() {
			actual := Map(input, func(e int) int {
				return e * 2
			})
			Convey("Then the result should be an empty slice", func() {
				So(len(actual), ShouldEqual, 0)
			})
		})
	})
}

func TestMax(t *testing.T) {
	Convey("Given I have an unordered slice of numbers between 1 to 10", t, func() {
		input := []int{2, 4, 6, 8, 10, 1, 3, 5, 7, 9}
		Convey("When I call 'Min'", func() {
			actual := Max(input, func(e int, max int) bool {
				return e > max
			})
			Convey("Then the result should be '10'", func() {
				So(*actual, ShouldEqual, 10)
			})
		})
	})

	Convey("Given I have a slice of numbers between 1 to 10", t, func() {
		var input []int
		Convey("When I call 'Max'", func() {
			actual := Max(input, func(e int, max int) bool {
				return e < max
			})
			Convey("Then the result should be nil", func() {
				So(actual, ShouldEqual, nil)
			})
		})
	})
}

func TestMin(t *testing.T) {
	Convey("Given I have an unordered slice of numbers between 1 to 10", t, func() {
		input := []int{2, 4, 6, 8, 10, 1, 3, 5, 7, 9}
		Convey("When I call 'Min'", func() {
			actual := Min(input, func(e int, min int) bool {
				return e < min
			})
			Convey("Then the result should be '1'", func() {
				So(*actual, ShouldEqual, 1)
			})
		})
	})

	Convey("Given I have an empty slice", t, func() {
		var input []int
		Convey("When I call 'Min'", func() {
			actual := Min(input, func(e int, min int) bool {
				return e < min
			})
			Convey("Then the result should be nil", func() {
				So(actual, ShouldEqual, nil)
			})
		})
	})
}

func TestNoneMatch(t *testing.T) {
	Convey("Given I have a slice of even numbers from 2 to 20", t, func() {
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

	Convey("Given I have a slice of numbers from 2 to 21", t, func() {
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

func TestReduce(t *testing.T) {
	Convey("Given I have a slice of numbers from 1 to 5", t, func() {
		input := []int{1, 2, 3, 4, 5}
		Convey("When I call 'Reduce'", func() {
			initAcc := 0
			actual := Reduce(input, initAcc, func(acc, e int) int {
				return acc + e
			})
			Convey("Then the result should be equal to 15", func() {
				So(actual, ShouldEqual, 15)
			})
		})
	})

	Convey("Given I have a slice of strings representing the first 5 letters of the alphabet", t, func() {
		input := []string{"a", "b", "c", "d", "e"}
		Convey("When I call 'Reduce'", func() {
			initAcc := ""
			actual := Reduce(input, initAcc, func(acc, e string) string {
				return acc + e
			})
			Convey("Then the result should be equal to 'abcde'", func() {
				So(actual, ShouldEqual, "abcde")
			})
		})
	})
}

func TestSkip(t *testing.T) {
	Convey("Given I have a slice of numbers between 1 to 10", t, func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Convey("When I 'Skip' the first 5 elements ", func() {
			actual := Skip(input, 5)
			Convey("Then the result should be [6, 7, 8, 9, 10]", func() {
				So(actual, ShouldResemble, []int{6, 7, 8, 9, 10})
			})
		})
	})

	Convey("Given I have an empty slice", t, func() {
		var input []int
		Convey("When I 'Skip' the first 5 elements ", func() {
			actual := Skip(input, 5)
			Convey("Then the result should be nil", func() {
				So(actual, ShouldResemble, make([]int, 0))
			})
		})
	})
}
