package goslice

import (
	"golang.org/x/exp/slices"
)

// ContainsSet checks if a slice contains all of provided values
func ContainsSet[T comparable](slice []T, s ...T) bool {
	if len(slice) == 0 {
		return false
	}
	if len(s) == 0 {
		return false
	}
	for _, v := range s {
		if !slices.Contains(slice, v) {
			return false
		}
	}
	return true
}
