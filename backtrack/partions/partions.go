package partions

// 划分为k个相等的子集. https://leetcode-cn.com/problems/partition-to-k-equal-sum-subsets/
// 视角一，如果我们切换到这 n 个数字的视角，每个数字都要选择进入到 k 个桶中的某一个。
// 视角二，如果我们切换到这 k 个桶的视角，对于每个桶，都要遍历 nums 中的 n 个数字，然后选择是否将当前遍历到的数字装进自己这个桶里。

func canPartitionKSubsets(nums []int, k int) bool {
	// 如果 nums 的元素总和 % k != 0, 那么肯定不能划分
	total := sum(nums)
	if total%k != 0 {
		return false
	}

	// 记录了每一个数是否已经划分到一个桶里面去了
	used := make([]bool, len(nums))

	// 得到每个桶中数据的和
	target := total / k

	return backtrack(k, 0, nums, 0, used, target)
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func backtrack(k int, bucketTarget int, nums []int, index int, used []bool, target int) bool {
	// base case
	if k == 0 {
		return true
	}

	// 当前桶已经装满了，开始选择下一个桶
	if bucketTarget == target {
		return backtrack(k-1, 0, nums, 0, used, target)
	}

	// 从 index 开始向后探查有效的 nums[i] 装入当前桶
	for i := index; i < len(nums); i++ {
		// 剪枝: nums[i] 已经被装入别的桶中
		if used[i] {
			continue
		}

		// 当前桶装不下 nums[i]
		if nums[i]+bucketTarget > target {
			continue
		}

		// 做选择，将 nums[i] 装入当前桶中
		used[i] = true
		bucketTarget += nums[i]

		// 递归穷举下一个数字是否装入当前桶
		if backtrack(k, bucketTarget, nums, i+1, used, target) {
			return true
		}

		// 撤销选择
		used[i] = false
		bucketTarget -= nums[i]

	}

	// 穷举了所有数字，都无法装满当前桶
	return false
}
