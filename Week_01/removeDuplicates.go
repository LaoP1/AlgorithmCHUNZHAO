package removeDuplicates

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    p := 0
    for q := 1; q<len(nums);q++ {
        if nums[p] != nums[q] {
            nums[p+1] = nums[q]
            p++
        }
    }
    return p+1
}
