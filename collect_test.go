package collect_test

import (
	"testing"

	"github.com/adamvduke/go-collect"
	"github.com/stretchr/testify/require"
)

func TestApply_StringToInt(t *testing.T) {
	input := []string{"a", "bb", "ccc", "dddd"}
	got := collect.Apply(input, func(s string) int {
		return len(s)
	})
	want := []int{1, 2, 3, 4}
	require.ElementsMatch(t, got, want, "Apply(%v, func(s string) int {...}) got: %v, want: %v", input, got, want)
}

func TestApply_StringToString(t *testing.T) {
	input := []string{"a", "bb", "ccc", "dddd"}
	got := collect.Apply(input, func(s string) string {
		return s + "-suffix"
	})
	want := []string{"a-suffix", "bb-suffix", "dddd-suffix", "ccc-suffix"}
	require.ElementsMatch(t, got, want, "Apply(%v, func(s string) string {...}) got: %v, want: %v", input, got, want)
}

func TestApply_StructToString(t *testing.T) {
	type Foo struct {
		bar string
	}
	input := []Foo{
		{bar: "a"},
		{bar: "bb"},
		{bar: "ccc"},
		{bar: "dddd"},
	}
	got := collect.Apply(input, func(foo Foo) string {
		return foo.bar + "-suffix"
	})
	want := []string{"a-suffix", "bb-suffix", "dddd-suffix", "ccc-suffix"}
	require.ElementsMatch(t, got, want, "Apply(%+v, func(s Foo) string {...}) got: %+v, want: %v", input, got, want)
}

func TestApply_StructToStruct(t *testing.T) {
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
	got := collect.Apply(input, func(foo In) Out {
		return Out{baz: foo.bar + "-suffix"}
	})
	want := []Out{{baz: "a-suffix"}, {baz: "bb-suffix"}, {baz: "dddd-suffix"}, {baz: "ccc-suffix"}}
	require.ElementsMatch(t, got, want, "Apply(%+v, func(s In) Out {...}) got: %+v, want: %v", input, got, want)
}

func TestKeys(t *testing.T) {
	input := map[int]any{1: "1", 2: 1.2, 12: struct{}{}}
	got := collect.Keys(input)
	want := []int{1, 12, 2}
	require.ElementsMatch(t, got, want)
}

func TestValues(t *testing.T) {
	input := map[int]any{1: "1", 2: 1.2, 10: struct{}{}}
	got := collect.Values(input)
	want := []any{"1", struct{}{}, 1.2}
	require.ElementsMatch(t, got, want)
}

func BenchmarkApply_StringToInt(b *testing.B) {
	b.ReportAllocs()
	input := []string{"a", "bb", "ccc", "dddd"}
	for i := 0; i < b.N; i++ {
		collect.Apply(input, func(s string) int {
			return len(s)
		})
	}
}

func BenchmarkApply_StructToString(b *testing.B) {
	b.ReportAllocs()
	type Foo struct {
		bar string
	}
	input := []Foo{
		{bar: "a"},
		{bar: "bb"},
		{bar: "ccc"},
		{bar: "dddd"},
	}
	for i := 0; i < b.N; i++ {
		collect.Apply(input, func(foo Foo) string {
			return foo.bar + "-suffix"
		})
	}
}

func BenchmarkApply_StructToStruct(b *testing.B) {
	b.ReportAllocs()
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
	for i := 0; i < b.N; i++ {
		collect.Apply(input, func(foo In) Out {
			return Out{baz: foo.bar + "-suffix"}
		})
	}
}
