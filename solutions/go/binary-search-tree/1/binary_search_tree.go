package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{left: nil, data: i, right: nil}
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
	if bst == nil {
        bst = &BinarySearchTree{left: nil, data: i, right: nil}
        return
    }

    if (i <= bst.data) {
        if bst.left == nil {
            bst.left = NewBst(i)
        } else {
            bst.left.Insert(i)
        }
    } else {
        if bst.right == nil {
            bst.right = NewBst(i)
        } else {
            bst.right.Insert(i)
        }
    }
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	if bst == nil {
        return []int{}
    }
    
    count := count_nodes(bst)
    result := make([]int, count)
    index := 0
    in_order_traversal(bst, result, &index)
    return result
}

func in_order_traversal(bst *BinarySearchTree, result []int, index *int) {
    if bst == nil {
        return
    }

    in_order_traversal(bst.left, result, index)
    result[*index] = bst.data
    *index++
    in_order_traversal(bst.right, result, index)
}

func count_nodes(bst *BinarySearchTree) int {
    if bst == nil {
        return 0
    }

    return 1 + count_nodes(bst.left) + count_nodes(bst.right)
}
