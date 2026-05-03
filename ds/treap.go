package ds

import "math/rand"

type Node struct {
	left, right              *Node
	key, cnt, size, priority int
}

type Treap struct {
	root *Node
}

func size(node *Node) int {
	if node == nil {
		return 0
	}
	return node.size
}

func updateSize(node *Node) {
	if node != nil {
		node.size = node.cnt + size(node.left) + size(node.right)
	}
}

func rightRotate(node *Node) *Node {
	newRoot := node.left
	node.left = newRoot.right
	newRoot.right = node

	updateSize(node)
	updateSize(newRoot)
	return newRoot
}

func leftRotate(node *Node) *Node {
	newRoot := node.right
	node.right = newRoot.left
	newRoot.left = node

	updateSize(node)
	updateSize(newRoot)
	return newRoot
}

func insert(node *Node, key int) *Node {
	if node == nil {
		return &Node{
			key:      key,
			priority: rand.Int(),
			cnt:      1,
			size:     1,
		}
	}

	if key < node.key {
		node.left = insert(node.left, key)
		if node.left.priority > node.priority {
			node = rightRotate(node)
		}
	} else if key > node.key {
		node.right = insert(node.right, key)
		if node.right.priority > node.priority {
			node = leftRotate(node)
		}
	} else {
		node.size++
		node.cnt++
	}

	updateSize(node)
	return node
}

func (t *Treap) Insert(key int) {
	t.root = insert(t.root, key)
}

func (t *Treap) Rank(key int) int {
	rank := 0
	node := t.root
	for node != nil {
		if key < node.key {
			node = node.left
		} else if key > node.key {
			rank += size(node.left) + node.cnt
			node = node.right
		} else {
			return rank + size(node.left) + 1
		}
	}
	return rank + 1
}
