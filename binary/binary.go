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


func search(arr []int, f func(int) bool) int {
	if len(arr) == 0 {
		return -1
	}
	lo, hi := 0, len(arr)-1

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

// 二分查找, 针对的是升序数组
func BinarySearch_1(arr []int, value int) int {
	f := func(i int) bool {
		return arr[i] >= value
	}
	return search(arr, f)
}

// 二分查找, 针对的是降序数组
func BinarySearch_2(arr []int, value int) int {
	f := func(i int) bool {
		return arr[i] <= value
	}
	return search(arr, f)
}

/* 
局部最小值定义:
arr[i-1], arr[i], arr[i+1], 若 arr[i] < arr[i-1] 同时 arr[i] < arr[i+1], 则 arr[i] 就是局部最小值
arr[N-2], arr[N-1], 若 arr[N-2] < arr[N-1], 则 arr[N-2] 就是局部最小值

局部最小值问题:
现有数组 arr[0 ~ N-1], 数组中的每相邻两个数都是不相等的, 需要返回其中一个局部最小值的下标

如果一个数组中的值的变化趋势如下, 因为每相邻两个数都是不相等的的原因, 可以判断数组中间一定存在局部最小值
------------\          /--------------
             \        /
arr[0]   arr[1] .... arr[N-2] arr[N-1]  
*/

// 搜索左边界
func findLeftBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	lo, hi := 0, len(nums)

	for lo < hi {
		mid := lo + (hi - lo) / 2
		if nums[mid] < target {
			lo = mid+1
		} else {
			hi = mid
		}
	}
	return lo
}

// 搜索右边界
func findRightBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	lo, hi := 0, len(nums)

	for lo < hi {
		mid := lo + (hi - lo) / 2
		if nums[mid] > target {
			hi = mid
		} else {
			lo = mid+1
		}
	}
	return lo
}