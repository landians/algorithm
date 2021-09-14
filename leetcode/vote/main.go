package main

import "fmt"

/*
 * 给定一个整型数组arr，长度为N
 * 如果存在某个数，其出现次数大于N/2，返回这个数
 * 如果不存在这样的数，返回-1
 *
 * 要求：时间复杂度O(N)，额外空间复杂度O(1)
 *
 * */

func vote(arr []int) (i int) {
	if len(arr) == 0 {
		return -1
	}

	// candidate 表示当前候选人的选票编号，为 0 表示没有候选人
	candidate := 0
	// restHP 表示候选人的当选几率
	restHP := 0

	for _, cur := range arr {
		if candidate == 0 { // 如果没有候选
			candidate = cur
			restHP = 1
		} else if cur == candidate { // 如果有候选，并且当前的数字和候选不一样
			restHP++
		} else { // 如果有候选，并且当前的数字和候选一样
			restHP--
		}
	}

	// // 如果遍历完成后，没有候选留下来，说明没有符合条件的数
	if restHP <= 0 {
		return -1
	}

	// 如果有候选留下来，再去遍历一遍，得到候选真正出现的次数
	count := 0
	for _, v := range arr {
		if v == candidate {
			count++
		}
	}

	// 如果候选真正出现的次数大于N/2，返回候选，否则说明没有符合条件的数
	if count > (len(arr) >> 1) {
		return candidate
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(vote(arr))

	arr = []int{2, 1, 2, 4, 1, 2, 2}
	fmt.Println(vote(arr))
}
