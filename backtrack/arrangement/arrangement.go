package arrangement

// 使用回溯算法来解决全排列的问题: 对于 n 个不重复的数字，全排列共有 n! = n * n-1 * n-2 * ... * 1
// for 选择 range 选择列表:
//    // 做选择
//    将该选择从选择列表移除
//    路径.add(选择)
//    backtrack(路径, 选择列表)
//    // 撤销选择
//    路径.remove(选择)
//    将该选择再加入选择列表
// 递归函数的事件复杂度 = 递归函数本身的时间复杂度 * 递归函数调用的次数

// 输入一组不重复的数字, 返回它的全排列

func permute(arr []int) [][]int {
	// track 用于记录选择的「路径」
	track := make([]int, 0, len(arr))

	// arrange 用于保存最终的排列结果
	arrange := make([][]int, 0)

	// backtrack 使用回溯法来获取全排列
	backtrack(arr, track, &arrange)

	return arrange
}

// 路径：记录在 track 中
// 选择列表：arr 中不存在于 track 的那些元素
// 结束条件：arr 中的元素全都在 track 中出现
func backtrack(arr []int, track []int, arrange *[][]int) {
	if len(track) == len(arr) {
		// 把临时结果复制出来保存到最终结果
		ans := make([]int, len(track))
		copy(ans, track)
		*arrange = append(*arrange, ans)
		return
	}

	for _, v := range arr {
		if contains(track, v) {
			continue
		}

		// 做选择
		track = append(track, v)

		// 进入下一层决策树
		backtrack(arr, track, arrange)

		// 取消选择
		track = track[:len(track)-1]
	}
}

func contains(arr []int, v int) bool {
	for _, item := range arr {
		if item == v {
			return true
		}
	}
	return false
}
