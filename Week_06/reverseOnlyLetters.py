# 方法1：栈
# 时间复杂度：O(N)
# 空间复杂度：O(N)
# 将字符串中的字符压入栈中，再次遍历S，确保非字符位置不变，若是字符则获取栈顶元素
class Solution:
    def reverseOnlyLetters(self, S: str) -> str:
        stack = [c for c in S if c.isalpha()]
        ans = []
        for x in S:
            if not x.isalpha():
                ans.append(x)
            else:
                ans.append(stack.pop())
        
        return "".join(ans)



# 方法2：反转指针
# 时间复杂度：O(N)
# 空间复杂度：O(N)
# 方法二实际上就是栈方法的一种改写，去掉了栈，用时间换空间。做了一次反向遍历
class Solution:
    def reverseOnlyLetters(self, S: str) -> str:
        ans = []
        j = len(S)-1
        for i, x in enumerate(S):
            if x.isalpha():
                while not S[j].isalpha():
                    j-=1
                ans.append(S[j])
                j-=1
            else:
                ans.append(x)
        return "".join(ans)
