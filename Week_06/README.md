学习笔记

## 作业
1. [仅仅反转字母](https://leetcode-cn.com/problems/reverse-only-letters/)
2. [LRU缓存机制](https://leetcode-cn.com/problems/lru-cache/#/)
3. [最长上升子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/)
4. [翻转对](https://leetcode-cn.com/problems/reverse-pairs/)
5. [合并区间](https://leetcode-cn.com/problems/merge-intervals/)
6. [翻转字符串里的单词](https://leetcode-cn.com/problems/reverse-words-in-a-string/)

## 学习总结
Go语言实现简单排序算法
1. 冒泡排序
	```Go
	func bubbleSort(slice []int) []int {
       for n := 0; n <= len(slice); n++ {
          for i := 1; i < len(slice)-n; i++ {
             if slice[i] < slice[i-1] {
                slice[i], slice[i-1] = slice[i-1], slice[i]
             }
          }
       }
       return slice
    }
	```

2. 比较排序
	```Go
	func selectionSort(slice []int) []int {
       for n := 0; n <= len(slice); n++ {
          for i := n + 1; i < len(slice); i++ {
             if slice[n] > slice[i] {
                slice[n], slice[i] = slice[i], slice[n]
             }
          }
       }
       return slice
    }
	```

3. 插入排序
	```Go
	func insertSort(arr []int) {
       for n := 1; n < len(arr); n++ {
          for i := n; i > 0; i-- {
             if arr[i] > arr[i-1] {
                break
             }
             arr[i], arr[i-1] = arr[i-1], arr[i]
          }
       }
    }
	```

4. 快速排序
	```Go
	func qsort(arr []int, start int, end int) {
       if start >= end {
          return
       }

       key := start
       value := arr[start] 
       for n := start + 1; n <= end; n++ {
          if arr[n] < value {
             arr[key] = arr[n]   
             arr[n] = arr[key+1] 
             key++ 

          }
       }

       arr[key] = value
       qsort(arr, start, key-1)
       qsort(arr, key+1, end) 
}
	```