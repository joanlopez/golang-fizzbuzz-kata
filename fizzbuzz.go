package fizzbuzz

import (
	"strconv"
	"strings"
)

type Verifier func(x int, results []string) bool
type Applier func(results []string) []string

type Rule struct {
	v Verifier
	a Applier
}

type RuleMachine struct {
	rules []Rule
}

func (m *RuleMachine) execute(x int) []string {
	var results []string
	for _, rule := range m.rules {
		if rule.v(x, results) {
			results = rule.a(results)
		}
	}
	return results
}

func isMultipleOf(x, y int) bool {
	return x%y == 0
}

func isBetween(n, x, y int) bool {
	return x <= n  && n <= y
}

var rules = []Rule{
	{
		func(x int, results []string) bool { return isMultipleOf(x, 3) },
		func(result []string) []string { return append(result, "Fizz") },
	},
	{
		func(x int, results []string) bool { return isMultipleOf(x, 5) },
		func(result []string) []string { return append(result, "Buzz") },
	},
	{
		func(x int, results []string) bool { return isMultipleOf(x, 7) },
		func(result []string) []string { return append(result, "Pop") },
	},
	{
		func(x int, results []string) bool { return isMultipleOf(x, 7) && isMultipleOf(x, 3) && len(results) == 2 },
		func(result []string) []string { return []string{result[1], result[0]} },
	},
	{
		func(x int, results []string) bool { return isBetween(x, 2, 20) && len(results) == 0},
		func(result []string) []string { return []string{"Bazz"} },
	},
}

func FizzBuzz(x int) string {
	machine := &RuleMachine{rules}
	result := machine.execute(x)

	if len(result) == 0 {
		return strconv.Itoa(x)
	}

	return strings.Join(result, " ")
}
