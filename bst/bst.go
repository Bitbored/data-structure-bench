package bst

type node struct {
	key     int
	element *interface{}
	left    *node
	right   *node
}

type BST struct {
	root *node
}

func (tree *BST) Add(key int, element interface{}) {
	if tree.root == nil {
		tree.root = &node{key, &element, nil, nil}
	} else {
		tree.root.add(key, &element)
	}

}

func (n *node) add(key int, element *interface{}) {
	if key == n.key {
		n.element = element
	} else if key < n.key {
		if n.left == nil {
			n.left = &node{key, element, nil, nil}
		} else {
			n.left.add(key, element)
		}
	} else {
		if n.right == nil {
			n.right = &node{key, element, nil, nil}
		} else {
			n.right.add(key, element)
		}
	}
}

func (tree *BST) Get(key int) interface{} {
	n := tree.root.find(key)

	if n == nil {
		return nil
	}

	return *n.element
}

func (n *node) find(key int) *node {
	if n == nil || key == n.key {
		return n
	}

	if key < n.key {
		return n.left.find(key)
	}

	return n.right.find(key)
}
