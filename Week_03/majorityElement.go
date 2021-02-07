package majorityElement

//投票法
//时间复杂度：O(n)
//空间复杂度：O(1)
func majorityElement(nums []int) int {
    count,candidate := 0, ^int(0)

    for _, num := range nums {
        if count==0 {
            candidate = num
        }
        if num == candidate {
            count += 1
        } else {
            count -= 1
        }
    }

    return candidate
}
