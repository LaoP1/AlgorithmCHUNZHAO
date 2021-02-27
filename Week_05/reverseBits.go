package reverseBits

//使用最简单的循环置换，新开数组记录结果。使用num>>1改变num，num=0时跳出循环，使用num&1取出num最右位，(nums&1)<<power调转位置
//时间复杂度：O(log 2 n)
//空间复杂度：O(1)
func reverseBits(num uint32) uint32 {
    res := uint32(0)
    power := uint32(31)
    for num != 0 {
        res += (num & 1) << power
        num = num >> 1
        power -= 1
    }
    return res
}
