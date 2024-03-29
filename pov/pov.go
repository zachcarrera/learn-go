package pov

type Tree struct {
	// Add the needed fields here
	value    string
	children []*Tree
	parent   *Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	root := &Tree{value, children, nil}
	for _, child := range children {
		child.parent = root
	}
	return root
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.value
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	return tr.children
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// POV problem-specific functions

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	povRoot := tr.FindNode(from)
	seenNodes := make(map[string]bool)
	return povRoot.flipNode(seenNodes)
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	if tr == nil {
		return nil
	}

	if from == to {
		return []string{to}
	}

	fromNode := tr.FromPov(from)
	if fromNode == nil {
		return nil
	}
	return fromNode.FindPathFromRoot(to)
}

func (tr *Tree) FindNode(value string) *Tree {
	if tr == nil || tr.Value() == value {
		return tr
	}
	for _, child := range tr.Children() {
		node := child.FindNode(value)
		if node != nil {
			return node
		}
	}
	return nil
}

func (tr *Tree) FindPathFromRoot(to string) []string {
	toNode := tr.FindNode(to)
	if toNode == nil {
		return nil
	}

	path := make([]string, 0)
	for toNode != tr {
		path = append(path, toNode.value)
		toNode = toNode.parent
	}

	path = append(path, tr.value)
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func (tr *Tree) flipNode(visited map[string]bool) *Tree {
	if tr == nil || visited[tr.value] {
		return nil
	}
	visited[tr.value] = true
	children := make([]*Tree, 0)
	for _, child := range tr.Children() {
		flippedChild := child.flipNode(visited)
		if flippedChild != nil {
			children = append(children, flippedChild)
		}
	}

	if tr.parent != nil {
		flippedParent := tr.parent.flipNode(visited)
		if flippedParent != nil {
			children = append(children, flippedParent)
		}
	}
	return New(tr.value, children...)
}
