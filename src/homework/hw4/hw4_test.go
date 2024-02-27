package main

import "testing"

func TestIsPower(t *testing.T) {

	assert := func(x, y int, expected bool) {
		if is_power(x, y) != expected {
			t.Errorf("is_power(%v, %v) == %v, expected %v", x, y, !expected, expected)
		}
	}

	assert(8, 2, true)
	assert(81, 3, true)
	assert(64, 3, false)
	assert(1, 9, true)
	assert(3, 9, false)
}

func TestListReverse(t *testing.T) {
}

func TestListPrimes(t *testing.T) {
}

func TestFilterList(t *testing.T) {
}

func TestListOverlap(t *testing.T) {
}
