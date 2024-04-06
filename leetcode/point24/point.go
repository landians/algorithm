package point24

import "math"

const (
	TARGET  = 24
	EPSILON = 1e-6
)

func judgePoint24(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	list := make([]float64, 0, len(nums))
	for _, num := range nums {
		list = append(list, float64(num))
	}
	return backtrack(list)
}

func backtrack(list []float64) bool {
	// 递归的出口，只剩 1 个数时，判断下是否够 24 点
	if len(list) == 1 {
		return math.Abs(list[0]-TARGET) < EPSILON
	}

	// flag 是用来减少 if 判断语句的
	flag := false

	// 开始选择
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			// 取出数组中两个不同的数, 并后续用这两个数的四则运算后的值当做一个新的数添加到新的数组，这样数组的数就少了一个
			n1, n2 := list[i], list[j]

			// 使用数组中剩下的数来构建新的数组
			newlist := make([]float64, 0, len(list))
			for k := 0; k < len(list); k++ {
				if k != i && k != j {
					newlist = append(newlist, list[k])
				}
			}

			// 加法
			flag = flag || backtrack(append(newlist, n1+n2))

			// 减法（减与被减）
			flag = flag || backtrack(append(newlist, n1-n2))
			flag = flag || backtrack(append(newlist, n2-n1))

			// 乘法
			flag = flag || backtrack(append(newlist, n1*n2))

			// 除法(除与被除)
			if n1 != 0 {
				flag = flag || backtrack(append(newlist, n2/n1))
			}
			if n2 != 0 {
				flag = flag || backtrack(append(newlist, n1/n2))
			}
			if flag {
				return true
			}
		}
	}
	return false
}
