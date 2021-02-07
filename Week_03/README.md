学习笔记

### 作业题目
1. [多数元素](https://leetcode-cn.com/problems/majority-element/)
2. [买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)
3. [子集](https://leetcode-cn.com/problems/subsets/)
4. [搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)
5. [寻找旋转排序数组中的最小值](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)
6. [N 皇后](https://leetcode-cn.com/problems/n-queens/)

---

### 学习总结
* 题目模式识别，就是找到一个拐点。其左右两边是两种趋势，一边单调递增、一边单调递减。直接的方案就是进行一次数组遍历，然后判断下标前后变化是否单调性一致
```Golang
func search(nums []int) int {
	left, right := 0, len(nums)-1;
	for left < right {
		mid := (right - left)/2 + left
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
```

