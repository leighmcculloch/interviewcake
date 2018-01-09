// Package solution contains solutions to the problem described at https://www.interviewcake.com/question/balanced-binary-tree.
package solution

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func (n *BinaryTreeNode) InsertLeft(value int) *BinaryTreeNode {
	n.left = &BinaryTreeNode{value: value}
	return n.left
}

func (n *BinaryTreeNode) InsertRight(value int) *BinaryTreeNode {
	n.right = &BinaryTreeNode{value: value}
	return n.right
}

// Solution0 solves the problem in O(n) time and O(1) space, recursively.
func Solution0(tree *BinaryTreeNode) bool {
	minDepth, maxDepth := solution0Depth(tree)
	return maxDepth-minDepth <= 1
}

func solution0Depth(tree *BinaryTreeNode) (minDepth, maxDepth int) {
	if tree == nil {
		return 0, 0
	}

	leftMinDepth, leftMaxDepth := solution0Depth(tree.left)
	rightMinDepth, rightMaxDepth := solution0Depth(tree.right)

	if leftMinDepth < rightMinDepth {
		minDepth = leftMinDepth
	} else {
		minDepth = rightMinDepth
	}

	if leftMaxDepth > rightMaxDepth {
		maxDepth = leftMaxDepth
	} else {
		maxDepth = rightMaxDepth
	}

	return 1 + minDepth, 1 + maxDepth
}

// Solution1 solves the problem in O(n) time and O(1) space, iteratively.
func Solution1(tree *BinaryTreeNode) bool {
	// TODO
	return false
}
