package preorderTraversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 //解法1：递归法
 //时间复杂度：O(n)
 //空间复杂度：O(n)
func preorderTraversal(root *TreeNode) []int {
    if root==nil {
        return nil
    }
    res := []int{}
    res = append(res, root.Val)
    res = append(res, preorderTraversal(root.Left)...)
    res = append(res, preorderTraversal(root.Right)...)
    return res
}

//解法2：迭代法
//时间复杂度：O(n)
//空间复杂度：O(N)
func preorderTraversal(root *TreeNode) []int {
    //边界条件处理
    if root == nil {
        return nil
    }

    //局部变量声明
    res := []int{}
    stack := []*TreeNode{root}
    marks := make(map[*TreeNode]bool)

    //主要循环
    for len(stack) !=0 {
        i := len(stack) - 1
        node := stack[i]
        stack = stack[:i]

        _, ok := marks[node]
        if ok {
            res = append(res, node.Val)
            //如果节点已经遍历过则输出节点值
        } else {
            //如果节点未曾被遍历过则入栈该节点和右、左孩子，并修改状态

            //修改此行代码位置实现不同遍历
            //位置1，实现后序遍历

            if node.Right != nil {
                stack = append(stack, node.Right)
            }

            //位置2，实现中序遍历

            if node.Left != nil {
                stack = append(stack, node.Left)
            }

            //位置3，实现前序遍历
            stack = append(stack, node)

            marks[node] = true
        }
    }
    return res
}
