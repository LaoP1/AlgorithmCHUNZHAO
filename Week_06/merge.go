package merge

// 排序后合并
// 时间复杂度：O(NlogN)
// 空间复杂度：O(logN)
// 主要时空复杂度都在排序算法上
func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    res := [][]int{}
    prev := intervals[0]

    for i:=1;i<len(intervals);i++ {
        cur := intervals[i]
        if prev[1] < cur[0] {
            res = append(res, prev)
            prev = cur
        } else {
            prev[1] = max(prev[1], cur[1])
        }
    }
    res = append(res, prev)
    return res
}

func max(a, b int) int {
    if a > b { return a}
    return b
}
