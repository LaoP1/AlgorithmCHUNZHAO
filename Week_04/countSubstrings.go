package countSubstrings

//查找回文中心并向两边拓展，统一了奇偶数长度两种情况
//时间复杂度：O(n^2)
//空间复杂度：O(1)
func countSubstrings(s string) int {
    n := len(s)
    ans := 0
    for i := 0; i < 2 * n - 1; i++ {
        l, r := i / 2, i / 2 + i % 2
        for l >= 0 && r < n && s[l] == s[r] {
            l--
            r++
            ans++
        }
    }
    return ans
}
