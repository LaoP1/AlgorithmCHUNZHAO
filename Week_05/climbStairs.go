package climbStairs

//拆解之后可以发现，最终结果就是前一级台阶和前两级台阶的走法和，依据递推公式很简单可以看出是斐波那契数列，所以使用了一个数组做了循环的求取没有用递归
//时间复杂度：O(n)
//空间复杂度：O(n)
func climbStairs(n int) int {
    res := make([]int, n+1)
    res[0], res[1] = 1, 1
    for i:=2;i<=n;i++ {
        res[i] = res[i-1] + res[i-2]
    }
    return res[n]
}
