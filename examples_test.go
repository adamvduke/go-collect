package collect_test

import (
	"fmt"
	"maps"
	"slices"

	"github.com/adamvduke/go-collect"
)

func ExampleApply_stringToInt() {
	input := []string{"a", "bb", "ccc", "dddd"}
	out := collect.Apply(input, func(s string) int {
		return len(s)
	})
	fmt.Println(out)
	// Output: [1 2 3 4]
}

func ExampleApply_structToString() {
	type Foo struct {
		bar string
	}
	input := []Foo{
		{bar: "a"},
		{bar: "bb"},
		{bar: "ccc"},
		{bar: "dddd"},
	}
	out := collect.Apply(input, func(foo Foo) string {
		return foo.bar + "a"
	})
	fmt.Println(out)
	// Output: [aa bba ccca dddda]
}

func ExampleApply_structToStruct() {
	type In struct {
		bar string
	}
	type Out struct {
		baz string
	}
	input := []In{
		{bar: "a"},
		{bar: "bb"},
		{bar: "ccc"},
		{bar: "dddd"},
	}
	out := collect.Apply(input, func(foo In) Out {
		return Out{baz: foo.bar + "a"}
	})
	fmt.Printf("%+v", out)
	// Output: [{baz:aa} {baz:bba} {baz:ccca} {baz:dddda}]
}

func ExampleSelect() {
	input := []int{1, 2, 3, 4}
	got := collect.Select(input, func(s int) bool {
		return s > 2
	})
	fmt.Println(got)
	// Output: [3 4]
}

func ExampleReject() {
	input := []int{1, 2, 3, 4}
	got := collect.Reject(input, func(s int) bool {
		return s > 2
	})
	fmt.Println(got)
	// Output: [1 2]
}

func ExampleKeys() {
	input := map[int]any{1: "1", 2: 1.2, 12: struct{}{}}
	got := collect.Keys(input)
	slices.Sort(got)
	fmt.Println(got)
	// Output: [1 2 12]
}

func ExampleValues() {
	input := map[any]any{1: "1", 2: 1.2, 10: struct{}{}}
	got := slices.Collect(maps.Values(input))
	fmt.Println("contains \"1\":", slices.Contains(got, "1"))
	fmt.Println("contains 1.2:", slices.Contains(got, 1.2))
	fmt.Println("contains struct{}{}:", slices.Contains(got, any(struct{}{})))
	// Output:
	// contains "1": true
	// contains 1.2: true
	// contains struct{}{}: true
}
