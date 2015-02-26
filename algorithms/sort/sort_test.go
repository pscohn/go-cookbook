package sort

import "testing"

var numbers = []int{5, 4, 95, 45, 5, 4, 93, 194, 3844}

func TestMin(t *testing.T) {
	mindex := min(numbers)
	correct := 1
	if mindex != correct {
		t.Errorf("Result index %d does not match index %d\n", mindex, correct)
	}
}

func TestMax(t *testing.T) {
	maxindex := max(numbers)
	correct := 8
	if maxindex != correct {
		t.Errorf("Result index %d does not match index %d\n", maxindex, correct)
	}
}

func TestLess(t *testing.T) {
}

//func TestSwap(t *testing.T) {
//	num := []IntComp{1, 2, 3}
//	num.Swap(0, 2)
//	if num[0] != 3 || num[2] != 1 {
//		t.Errorf("Array %v swap fail")
//	}
//}

func TestInsertion(t *testing.T) {

}
