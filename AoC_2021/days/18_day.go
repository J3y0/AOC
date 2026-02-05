package days

import (
	"fmt"
	"main/utils"
	"strings"
)

type snailNode struct {
	value  int
	parent *snailNode
	left   *snailNode
	right  *snailNode
}

// parse a tree from a single line string
func nodesFromString(s string) (*snailNode, error) {
	if !strings.HasPrefix(s, "[") || !strings.HasSuffix(s, "]") {
		return nil, fmt.Errorf("invalid pair string")
	}
	nodes := make([]*snailNode, 0)
	countOpenedBraces := 0
	for _, r := range s {
		switch r {
		case '[':
			countOpenedBraces++
			continue
		case ',':
			continue
		case ']':
			countOpenedBraces--
			if len(nodes) < 2 {
				return nil, fmt.Errorf("invalid pair string")
			}
			n1 := nodes[len(nodes)-2]
			n2 := nodes[len(nodes)-1]
			nodes = nodes[:len(nodes)-2]
			parent := &snailNode{
				value:  -1,
				parent: nil,
				left:   n1,
				right:  n2,
			}
			n1.parent = parent
			n2.parent = parent
			nodes = append(nodes, parent)
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			nodes = append(nodes, newLiteralNode(int(r-'0'), nil))
		default:
			return nil, fmt.Errorf("invalid pair string")
		}
	}
	// countOpenedBraces only accounts for verification
	if countOpenedBraces != 0 {
		return nil, fmt.Errorf("invalid pair string")
	}
	if len(nodes) != 1 {
		return nil, fmt.Errorf("invalid pair string")
	}
	return nodes[0], nil
}

// helper function to create a node holding a value
func newLiteralNode(value int, parent *snailNode) *snailNode {
	return &snailNode{
		value:  value,
		parent: parent,
		left:   nil,
		right:  nil,
	}
}

// n and other have their parents modified.
// It returns the new root of the merged tree where n is the left branch and other the right branch
func (n *snailNode) join(other *snailNode) *snailNode {
	newRoot := &snailNode{
		value:  -1,
		parent: nil,
		left:   n,
		right:  other,
	}
	n.parent = newRoot
	other.parent = newRoot
	return newRoot
}

// represent the tree as a single line string with brackets
func (n *snailNode) String() string {
	if n == nil {
		return "nil"
	}
	if n.isLeaf() {
		return fmt.Sprintf("%d", n.value)
	}
	var sb strings.Builder
	sb.WriteString("[")
	sb.WriteString(n.left.String())
	sb.WriteString(",")
	sb.WriteString(n.right.String())
	sb.WriteString("]")
	return sb.String()
}

// clone a node and its sub-tree
func (n *snailNode) Clone() *snailNode {
	if n == nil {
		return nil
	}
	copy := &snailNode{
		value: n.value,
	}
	if n.left != nil {
		copy.left = n.left.Clone()
		copy.left.parent = copy
	}
	if n.right != nil {
		copy.right = n.right.Clone()
		copy.right.parent = copy
	}
	return copy
}

// determine if current node n is a leaf
func (n *snailNode) isLeaf() bool {
	return n.left == nil && n.right == nil
}

// determine if current node n is the root of the tree
func (n *snailNode) isRoot() bool {
	return n.parent == nil
}

// true if n is the left child
func (n *snailNode) isLeftChild() bool {
	return !n.isRoot() && n.parent.left == n
}

// true if n is the right child
func (n *snailNode) isRightChild() bool {
	return !n.isRoot() && n.parent.right == n
}

// apply explode when possible, otherwise try to split.
// Keep going if there is still a reduction operation to perform
func (n *snailNode) reduce() {
	explode, split := true, true
	for explode || split {
		explode = n.explode(0)
		if !explode {
			split = n.split()
		}
	}
}

// return true if a split occured
func (n *snailNode) split() bool {
	if n.isLeaf() && n.value >= 10 {
		l := n.value / 2
		r := (n.value + 1) / 2
		n.left = newLiteralNode(l, n)
		n.right = newLiteralNode(r, n)
		n.value = -1
		return true
	}
	if n.left != nil && n.right != nil {
		return n.left.split() || n.right.split()
	}
	return false
}

// traverse the tree to apply explosion if
func (n *snailNode) explode(depth int) bool {
	if depth >= 4 && n.left != nil && n.left.isLeaf() && n.right != nil && n.right.isLeaf() {
		n.explodeNode()
		return true
	}

	if n.left != nil && n.right != nil {
		return n.left.explode(depth+1) || n.right.explode(depth+1)
	}
	return false
}

// explodeNode applies directly on the node and performs the modification for an explosion (not on the root node)
func (n *snailNode) explodeNode() {
	leftVal := n.left.value
	rightVal := n.right.value
	leftOk, rightOk := false, false
	cur := n
	for !cur.isRoot() {
		if !rightOk && cur.isLeftChild() {
			// explore right
			child := cur.parent.right
			for child != nil && !child.isLeaf() {
				child = child.left
			}

			if child != nil {
				child.value += rightVal
				rightOk = true
			}
		} else if !leftOk && cur.isRightChild() {
			// explore left
			child := cur.parent.left
			for child != nil && !child.isLeaf() {
				child = child.right
			}

			if child != nil {
				child.value += leftVal
				leftOk = true
			}
		}

		cur = cur.parent
	}
	// update node
	n.left = nil
	n.right = nil
	n.value = 0
}

// compute the magnitude of a tree
func (n *snailNode) magnitude() int {
	if n.isLeaf() {
		return n.value
	} else {
		return 3*n.left.magnitude() + 2*n.right.magnitude()
	}
}

type Day18 struct {
	snailNumbers []*snailNode
}

func (d *Day18) Parse(input string) error {
	lines := utils.ParseLines(input)
	d.snailNumbers = make([]*snailNode, len(lines))
	for i, l := range lines {
		n, err := nodesFromString(l)
		if err != nil {
			return err
		}
		d.snailNumbers[i] = n
	}
	return nil
}

func (d *Day18) Part1() (int, error) {
	res := d.snailNumbers[0]
	for i := 1; i < len(d.snailNumbers); i++ {
		res = add(res, d.snailNumbers[i])
	}
	return res.magnitude(), nil
}

func (d *Day18) Part2() (int, error) {
	maxi := 0
	for i := range d.snailNumbers {
		for j := i + 1; j < len(d.snailNumbers); j++ {
			maxi = max(maxi, add(d.snailNumbers[i], d.snailNumbers[j]).magnitude())
			maxi = max(maxi, add(d.snailNumbers[j], d.snailNumbers[i]).magnitude())
		}
	}
	return maxi, nil
}

func add(a, b *snailNode) *snailNode {
	res := a.Clone().join(b.Clone())
	res.reduce()
	return res
}
