package nsum

import "testing"

func Test_TwoSumArray(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 3, 3}
	sumArr := twoSumArray(nums, 4)
	for _, arr := range sumArr {
		t.Log(arr)
	}
}

func Test_ThreeSumArray(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 3, 3}
	sumArr := threeSumArray(nums, 6)
	for _, arr := range sumArr {
		t.Log(arr)
	}
}

func Test_FourSumArray(t *testing.T) {
	nums := []int{2, 2, 2, 2, 1, 3, 3}
	sumArr := fourSumArr(nums, 8)
	for _, arr := range sumArr {
		t.Log(arr)
	}
}