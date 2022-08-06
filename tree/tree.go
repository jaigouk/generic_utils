package tree

import (
	"golang.org/x/exp/constraints"
	"fmt"
	"strings"
)

// Tree using generic nodes
type Tree[Value constraints.Ordered, Data any] struct {
	Root *Node[Value, Data]
}

// Insert a generic node to a tree
func (t *Tree[Value, Data]) Insert(value Value, data Data) {
	t.Root = t.Root.Insert(value, data)
	if t.Root.Bal() < -1 || t.Root.Bal() > 1 {
		t.rebalance()
	}
}

func (t *Tree[Value, Data]) rebalance() {
	if t == nil || t.Root == nil {
		return
	}
	t.Root = t.Root.rebalance()
}

// Find a node with a value
func (t *Tree[Value, Data]) Find(s Value) (Data, bool) {
	if t == nil || t.Root == nil {
		// Same situation as in method `Find` above.\
		// Here, we use `new` to create a zero value on the fly.\
		// `new` returns a pointer, and hence we need to add the dereferencing operator.
		return *new(Data), false
	}
	return t.Root.Find(s)
}

// Traverse tree from a node
func (t *Tree[Value, Data]) Traverse(n *Node[Value, Data], f func(*Node[Value, Data])) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

func (t *Tree[Value, Data]) PrettyPrint() {

	printNode := func(n *Node[Value, Data], depth int) {
		fmt.Printf("%s%v\n", strings.Repeat("  ", depth), n.Value)
	}
	var walk func(*Node[Value, Data], int)
	walk = func(n *Node[Value, Data], depth int) {
		if n == nil {
			return
		}
		walk(n.Right, depth+1)
		printNode(n, depth)
		walk(n.Left, depth+1)
	}

	walk(t.Root, 0)
}

func (t *Tree[Value, Data]) Dump() {
	t.Root.Dump(0, "")
}


// Node is a node in a tree. It has a value, a left child, a right child and a height.
type Node[Value constraints.Ordered, Data any] struct {
	Value  Value
	Data   Data
	Left   *Node[Value, Data]
	Right  *Node[Value, Data]
	height int
}

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Height of a tree
func (n *Node[Value, Data]) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}

// Bal between left and right
func (n *Node[Value, Data]) Bal() int {
	return n.Right.Height() - n.Left.Height()
}

// Insert a node to tree
func (n *Node[Value, Data]) Insert(value Value, data Data) *Node[Value, Data] {
	if n == nil {
		return &Node[Value, Data]{
			Value:  value,
			Data:   data,
			height: 1,
		}
	}
	if n.Value == value {
		n.Data = data
		return n
	}

	if value < n.Value {
		n.Left = n.Left.Insert(value, data)
	} else {
		n.Right = n.Right.Insert(value, data)
	}

	n.height = max(n.Left.Height(), n.Right.Height()) + 1

	return n.rebalance()
}

// rotaeLeft is a helper function for rebalance.
func (n *Node[Value, Data]) rotateLeft() *Node[Value, Data] {
	r := n.Right
	n.Right = r.Left
	r.Left = n
	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	r.height = max(r.Left.Height(), r.Right.Height()) + 1
	return r
}

// rotateRight is a helper function for rebalance.
func (n *Node[Value, Data]) rotateRight() *Node[Value, Data] {
	l := n.Left
	n.Left = l.Right
	l.Right = n
	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	l.height = max(l.Left.Height(), l.Right.Height()) + 1
	return l
}

// rotateRightLeft is a helper function for rebalance.
func (n *Node[Value, Data]) rotateRightLeft() *Node[Value, Data] {
	n.Right = n.Right.rotateRight()
	n = n.rotateLeft()
	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	return n
}

// rotateLeftRight is a helper function for rebalance.
func (n *Node[Value, Data]) rotateLeftRight() *Node[Value, Data] {
	n.Left = n.Left.rotateLeft()
	n = n.rotateRight()
	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	return n
}

// rebalance is a helper function for Insert.
func (n *Node[Value, Data]) rebalance() *Node[Value, Data] {
	switch {
	case n.Bal() < -1 && n.Left.Bal() == -1:
		return n.rotateRight()
	case n.Bal() > 1 && n.Right.Bal() == 1:
		return n.rotateLeft()
	case n.Bal() < -1 && n.Left.Bal() == 1:
		return n.rotateLeftRight()
	case n.Bal() > 1 && n.Right.Bal() == -1:
		return n.rotateRightLeft()
	}
	return n
}

// Find a node with a value
func (n *Node[Value, Data]) Find(s Value) (Data, bool) {

	if n == nil {
		var zero Data
		return zero, false
	}

	switch {
	case s == n.Value:
		return n.Data, true
	case s < n.Value:
		return n.Left.Find(s)
	default:
		return n.Right.Find(s)
	}
}

// Dump a tree
func (n *Node[Value, Data]) Dump(i int, lr string) {
	if n == nil {
		return
	}
	indent := ""
	if i > 0 {
		indent = strings.Repeat(" ", (i-1)*4) + "+" + lr + "--"
	}
	fmt.Printf("%s%v[%d,%d]\n", indent, n.Value, n.Bal(), n.Height())
	n.Left.Dump(i+1, "L")
	n.Right.Dump(i+1, "R")
}

