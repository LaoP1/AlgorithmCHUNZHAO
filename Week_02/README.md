学习笔记

选择的题目：
[二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
[二叉树的前序遍历](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)
[丑数](https://leetcode-cn.com/problems/chou-shu-lcof/)
[从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)
[全排列](https://leetcode-cn.com/problems/permutations/)
[全排列 II ](https://leetcode-cn.com/problems/permutations-ii/)

---

###  二叉树
1. 遍历
   * 前序遍历(preorder)：根-左-右
   * 中序遍历(inorder)：左-根-右；通常来说，对于二叉搜索树，我们可以通过中序遍历得到一个递增的有序序列。
   * 后序遍历(postorder)：左-右-根；删除树中的节点时，删除过程将按照后序遍历的顺序进行；后序遍历在数学表达中被广泛使用
   * 遍历的模板：其实有四种但是我只觉得先学两种吧
      1. 递归
         ```Golang
         func traverse(root *TreeNode) []int {
             if root == nil {
                 return nil
             }
             res := []int{}
             res = append(res, (*root).Val)
             res = append(res, traverse(root.Left)...)
             res = append(res, traverse(root.Right)...)
             return res
         }
         ```
      2. 迭代（stack）
         ```Golang
         func iterationTraverse(root *TreeNode) []int {
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
                     stack = append(stack, node) //位置1，实现后序遍历
                     
                     if node.Right != nil {
                         stack = append(stack, node.Right)
                     }
                     
                     //位置2，实现中序遍历
                     
                     if node.Left != nil {
                         stack = append(stack, node.Left)
                     }
                     
                     //位置3，实现前序遍历
                     
                     marks[node] = true
                 }
             }
             return res
         }
         ```
2. 重建二叉树
	
	* 通过前序/后序遍历+中序遍历可重建二叉树