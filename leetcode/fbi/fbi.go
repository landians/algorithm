package fbi

func fbi1(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}

	return dp[n]
}

func fbi2(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	ni1, ni2 := 0, 1
	for i := 0; i < n; i++ {
		ni1, ni2 = ni2, (ni1+ni2)%1000000007
	}

	return ni1
}
