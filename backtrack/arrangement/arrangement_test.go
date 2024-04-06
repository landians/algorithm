package arrangement

import "testing"

func Test_Permute(t *testing.T) {
	arr := []int{1, 2, 3}
	arrange := permute(arr)
	for _, v := range arrange {
		t.Log(v)
	}
}
