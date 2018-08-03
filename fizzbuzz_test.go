package fizzbuzz

import "testing"

type tc struct {
	title string
	input int
	expected string
}

var tests = []tc{
	{"ReturnsTheSameNumberWhenNormal", 1, "1"},
	{"ReturnFizzWhenIsMultipleOfThree", 3, "Fizz"},
	{"ReturnBuzzWhenIsMultipleOfFive", 5, "Buzz"},
	{"ReturnPopWhenIsMultipleOfSeven", 7, "Pop"},
	{"ReturnFizz_BuzzWhenIsMultipleOfThreeAndFive", 15, "Fizz Buzz"},
	{"ReturnPop_FizzWhenIsMultipleOfThreeAndSeven", 21, "Pop Fizz"},
	{"ReturnBuzz_PopWhenIsMultipleOfSevenAndFive", 35, "Buzz Pop"},
	{"ReturnFizzBuzzPopWhenIsMultipleOfThreeAndFiveAndSeven", 735, "Fizz Buzz Pop"},
	{"ReturnBazzWhenIsBetweenTwoAndTwenty", 19, "Bazz"},
}

func TestFizzBuzz(t *testing.T) {
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			out := FizzBuzz(test.input)
			if out != test.expected {
				t.Errorf("Wrong FizzBuzz output, expected: %v, given: %v", test.expected, out)
			}
		})
	}
}