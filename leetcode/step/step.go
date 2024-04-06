package step

func walkStep(target int) int {
	if target == 1 || target == 2 {
		return 1
	}

	return walkStep(target-1) + walkStep(target-2)
}

func stepDp(n int) int {
	if n < 1 {
		return 0
	}

	// 备忘录
	mem := make(map[int]int)

	return help(mem, n)
}

func help(mem map[int]int, n int) int {
	// base case
	if n == 1 || n == 2 {
		return 1
	}

	// 已经计算够了
	if _, ok := mem[n]; ok {
		return mem[n]
	}

	mem[n] = help(mem, n-1) + help(mem, n-2)
	return mem[n]
}

func stepDpUP(n int) int {
	if n < 1 {
		return 0
	}

	dp := make([]int, n+1)
	// base case
	dp[1], dp[2] = 1, 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}