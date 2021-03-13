# 无情的API选手
# 时间复杂度：O(N)
# 空间复杂度：O(N)
class Solution:
    def reverseWords(self, s: str) -> str:
        return " ".join(reversed(s.split()))
