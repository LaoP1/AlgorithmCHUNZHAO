package nthUglyNumber

//合并三个有序数组，类似于动态规划
//时间复杂度：O(n)
//空间复杂度：O(n)
func nthUglyNumber(n int) int {
    dp := make([]int, n)
    p2, p3, p5 := 0, 0, 0
    dp[0] = 1
    for i:=1;i<n;i++ {
        dp[i] = min(min(dp[p2]*2, dp[p3]*3), dp[p5]*5)
        if dp[i]==dp[p2] * 2 {
            p2++
        }
        if dp[i]==dp[p3] * 3 {
            p3++
        }
        if dp[i]==dp[p5] * 5 {
            p5++
        }
    }
    return dp[n-1]
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}
