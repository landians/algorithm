package nsum

import (
	"fmt"
	"sort"
)

// 两数之和: https://leetcode-cn.com/problems/two-sum/

// 不考虑原下标变化的解答
func twoSumArray(nums []int, target int) [][]int {
	fmt.Println("=== before sort ===")
	fmt.Println(nums)
	fmt.Println("=== before sort ===")

	// 先排序
	sort.Ints(nums)

	fmt.Println("=== after sort ===")
	fmt.Println(nums)
	fmt.Println("=== after sort ===")

	var sumArr [][]int

	// 通过左右双指针的方式来查找
	lo, hi := 0, len(nums)-1

	for lo < hi {
		left := nums[lo]
		right := nums[hi]
		// 根据 sum 和 target 的比较来决定 lo, hi 的指向
		sum := left + right
		if sum > target {
			hi--
		} else if sum < target {
			lo++
		} else {
			// 得到对应的值的集合
			sumArr = append(sumArr, []int{left, right})
			// 跳过重复的元素
			for lo < hi && nums[lo] == left {
				lo++
			}
			for lo < hi && nums[hi] == right {
				hi--
			}
		}
	}

	return sumArr
}

func twoSum(nums []int, target int) []int {
	// 构建 nums 组成的 hash 表
	numMap := make(map[int]int, len(nums))

	for i, num := range nums {
		if p, ok := numMap[target-num]; ok {
			return []int{p, i}
		} else {
			numMap[num] = i
		}
	}

	return []int{}
}

func twoSumTarget(nums []int, start int, end int, target int) [][]int {
	lo, hi := start, end
	var sumArr [][]int

	for lo < hi {
		left := nums[lo]
		right := nums[hi]

		// 根据 sum 和 target 的比较来决定 lo, hi 的指向
		sum := left + right
		if sum > target {
			hi--
		} else if sum < target {
			lo++
		} else {
			// 得到对应的值的集合
			sumArr = append(sumArr, []int{left, right})
			// 跳过重复的元素
			for lo < hi && nums[lo] == left {
				lo++
			}
			for lo < hi && nums[hi] == right {
				hi--
			}
		}
	}

	return sumArr
}

func threeSumArray(nums []int, target int) [][]int {
	fmt.Println("=== before sort ===")
	fmt.Println(nums)
	fmt.Println("=== before sort ===")

	// 先排序
	sort.Ints(nums)

	fmt.Println("=== after sort ===")
	fmt.Println(nums)
	fmt.Println("=== after sort ===")

	var sumArr [][]int

	for i := 0; i < len(nums); i++ {
		// 定下来一个数 nums[i]，然后求剩下的数两数之后为 target-nums[i] 的，即可得到三数之和
		twoSumArr := twoSumTarget(nums, i+1, len(nums)-1, target-nums[i])

		// 如果存在满足条件的二元组，再加上 nums[i] 就是结果三元组
		for _, sum := range twoSumArr {
			sum = append(sum, nums[i])
			sumArr = append(sumArr, sum)
		}

		//  跳过第一个数字重复的情况，否则会出现重复结果
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	return sumArr
}

// n 数之和的模板
func nSumTarget(nums []int, n int, start int, end int, target int) [][]int {
	lo, hi := start, end
	var sumArr [][]int

	if n == 2 { // twoSum 基本计算
		for lo < hi {
			left := nums[lo]
			right := nums[hi]

			// 根据 sum 和 target 的比较来决定 lo, hi 的指向
			sum := left + right
			if sum > target {
				hi--
			} else if sum < target {
				lo++
			} else {
				// 得到对应的值的集合
				sumArr = append(sumArr, []int{left, right})
				// 跳过重复的元素
				for lo < hi && nums[lo] == left {
					lo++
				}
				for lo < hi && nums[hi] == right {
					hi--
				}
			}
		}
	} else { // (n-1)Sum 递归计算
		for i := start; i <= end; i++ {
			subArr := nSumTarget(nums, n-1, i+1, len(nums)-1, target-nums[i])
			for _, arr := range subArr {
				arr = append(arr, nums[i])
				sumArr = append(sumArr, arr)
			}

			//  跳过第一个数字重复的情况，否则会出现重复结果
			for i < len(nums)-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}

	return sumArr
}

func fourSumArr(nums []int, target int) [][]int  {
	fmt.Println("=== before sort ===")
	fmt.Println(nums)
	fmt.Println("=== before sort ===")

	// 先排序
	sort.Ints(nums)

	fmt.Println("=== after sort ===")
	fmt.Println(nums)
	fmt.Println("=== after sort ===")

	return nSumTarget(nums, 4, 0, len(nums)-1, target)
}