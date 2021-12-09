package main

import "fmt"

// 只要涉及递归，都可以抽象成二叉树的问题
// 写递归算法的关键是要明确函数的「定义」是什么，然后相信这个定义，利用这个定义推导最终结果，绝不要跳入递归的细节。
func traverse(root *TreeNode) {
	// 前序遍历
	// 如果按照 根节点 -> 左节点 -> 右结点的方式遍历,叫先序遍历
	traverse(root.Left)

	// 中序遍历
	// 如果按照 左结点 -> 根节点 -> 右结点的方式遍历,叫中序遍历
	traverse(root.Right)
	// 后序遍历
	// 如果按照 左结点 -> 右结点 -> 根结点的方式遍历,叫后序遍历
}

// 二叉树算法题
// 	   4
//   /   \
//  2     7
// / \   / \
//1   3 6   9

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}

func main() {
	var one = TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	var three = TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	var two = TreeNode{
		Val:   2,
		Left:  &one,
		Right: &three,
	}

	var six = TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}

	var nine = TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}

	var seven = TreeNode{
		Val:   7,
		Left:  &six,
		Right: &nine,
	}

	var head = TreeNode{
		Val:   4,
		Left:  &two,
		Right: &seven,
	}

	//fmt.Println(mirrorTree(&head))
	recursionMiddleorderTraversal(&head)
}

func count(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + count(node.Left) + count(node.Right)
}

// 翻转二叉树
func invertTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	var lin = node.Left
	node.Left = node.Right
	node.Right = lin
	invertTree(node.Left)
	invertTree(node.Right)
	return node
}

// 连接两个二叉树
// 思路,这个题的重点是,要把跨越父节点的两个子节点相互连接起来
func connect(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

func connectTwoNode(node1 *TreeNode, node2 *TreeNode) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	// 连接相同父节点的两个子节点
	connectTwoNode(node1.Left, node1.Right)
	connectTwoNode(node2.Left, node2.Right)
	// 连接跨越父节点的两个子节点
	connectTwoNode(node1.Right, node2.Left)
}

// 二叉树展开为链表
// 思路:使用后序遍历 左结点 -> 右结点 -> 根结点

func flatten(node *TreeNode) {
	if node == nil {
		return
	}
	flatten(node.Left)
	flatten(node.Right)

	// 后序遍历
	var left = node.Left
	var right = node.Left
	node.Left = nil
	node.Right = left
	var p = node
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}

// 递归中序遍历
func recursionMiddleorderTraversal(root *TreeNode) {
	if root != nil {
		recursionMiddleorderTraversal(root.Left)
		fmt.Println(root.Val)
		recursionMiddleorderTraversal(root.Right)
	}
}

//输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
func isBalanced(root *TreeNode) bool {
	if maxDepth(root) == -1 {
		return false
	}
	return true
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	// 为什么返回-1呢？（变量具有二义性）
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

// 输入一个二叉树,翻转它
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := mirrorTree(root.Left)
	right := mirrorTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

// 请 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	return root
}
