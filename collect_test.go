package collect_test

import (
	"testing"

	"github.com/adamvduke/go-collect"
	"github.com/stretchr/testify/require"
)

type TypeA struct {
	bar string
}

type TypeB struct {
	baz string
}

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
	input := []TypeA{{bar: "a"}, {bar: "bb"}, {bar: "ccc"}, {bar: "dddd"}}
	got := collect.Apply(input, func(foo TypeA) string {
		return foo.bar + "-suffix"
	})
	want := []string{"a-suffix", "bb-suffix", "dddd-suffix", "ccc-suffix"}
	require.ElementsMatch(t, got, want, "Apply(%+v, func(s TypeA) string {...}) got: %+v, want: %v", input, got, want)
}

func TestApply_StructToStruct(t *testing.T) {
	input := []TypeA{{bar: "a"}, {bar: "bb"}, {bar: "ccc"}, {bar: "dddd"}}
	got := collect.Apply(input, func(foo TypeA) TypeB {
		return TypeB{baz: foo.bar + "-suffix"}
	})
	want := []TypeB{{baz: "a-suffix"}, {baz: "bb-suffix"}, {baz: "dddd-suffix"}, {baz: "ccc-suffix"}}
	require.ElementsMatch(t, got, want, "Apply(%+v, func(s TypeA) TypeB {...}) got: %+v, want: %v", input, got, want)
}

func TestSelect_Int(t *testing.T) {
	input := []int{1, 2, 3, 4}
	got := collect.Select(input, func(s int) bool {
		return s > 2
	})
	want := []int{3, 4}
	require.ElementsMatch(t, got, want, "Select(%v, func(s int) bool {...}) got: %v, want: %v", input, got, want)
}

func TestSelect_String(t *testing.T) {
	input := []string{"a", "bb", "ccc", "dddd"}
	got := collect.Select(input, func(s string) bool {
		return len(s) > 2
	})
	want := []string{"ccc", "dddd"}
	require.ElementsMatch(t, got, want, "Select(%v, func(s string) bool {...}) got: %v, want: %v", input, got, want)
}

func TestSelect_Struct(t *testing.T) {
	input := []TypeA{{bar: "a"}, {bar: "bb"}, {bar: "ccc"}, {bar: "dddd"}}
	got := collect.Select(input, func(foo TypeA) bool {
		return len(foo.bar) > 2
	})
	want := []TypeA{{bar: "ccc"}, {bar: "dddd"}}
	require.ElementsMatch(t, got, want, "Select(%+v, func(s TypeA) bool {...}) got: %+v, want: %v", input, got, want)
}

func TestReject_Int(t *testing.T) {
	input := []int{1, 2, 3, 4}
	got := collect.Reject(input, func(s int) bool {
		return s > 2
	})
	want := []int{1, 2}
	require.ElementsMatch(t, got, want, "Reject(%v, func(s int) bool {...}) got: %v, want: %v", input, got, want)
}

func TestReject_String(t *testing.T) {
	input := []string{"a", "bb", "ccc", "dddd"}
	got := collect.Reject(input, func(s string) bool {
		return len(s) > 2
	})
	want := []string{"a", "bb"}
	require.ElementsMatch(t, got, want, "Reject(%v, func(s string) bool {...}) got: %v, want: %v", input, got, want)
}

func TestReject_Struct(t *testing.T) {
	input := []TypeA{{bar: "a"}, {bar: "bb"}, {bar: "ccc"}, {bar: "dddd"}}
	got := collect.Reject(input, func(foo TypeA) bool {
		return len(foo.bar) > 2
	})
	want := []TypeA{{bar: "a"}, {bar: "bb"}}
	require.ElementsMatch(t, got, want, "Reject(%+v, func(s TypeA) bool {...}) got: %+v, want: %v", input, got, want)
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
	input := []TypeA{{bar: "a"}, {bar: "bb"}, {bar: "ccc"}, {bar: "dddd"}}
	for i := 0; i < b.N; i++ {
		collect.Apply(input, func(foo TypeA) string {
			return foo.bar + "-suffix"
		})
	}
}

func BenchmarkApply_StructToStruct(b *testing.B) {
	b.ReportAllocs()
	input := []TypeA{{bar: "a"}, {bar: "bb"}, {bar: "ccc"}, {bar: "dddd"}}
	for i := 0; i < b.N; i++ {
		collect.Apply(input, func(foo TypeA) TypeB {
			return TypeB{baz: foo.bar + "-suffix"}
		})
	}
}
