# dp方法
# 时间复杂度：O(N^2)
# 空间复杂度：O(N)
# 目前只能比较好的理解dp方法，dp的再优化方法还正在学习
class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        if not nums: return 0
        dp = [1]*len(nums)
        for i in range(len(nums)):
            for j in range(i):
                if nums[i] > nums[j]:
                    dp[i] = max(dp[i], dp[j]+1)
        return max(dp)
