package isPowerOfTwo

//采取位运算的方式，采取去掉最高位1检测n是否为0
//时间复杂度：O(1)
//空间复杂度：O(1)
func isPowerOfTwo(n int) bool {
    if n==0 {
        return false
    }
    return n&(n-1)==0
}
