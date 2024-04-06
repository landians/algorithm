package slide

import (
	"math"
)

// 最小覆盖子串问题: https://leetcode-cn.com/problems/minimum-window-substring/
func minWindow(s string, t string) string {
	// 1. 初始化窗口
	lo, hi := 0, 0

	// needs和window相当于计数器，分别记录T中字符出现次数和「窗口」中的相应字符的出现次数, 因为 T 中的字符串是可能重复的
	need, window := make(map[byte]int), make(map[byte]int)

	// valid变量表示窗口中满足need条件的字符个数，如果valid和need.size的大小相同，则说明窗口已满足条件，已经完全覆盖了串T。
	valid := 0

	// 记录最小覆盖子串的起始索引及长度
	start, length := 0, math.MaxInt

	byteS := []byte(s)
	byteT := []byte(t)

	for _, bt := range byteT {
		need[bt]++
	}

	// 开始滑动
	for hi < len(s) {
		// c 是将移入窗口的字符
		c := byteS[hi]
		// 右移窗口
		hi++

		// 进行窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		//判断左侧窗口是否要收缩
		for valid == len(need) {
			// 在这里更新最小覆盖子串
			if hi-lo < length {
				start = lo
				length = hi - lo
			}

			// d 是将移出窗口的字符
			d := byteS[lo]
			// 左移窗口
			lo++

			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	// 返回最小覆盖子串
	if length == math.MaxInt {
		return ""
	}
	return string(byteS[start:length])

	// 2. 不断增大窗口，即 hi++, 直到窗口中的字符串包含了 t 这个子串

	// 3. 不断缩小窗口，即 lo++, 直到窗口中的字符串不包含 t 这个子串

	// 4. 重复 2，3 步骤，直到 hi 走到尽头
}
