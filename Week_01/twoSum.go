package twoSum

func twoSum(nums []int, target int) []int {
    marks := make(map[int]int)
    res := make([]int, 0)
    for i:=0;i<len(nums);i++ {
        tmp := target - nums[i]
        if _, ok := marks[tmp]; ok==true {
            return append(res, i, marks[tmp])
        } else {
            marks[nums[i]] = i
        }
    }
    return append(res, -1, -1)
}
