// Package collect provides some utilities for slices and maps.
package collect

import (
	"maps"
	"slices"
)

// Apply returns a new []V created by collecting the result of applying the given
// function to each element in the given slice.
func Apply[Slice ~[]T, T any, V any](s Slice, applyFn func(T) V) []V {
	out := make([]V, len(s))
	for idx, t := range s {
		out[idx] = applyFn(t)
	}
	return out
}

// Select returns a new []T created by applying the given function to each
// element, and collecting the elements where the function returns true.
func Select[Slice ~[]T, T any](s Slice, include func(T) bool) (out []T) {
	for _, t := range s {
		if include(t) {
			out = append(out, t)
		}
	}
	return out
}

// Reject returns a new []T created by applying the given function to each
// element, and collecting the elements where the function returns false.
func Reject[Slice ~[]T, T any](s Slice, exclude func(T) bool) (out []T) {
	for _, t := range s {
		if !exclude(t) {
			out = append(out, t)
		}
	}
	return out
}

// Keys returns a slice containing the keys from the given map. This is
// different from the stdlib maps.Keys() in that the keys are materialized into
// a slice rather than returned as an iter.Seq.
func Keys[Map ~map[K]V, K comparable, V any](m Map) []K {
	return slices.Collect(maps.Keys(m))
}

// Values returns a slice containing the values from the given map. This is
// different from the stdlib maps.Values() in that the values are materialized into
// a slice rather than returned as an iter.Seq.
func Values[Map ~map[K]V, K comparable, V any](m Map) []V {
	return slices.Collect(maps.Values(m))
}
