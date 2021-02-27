package hammingWeight

//解法是采取位运算操作，保持num位数为32，采用x & x-1的操作去除低位的0，只要x不为0就计数一次，num为0退出循环
//时间复杂度：O(n)
//空间复杂度：O(1)
func hammingWeight(num uint32) int {
    count := 0
    for num!=0 {
        count++
        num &= (num - 1)
    }
    return count
}
