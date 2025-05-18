package main

import "fmt"

// 01 背包问题：给定 n 个物品和一个容量为 W 的背包，每个物品有一个重量 weight 和价值 value
// 问：如何选择物品装入背包，使得总重量不超过 W 的情况下，总价值最大

// 使用动态规划解决 01 背包问题
// items：物品数组，每个元素包含物品的重量和价值
// capacity：背包容量
// 返回最大总价值
func knapsack01(items []struct{ weight, value int }, capacity int) int {
	n := len(items)
	if n == 0 || capacity <= 0 {
		return 0
	}

	// 创建动态规划表，dp[i][w] 表示前 i 个物品放入容量为 w 的背包中的最大价值
	// 为了节省空间，可以优化为一维数组
	dp := make([]int, capacity+1)

	// 遍历每个物品
	for _, item := range items {
		weight := item.weight
		value := item.value

		// 从后往前更新 dp 数组，避免重复使用物品
		for w := capacity; w >= weight; w-- {
			if dp[w-weight]+value > dp[w] {
				dp[w] = dp[w-weight] + value
			}
		}
	}

	return dp[capacity]
}

func main() {
	// 测试用例
	items := []struct{ weight, value int }{
		{2, 3},
		{3, 4},
		{4, 8},
		{5, 8},
		{9, 10},
	}

	capacity := 20

	maxValue := knapsack01(items, capacity)
	fmt.Printf("最大总价值为: %d\n", maxValue)
}
