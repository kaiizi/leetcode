/*
leetcode 300.最长递增子序列
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

输入：nums = [0,1,0,3,2,3]
输出：4
*/

package main

import "fmt"

// 动态规划, 最优子结构：dp[i]= max(dp[j]) + 1（j属于[0, i-1],并且num[i]>num[j]）
// 可优化：增加二分查找
func lengthOfLIS(nums []int) int {
	// 边界
	dp := make([]int, len(nums))
	dp[0] = 1
	// 构建最优子结构
	max := func(index int) int {
		maxValue := 0
		for i := 0; i <= index; i++ {
			if nums[index+1] > nums[i] && dp[i] > maxValue {
				maxValue = dp[i]
			}
		}
		return maxValue
	}
	// 自顶向上计算最优子结构
	for i := 1; i < len(nums); i++ {
		//dp = append(dp, max(i - 1) + 1)
		dp[i] = max(i-1) + 1
	}
	// 返回计算结果
	result := 1
	for _, v := range dp {
		if v > result {
			result = v
		}
	}

	return result
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}
