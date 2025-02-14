package main

import "fmt"

type BinaryTree struct {
	key        int
	leftChild  *BinaryTree
	rightChild *BinaryTree
}

func NewBinaryTree(value int) *BinaryTree {
	return &BinaryTree{key: value}
}

func (t *BinaryTree) InsertLeft(value int) {
	if t.leftChild == nil {
		t.leftChild = NewBinaryTree(value)
	} else {
		newTree := NewBinaryTree(value)
		newTree.leftChild = t.leftChild
		t.leftChild = newTree
	}
}

func (t *BinaryTree) InsertRight(value int) {
	if t.rightChild == nil {
		t.rightChild = NewBinaryTree(value)
	} else {
		newTree := NewBinaryTree(value)
		newTree.rightChild = t.rightChild
		t.rightChild = newTree
	}
}

func (t *BinaryTree) BreadthFirstSearch(n int) bool {
	queue := []*BinaryTree{t}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.key == n {
			return true
		}
		if node.leftChild != nil {
			queue = append(queue, node.leftChild)
		}
		if node.rightChild != nil {
			queue = append(queue, node.rightChild)
		}
	}
	return false
}

func (t *BinaryTree) Preorder() {
	if t != nil {
		fmt.Println(t.key)
		if t.leftChild != nil {
			t.leftChild.Preorder()
		}
		if t.rightChild != nil {
			t.rightChild.Preorder()
		}
	}
}

func (t *BinaryTree) Postorder() {
	if t != nil {
		if t.leftChild != nil {
			t.leftChild.Postorder()
		}
		if t.rightChild != nil {
			t.rightChild.Postorder()
		}
		fmt.Println(t.key)
	}
}

func (t *BinaryTree) Inorder() {
	if t != nil {
		if t.leftChild != nil {
			t.leftChild.Inorder()
		}
		fmt.Println(t.key)
		if t.rightChild != nil {
			t.rightChild.Inorder()
		}
	}
}

func (t *BinaryTree) Invert() {
	queue := []*BinaryTree{t}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		node.leftChild, node.rightChild = node.rightChild, node.leftChild
		if node.leftChild != nil {
			queue = append(queue, node.leftChild)
		}
		if node.rightChild != nil {
			queue = append(queue, node.rightChild)
		}
	}
}

func main() {
	tree := NewBinaryTree(1)
	tree.InsertLeft(2)
	tree.InsertRight(3)
	tree.InsertLeft(4)
	tree.leftChild.InsertRight(6)
	tree.InsertRight(5)

	fmt.Println("Pre-order:")
	tree.Preorder()

	fmt.Println("Post-order:")
	tree.Postorder()

	fmt.Println("In-order:")
	tree.Inorder()

	fmt.Println("Before inversion:")
	tree.Inorder()
	tree.Invert()
	fmt.Println("After inversion:")
	tree.Inorder()

	if tree.BreadthFirstSearch(5) {
		fmt.Println("Value 5 found in the tree.")
	} else {
		fmt.Println("Value 5 not found in tree.")
	}
}
