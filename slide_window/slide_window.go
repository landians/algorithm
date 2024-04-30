package slide_window

import "math"

// leetcode 3
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	window := make(map[byte]int, 1)
	l := 0

	lo, hi := 0, 0
	for hi < len(s) {
		in := s[hi]
		hi++
		window[in]++

		// window[in] > 1 说明窗口内有重复的子串了，可以收缩窗口
		for window[in] > 1 {
			out := s[lo]
			lo++
			window[out]--
		}

		l = int(math.Max(float64(hi-lo), float64(l)))
	}

	return l
}

// leetcode 17
func minWindow(s string, t string) string {
	// need 和 window 相当于计数器
	// need 记录了 t 中字符的出现次数
	// window 记录了 "窗口" 中相应字符的出现次数
	need := make(map[byte]int)
	window := make(map[byte]int)

	// 将需要判断的子串，使用 map 存储起来，key 为 字串的单个字符，value 为这个字符的 统计数量
	for _, b := range []byte(t) {
		need[b]++
	}

	// lo, hi 表示变量初始化窗口的两端，[lo, hi)
	// valid 表示窗口中满足 need 条件的字符个数，如果 valid 和 need.size 的大小相同，
	// 则说明窗口已经满足条件，已经完全覆盖了字符串 t
	lo, hi, valid := 0, 0, 0

	// i 记录了最小覆盖字串的起始索引，l 则记录了最小覆盖字串长度
	i, l := 0, math.MaxInt

	// 当加入字符 x 时，window[x]++
	// 当移出字符 x 时, window[x]--
	// 当 valid 满足 need 时，应该收缩窗口
	// 在收缩窗口的时候，应该更新最终结果

	for hi < len(s) {
		// in 表示将要移入窗口的字符
		in := s[hi]
		// 扩大窗口
		hi++
		// 进行窗口内数据的一系列更新
		if _, ok := need[in]; ok {
			// 加入窗口中
			window[in]++
			// 如果字符 in 在窗口中的数量已经满足其在字串 t 中的数量,  计数器 valid ++
			if window[in] == need[in] {
				valid++
			}
		}

		// 判断左侧窗口是否需要收缩， 如果滑动窗口中的字符已经完全覆盖字串 t 中的字符
		for valid == len(need) {
			// 如果此时的覆盖子串更短
			if (hi - lo) < l {
				// 更新最小覆盖子串的起始索引
				i = lo
				// 更新最小子串的长度
				l = hi - lo
			}

			// out 表示将要移出窗口的字符
			out := s[lo]
			// 缩小窗口
			lo++
			// 进行窗口内一系列数据的更新
			if _, ok := need[out]; ok {
				// 如果这个字符已经满足了他在字串 t 中的需求, 计数器 valid --
				if window[out] == need[out] {
					valid--
				}
				// 移出窗口
				window[out]--
			}
		}
	}

	if l == math.MaxInt {
		return ""
	} else {
		return s[i : i+l]
	}
}

// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引
func findAnagrams(s string, p string) []int {
	need := make(map[byte]int, len(p))
	window := make(map[byte]int, len(p))

	for _, b := range []byte(p) {
		need[b]++
	}

	anagrams := make([]int, 0)

	lo, hi, valid := 0, 0, 0

	for hi < len(s) {
		// in 表示将要移入窗口的字符
		in := s[hi]
		// 扩大窗口
		hi++
		// 进行窗口内数据的一系列更新
		if _, ok := need[in]; ok {
			// 加入窗口中
			window[in]++
			// 如果字符 in 在窗口中的数量已经满足其在字串 t 中的数量,  计数器 valid ++
			if window[in] == need[in] {
				valid++
			}
		}

		for hi-lo >= len(p) {
			if valid == len(need) {
				anagrams = append(anagrams, lo)
			}

			// out 表示将要移出窗口的字符
			out := s[lo]
			// 缩小窗口
			lo++
			// 进行窗口内一系列数据的更新
			if _, ok := need[out]; ok {
				// 如果这个字符已经满足了他在字串 t 中的需求, 计数器 valid --
				if window[out] == need[out] {
					valid--
				}
				// 移出窗口
				window[out]--
			}
		}
	}

	return anagrams
}

// 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int, len(s1))
	window := make(map[byte]int, len(s1))

	for _, b := range []byte(s1) {
		need[b]++
	}

	lo, hi, valid := 0, 0, 0

	for hi < len(s2) {
		// in 表示将要移入窗口的字符
		in := s2[hi]
		// 扩大窗口
		hi++
		// 进行窗口内数据的一系列更新
		if _, ok := need[in]; ok {
			// 加入窗口中
			window[in]++
			// 如果字符 in 在窗口中的数量已经满足其在字串 t 中的数量,  计数器 valid ++
			if window[in] == need[in] {
				valid++
			}
		}

		// 判断窗口是否需要进行收缩，判断是全排列，所以满足 t 的长度就是 ok
		for hi-lo >= len(s1) {
			if valid == len(need) {
				return true
			}

			// out 表示将要移出窗口的字符
			out := s2[lo]
			// 缩小窗口
			lo++
			// 进行窗口内一系列数据的更新
			if _, ok := need[out]; ok {
				// 如果这个字符已经满足了他在字串 t 中的需求, 计数器 valid --
				if window[out] == need[out] {
					valid--
				}
				// 移出窗口
				window[out]--
			}
		}
	}

	return false
}
