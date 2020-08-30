package leetcode_go

/*
有序链表转二叉搜索树
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
*/

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func sortedListToBST(head *ListNode) *TreeNode2 {
	return buildTree(head, nil)
}

func getMedian(left, right *ListNode) *ListNode {
	fast, slow := left, left
	for fast != right && fast.Next != right {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func buildTree(left, right *ListNode) *TreeNode2 {
	if left == right {
		return nil
	}
	mid := getMedian(left, right)
	root := &TreeNode2{
		Val:   mid.Val,
		Left:  nil,
		Right: nil,
	}
	root.Left = buildTree(left, mid)
	root.Right = buildTree(mid.Next, right)
	return root
}
