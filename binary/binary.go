package binary

/*
二分法可以处理的问题:
1. 在一个有序数组中, 找 == 某个数存在的位置
2. 在一个有序数组中, 找 >= 某个数最左侧的位置
3. 在一个有序数组中, 找 <= 某个数最右侧的位置
4. 局部最小值问题, 这个不需要有序!

时间复杂度为: O(logN)
二分查找在最坏的情况下依次是n/2,n/4,n/8... 一直到1为止, 时间复杂度就指的是查询的次数
假设查询的次数为 x, 那么就可以得到这个公式: n(1/2)^x = 1 => n/2^x = 1 => 2^x = n => x = log2(N),
忽略底数可知, 二分查找的时间复杂度为 O(logN)

注: go标准库中的二分查找函数是: sort.Search
*/

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	i := FindLeftBound(nums, target)
	if i == len(nums) || nums[i] != target {
		return []int{-1, -1}
	}

	j := FindLeftBound(nums, target+1)

	return []int{i, j - 1}
}

func FindLeftBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	f := func(i int) bool {
		return nums[i] >= target
	}

	i := binarySearch(len(nums), f)

	// i == len(nums) || nums[i] != target 则认为是找不到

	return i
}

func binarySearch(n int, f func(int) bool) int {
	lo, hi := 0, n
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if !f(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
