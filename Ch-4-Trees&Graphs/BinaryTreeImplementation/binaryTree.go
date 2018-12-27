package BinaryTree

type BinaryTree struct {
	Root *Node
}

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}


func(b *BinaryTree) Insert(value int) {
	if b == nil {
		b.Root = &Node{Key: value}
	} else {
		b.Root.insert(value)
	}
}

func(n *Node) insert(value int) {
	// Left Insert
	if n.Key > value {
		if n.Left == nil {
			n.Left = &Node{Key:value}
		} else {
			n.Left.insert(value)
		}

	// Right Insert
	} else if n.Key < value {
		if n.Right == nil {
			n.Right = &Node{Key:value}
		} else {
			n.Right.insert(value)
		}
	}
}

func(b *BinaryTree) Search(value int) bool {
	if b.Root == nil {
		return false
	} else {
		return b.Root.search(value)
	}
	return true
}

func(n *Node) search(value int) bool {
	if n.Key > value {
		if n.Left == nil {
			return false
		} else {
			return n.Left.search(value)
		}
	} else if n.Key < value {
		if n.Right == nil {
			return false
		} else {
			return n.Right.search(value)
		}
	}
	return true
}

func(b *BinaryTree) Delete(value int) *Node {
	if b.Root == nil {
		return nil
	} else {
		return b.Root.delete(value)
	}
}

func(n *Node) delete(value int) *Node {
	if n.Key < value {
		n.Right = n.Right.delete(value)
	} else if n.Key > value {
		n.Left = n.Left.delete(value)
	} else {
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}
		min := n.Right.min()
		n.Key = min
		n.Right = n.Right.delete(min)
	}
	return n
}


func(n *Node) min() int {
	if n.Left != nil {
		return n.Key
	}
	return n.Left.min()
}

