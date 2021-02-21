package canCross

//记忆化递归方法，参考的LC题解，但是想翻译一个Java动归到Go有点不太行
//时间复杂度：O(n^3)
//空间复杂度：O(n^2)
func canCross(stones []int) bool {
	hashMap := map[int]bool{} // 构建空map
	return helper(stones, 0, 0, hashMap) // 上一步跳了0步，现在在索引0处，基于此，能否到终点
}
// helper：上一步跳了k步，来到index处，基于此，能否到达终点
func helper(stones []int, index, k int, hashMap map[int]bool) bool {
    key := index*1000 + k 	// 构造一个唯一的key，代表当前子问题
    if hashMap[key] { 	// 如果这个子问题之前遇到过，直接返回false
		return false 		// 因为都第二次遇到相同子问题了，说明之前的子问题肯定返回false
	} else {    			// 第一次遇到这个子问题，在hashMap记录一下
		hashMap[key] = true
	}
	for i := index + 1; i < len(stones); i++ { //枚举出可跳的选择，从下一个石头到最后一个
		gap := stones[i] - stones[index]		  // 本轮迭代选择跳到i，算出第i石头到当前石头的距离
		if gap >= k-1 && gap <= k+1 {    	  // 这一步能跳[k-1,k+1]，距离gap在这范围内，就能尝试
			if helper(stones, i, gap, hashMap) { // 如果 基于此的递归返回true，说明能到终点
				return true					  // 递归压栈压到底，如果true了，则true一层向上返回
			}
		} else if gap > k+1 { // 超出范围，跳不到，它之后的石头更跳不到，不用考察了，break
			break
		} // 这是gap<k-1的情况，说明这个点离太近了，要继续下一轮迭代，看看远一点的石头
	}
    return index == len(stones)-1 // 考察完当前的所有选项，看看来到的index是不是终点
}
