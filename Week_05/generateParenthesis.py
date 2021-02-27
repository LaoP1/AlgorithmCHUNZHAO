from typing import List

# dfs+回溯算法，剪纸方法是左括号使用数量小于右括号使用数量
# 时间复杂度：O(4^n/n^0.5i)
# 空间复杂度：O(n)
class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        res = []
        cur_str = ''

        def dfs(cur_str, left, right, n):
            """
            :param cur_str: 从根结点到叶子结点的路径字符串
            :param left: 左括号已经使用的个数
            :param right: 右括号已经使用的个数
            :return:
            """
            if left == n and right == n:
                res.append(cur_str)
                return
            if left < right:
                return

            if left < n:
                dfs(cur_str + '(', left + 1, right, n)

            if right < n:
                dfs(cur_str + ')', left, right + 1, n)

        dfs(cur_str, 0, 0, n)
        return res
