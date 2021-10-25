package dp

import "fmt"

// 01背包
func bag01(items [][]int, cap int) int {
	itemsLen := len(items)

	// 初始化 行
	dp := make([][]int, itemsLen)

	// 初始化 列
	for i := 0; i < itemsLen; i++ {
		dp[i] = make([]int, cap+1)
	}

	// 初始化 第一行 （价值）
	for j := items[0][0]; j <= cap; j++ {
		dp[0][j] = items[0][1]
	}

	for i := 1; i < itemsLen; i++ {
		for j := 0; j <= cap; j++ {
			if j >= items[i][0] {
				// TODO
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-items[i][0]]+items[i][1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	fmt.Println(dp)

	return dp[itemsLen-1][cap]

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
