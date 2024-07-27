package calculate_pack

import (
	"reflect"
	"testing"
)

// Test cases

// Ordered 1, deliver 1 x 250
func TestCase1(t *testing.T) {
	ordered := 1
	var pack_sizes = []int{250, 500, 1000, 2000, 5000}
	exp := map[int]int{250: 1, 500: 0, 1000: 0, 2000: 0, 5000: 0}
	ans := CalculatePacks(ordered, pack_sizes)
	if !reflect.DeepEqual(exp, ans) {
		t.Errorf("Maps not equal")
	}
}

// Ordered 250, deliver 1 x 250
func TestCase2(t *testing.T) {
	ordered := 250
	var pack_sizes = []int{250, 500, 1000, 2000, 5000}
	exp := map[int]int{250: 1, 500: 0, 1000: 0, 2000: 0, 5000: 0}
	ans := CalculatePacks(ordered, pack_sizes)
	if !reflect.DeepEqual(exp, ans) {
		t.Errorf("Maps not equal")
	}
}

// Ordered 251, deliver 1 x 500
func TestCase3(t *testing.T) {
	ordered := 251
	var pack_sizes = []int{250, 500, 1000, 2000, 5000}
	exp := map[int]int{250: 0, 500: 1, 1000: 0, 2000: 0, 5000: 0}
	ans := CalculatePacks(ordered, pack_sizes)
	if !reflect.DeepEqual(exp, ans) {
		t.Errorf("Maps not equal")
	}
}

// Ordered 501, deliver 1 x 500, 1 x 250
func TestCase4(t *testing.T) {
	ordered := 501
	var pack_sizes = []int{250, 500, 1000, 2000, 5000}
	exp := map[int]int{250: 1, 500: 1, 1000: 0, 2000: 0, 5000: 0}
	ans := CalculatePacks(ordered, pack_sizes)
	if !reflect.DeepEqual(exp, ans) {
		t.Errorf("Maps not equal")
	}
}

// Ordered 12001, deliver 2 x 5000, 1 x 2000, 1 x 250
func TestCase5(t *testing.T) {
	ordered := 12001
	var pack_sizes = []int{250, 500, 1000, 2000, 5000}
	exp := map[int]int{250: 1, 500: 0, 1000: 0, 2000: 1, 5000: 2}
	ans := CalculatePacks(ordered, pack_sizes)
	if !reflect.DeepEqual(exp, ans) {
		t.Errorf("Maps not equal")
	}
}
