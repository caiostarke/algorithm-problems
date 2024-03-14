package binary_tree

type Node struct {
	Value       int
	Left, Right *Node
}

type BynaryTree struct {
	Root *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

func (bt *BynaryTree) Insert(value int) {
	if bt.Root == nil {
		bt.Root = NewNode(value)
		return
	}

	insertNode(bt.Root, value)
}

// values := []int{4, 2, 7, 1, 3, 6, 9}
func insertNode(node *Node, value int) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = NewNode(value)
		} else {
			insertNode(node.Left, value)
		}
	} else {
		if node.Right == nil {
			node.Right = NewNode(value)
		} else {
			insertNode(node.Right, value)
		}
	}
}

func (BT *BynaryTree) FindMax() int {
	return findMax(BT.Root)
}

func findMax(node *Node) int {
	if node == nil {
		return -1
	}
	if node.Right == nil {
		return node.Value
	}
	return findMax(node.Right)
}
