package collect_test

import (
	"fmt"

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
