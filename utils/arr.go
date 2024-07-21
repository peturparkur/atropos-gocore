// Package utils contains utility functions that are used throughout the application.
package utils

import "golang.org/x/exp/constraints"

// ArrContainsArr checks if an array contains all elements of another array
func ArrContainsArr[T constraints.Ordered](arr []T, subArr []T) bool {
	for _, a := range subArr {
		if !ArrContains(arr, a) {
			return false
		}
	}
	return true
}

// ArrContains checks if an array contains obj
func ArrContains[T constraints.Ordered](arr []T, obj T) bool {
	for _, a := range arr {
		if a == obj {
			return true
		}
	}
	return false
}
