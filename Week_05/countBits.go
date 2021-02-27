package countBits

//可以参考汉明权重的方法，拿那个做一个子函数进行调用。或者使用动态规划和最低设置位进行循环
//时间复杂度：O(n)
//空间复杂度：O(n)
func countBits(num int) []int {
    res := make([]int, num+1)
    for i := 1; i <= num; i++ {
        res[i] = res[i & (i - 1)] + 1
    }
    return res
}
