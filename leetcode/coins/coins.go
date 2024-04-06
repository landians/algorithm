package coins

// 给你k种面值的硬币，面值分别为c1, c2 ... ck，每种硬币的数量无限，再给一个总金额amount，
// 问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	// 备忘录
	mem := make(map[int]int)

	return dp(amount, coins, mem)
}

func dp(n int, coins []int, mem map[int]int) int {
	if _, ok := mem[n]; ok {
		return mem[n]
	}

	// base case
	if n == 0 {
		return 0
	}

	if n < 0 {
		return -1
	}

	res := 100000000

	for _, coin := range coins {
		// 子问题求解
		subN := dp(n-coin, coins, mem)
		// 换一种硬币试试
		if subN == -1 {
			continue
		}

		res = min(res, 1+subN)
	}

	if res != 100000000 {
		mem[n] = res
	} else {
		mem[n] = -1
	}

	return mem[n]
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
