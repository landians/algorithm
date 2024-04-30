package array

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			// 维护 nums[0..slow] 无重复
			nums[slow] = nums[fast]
		}
		fast++
	}

	return slow + 1
}

// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return []int{-1, -1}
	}

	lo, hi := 0, len(numbers)-1
	for lo <= hi {
		sum := numbers[lo] + numbers[hi]
		switch {
		case sum > target:
			hi--
		case sum == target:
			return []int{lo + 1, hi + 1}
		case sum < target:
			lo++
		}
	}

	return []int{-1, -1}
}
