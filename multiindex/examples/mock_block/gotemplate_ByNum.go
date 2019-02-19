// Code generated by gotemplate. DO NOT EDIT.

package mock_block

import (
	"fmt"

	"foundation/container"
	"foundation/multiindex"
)

// template type OrderedIndex(FinalIndex,FinalNode,SuperIndex,SuperNode,Value,Key,KeyFunc,Comparator,Multiply)

// OrderedIndex holds elements of the red-black tree
type ByNum struct {
	super *ByLibNum  // index on the OrderedIndex, IndexBase is the last super index
	final *TestIndex // index under the OrderedIndex, MultiIndex is the final index

	Root *ByNumNode
	size int
}

func (tree *ByNum) init(final *TestIndex) {
	tree.final = final
	tree.super = &ByLibNum{}
	tree.super.init(final)
}

func (tree *ByNum) clear() {
	tree.Clear()
	tree.super.clear()
}

/*generic class*/

/*generic class*/

// OrderedIndexNode is a single element within the tree
type ByNumNode struct {
	Key    ByNumComposite
	super  *ByLibNumNode
	final  *TestIndexNode
	color  colorByNum
	Left   *ByNumNode
	Right  *ByNumNode
	Parent *ByNumNode
}

/*generic class*/

/*generic class*/

func (node *ByNumNode) value() *ValueType {
	return node.super.value()
}

type colorByNum bool

const (
	blackByNum, redByNum colorByNum = true, false
)

func (tree *ByNum) Insert(v ValueType) (IteratorByNum, bool) {
	fn, res := tree.final.insert(v)
	if res {
		return tree.makeIterator(fn), true
	}
	return tree.End(), false
}

func (tree *ByNum) insert(v ValueType, fn *TestIndexNode) (*ByNumNode, bool) {
	key := ByNumKeyFunc(v)

	node, res := tree.put(key)
	if !res {
		container.Logger.Warn("#ordered index insert failed")
		return nil, false
	}
	sn, res := tree.super.insert(v, fn)
	if res {
		node.super = sn
		node.final = fn
		return node, true
	}
	tree.remove(node)
	return nil, false
}

func (tree *ByNum) Erase(iter IteratorByNum) (itr IteratorByNum) {
	itr = iter
	itr.Next()
	tree.final.erase(iter.node.final)
	return
}

func (tree *ByNum) Erases(first, last IteratorByNum) {
	for first != last {
		first = tree.Erase(first)
	}
}

func (tree *ByNum) erase(n *ByNumNode) {
	tree.remove(n)
	tree.super.erase(n.super)
	n.super = nil
	n.final = nil
}

func (tree *ByNum) erase_(iter multiindex.IteratorType) {
	if itr, ok := iter.(IteratorByNum); ok {
		tree.Erase(itr)
	} else {
		tree.super.erase_(iter)
	}
}

func (tree *ByNum) Modify(iter IteratorByNum, mod func(*ValueType)) bool {
	if _, b := tree.final.modify(mod, iter.node.final); b {
		return true
	}
	return false
}

func (tree *ByNum) modify(n *ByNumNode) (*ByNumNode, bool) {
	n.Key = ByNumKeyFunc(*n.value())

	if !tree.inPlace(n) {
		tree.remove(n)
		node, res := tree.put(n.Key)
		if !res {
			container.Logger.Warn("#ordered index modify failed")
			tree.super.erase(n.super)
			return nil, false
		}

		//n.Left = node.Left
		//if n.Left != nil {
		//	n.Left.Parent = n
		//}
		//n.Right = node.Right
		//if n.Right != nil {
		//	n.Right.Parent = n
		//}
		//n.Parent = node.Parent
		//if n.Parent != nil {
		//	if n.Parent.Left == node {
		//		n.Parent.Left = n
		//	} else {
		//		n.Parent.Right = n
		//	}
		//} else {
		//	tree.Root = n
		//}
		node.super = n.super
		node.final = n.final
		n = node
	}

	if sn, res := tree.super.modify(n.super); !res {
		tree.remove(n)
		return nil, false
	} else {
		n.super = sn
	}

	return n, true
}

func (tree *ByNum) modify_(iter multiindex.IteratorType, mod func(*ValueType)) bool {
	if itr, ok := iter.(IteratorByNum); ok {
		return tree.Modify(itr, mod)
	} else {
		return tree.super.modify_(iter, mod)
	}
}

// Get searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *ByNum) Find(key ByNumComposite) IteratorByNum {
	if true {
		lower := tree.LowerBound(key)
		if !lower.IsEnd() && ByNumCompare(key, lower.Key()) == 0 {
			return lower
		}
		return tree.End()
	} else {
		if node := tree.lookup(key); node != nil {
			return IteratorByNum{tree, node, betweenByNum}
		}
		return tree.End()
	}
}

// LowerBound returns an iterator pointing to the first element that is not less than the given key.
// Complexity: O(log N).
func (tree *ByNum) LowerBound(key ByNumComposite) IteratorByNum {
	result := tree.End()
	node := tree.Root

	if node == nil {
		return result
	}

	for {
		if ByNumCompare(key, node.Key) > 0 {
			if node.Right != nil {
				node = node.Right
			} else {
				return result
			}
		} else {
			result.node = node
			result.position = betweenByNum
			if node.Left != nil {
				node = node.Left
			} else {
				return result
			}
		}
	}
}

// UpperBound returns an iterator pointing to the first element that is greater than the given key.
// Complexity: O(log N).
func (tree *ByNum) UpperBound(key ByNumComposite) IteratorByNum {
	result := tree.End()
	node := tree.Root

	if node == nil {
		return result
	}

	for {
		if ByNumCompare(key, node.Key) >= 0 {
			if node.Right != nil {
				node = node.Right
			} else {
				return result
			}
		} else {
			result.node = node
			result.position = betweenByNum
			if node.Left != nil {
				node = node.Left
			} else {
				return result
			}
		}
	}
}

// Remove remove the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *ByNum) Remove(key ByNumComposite) {
	if true {
		for lower := tree.LowerBound(key); lower.position != endByNum; {
			if ByNumCompare(lower.Key(), key) == 0 {
				node := lower.node
				lower.Next()
				tree.remove(node)
			} else {
				break
			}
		}
	} else {
		node := tree.lookup(key)
		tree.remove(node)
	}
}

func (tree *ByNum) put(key ByNumComposite) (*ByNumNode, bool) {
	var insertedNode *ByNumNode
	if tree.Root == nil {
		// Assert key is of comparator's type for initial tree
		ByNumCompare(key, key)
		tree.Root = &ByNumNode{Key: key, color: redByNum}
		insertedNode = tree.Root
	} else {
		node := tree.Root
		loop := true
		if true {
			for loop {
				compare := ByNumCompare(key, node.Key)
				switch {
				case compare < 0:
					if node.Left == nil {
						node.Left = &ByNumNode{Key: key, color: redByNum}
						insertedNode = node.Left
						loop = false
					} else {
						node = node.Left
					}
				case compare >= 0:
					if node.Right == nil {
						node.Right = &ByNumNode{Key: key, color: redByNum}
						insertedNode = node.Right
						loop = false
					} else {
						node = node.Right
					}
				}
			}
		} else {
			for loop {
				compare := ByNumCompare(key, node.Key)
				switch {
				case compare == 0:
					node.Key = key
					return node, false
				case compare < 0:
					if node.Left == nil {
						node.Left = &ByNumNode{Key: key, color: redByNum}
						insertedNode = node.Left
						loop = false
					} else {
						node = node.Left
					}
				case compare > 0:
					if node.Right == nil {
						node.Right = &ByNumNode{Key: key, color: redByNum}
						insertedNode = node.Right
						loop = false
					} else {
						node = node.Right
					}
				}
			}
		}
		insertedNode.Parent = node
	}
	tree.insertCase1(insertedNode)
	tree.size++

	return insertedNode, true
}

func (tree *ByNum) swapNode(node *ByNumNode, pred *ByNumNode) {
	if node == pred {
		return
	}

	tmp := ByNumNode{color: pred.color, Left: pred.Left, Right: pred.Right, Parent: pred.Parent}

	pred.color = node.color
	node.color = tmp.color

	pred.Right = node.Right
	if pred.Right != nil {
		pred.Right.Parent = pred
	}
	node.Right = tmp.Right
	if node.Right != nil {
		node.Right.Parent = node
	}

	if pred.Parent == node {
		pred.Left = node
		node.Left = tmp.Left
		if node.Left != nil {
			node.Left.Parent = node
		}

		pred.Parent = node.Parent
		if pred.Parent != nil {
			if pred.Parent.Left == node {
				pred.Parent.Left = pred
			} else {
				pred.Parent.Right = pred
			}
		} else {
			tree.Root = pred
		}
		node.Parent = pred

	} else {
		pred.Left = node.Left
		if pred.Left != nil {
			pred.Left.Parent = pred
		}
		node.Left = tmp.Left
		if node.Left != nil {
			node.Left.Parent = node
		}

		pred.Parent = node.Parent
		if pred.Parent != nil {
			if pred.Parent.Left == node {
				pred.Parent.Left = pred
			} else {
				pred.Parent.Right = pred
			}
		} else {
			tree.Root = pred
		}

		node.Parent = tmp.Parent
		if node.Parent != nil {
			if node.Parent.Left == pred {
				node.Parent.Left = node
			} else {
				node.Parent.Right = node
			}
		} else {
			tree.Root = node
		}
	}
}

func (tree *ByNum) remove(node *ByNumNode) {
	var child *ByNumNode
	if node == nil {
		return
	}
	if node.Left != nil && node.Right != nil {
		pred := node.Left.maximumNode()
		tree.swapNode(node, pred)
	}
	if node.Left == nil || node.Right == nil {
		if node.Right == nil {
			child = node.Left
		} else {
			child = node.Right
		}
		if node.color == blackByNum {
			node.color = nodeColorByNum(child)
			tree.deleteCase1(node)
		}
		tree.replaceNode(node, child)
		if node.Parent == nil && child != nil {
			child.color = blackByNum
		}
	}
	tree.size--
}

func (tree *ByNum) lookup(key ByNumComposite) *ByNumNode {
	node := tree.Root
	for node != nil {
		compare := ByNumCompare(key, node.Key)
		switch {
		case compare == 0:
			return node
		case compare < 0:
			node = node.Left
		case compare > 0:
			node = node.Right
		}
	}
	return nil
}

// Empty returns true if tree does not contain any nodes
func (tree *ByNum) Empty() bool {
	return tree.size == 0
}

// Size returns number of nodes in the tree.
func (tree *ByNum) Size() int {
	return tree.size
}

// Keys returns all keys in-order
func (tree *ByNum) Keys() []ByNumComposite {
	keys := make([]ByNumComposite, tree.size)
	it := tree.Iterator()
	for i := 0; it.Next(); i++ {
		keys[i] = it.Key()
	}
	return keys
}

// Values returns all values in-order based on the key.
func (tree *ByNum) Values() []ValueType {
	values := make([]ValueType, tree.size)
	it := tree.Iterator()
	for i := 0; it.Next(); i++ {
		values[i] = it.Value()
	}
	return values
}

// Left returns the left-most (min) node or nil if tree is empty.
func (tree *ByNum) Left() *ByNumNode {
	var parent *ByNumNode
	current := tree.Root
	for current != nil {
		parent = current
		current = current.Left
	}
	return parent
}

// Right returns the right-most (max) node or nil if tree is empty.
func (tree *ByNum) Right() *ByNumNode {
	var parent *ByNumNode
	current := tree.Root
	for current != nil {
		parent = current
		current = current.Right
	}
	return parent
}

// Clear removes all nodes from the tree.
func (tree *ByNum) Clear() {
	tree.Root = nil
	tree.size = 0
}

// String returns a string representation of container
func (tree *ByNum) String() string {
	str := "OrderedIndex\n"
	if !tree.Empty() {
		outputByNum(tree.Root, "", true, &str)
	}
	return str
}

func (node *ByNumNode) String() string {
	if !node.color {
		return fmt.Sprintf("(%v,%v)", node.Key, "red")
	}
	return fmt.Sprintf("(%v)", node.Key)
}

func outputByNum(node *ByNumNode, prefix string, isTail bool, str *string) {
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		outputByNum(node.Right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		outputByNum(node.Left, newPrefix, true, str)
	}
}

func (node *ByNumNode) grandparent() *ByNumNode {
	if node != nil && node.Parent != nil {
		return node.Parent.Parent
	}
	return nil
}

func (node *ByNumNode) uncle() *ByNumNode {
	if node == nil || node.Parent == nil || node.Parent.Parent == nil {
		return nil
	}
	return node.Parent.sibling()
}

func (node *ByNumNode) sibling() *ByNumNode {
	if node == nil || node.Parent == nil {
		return nil
	}
	if node == node.Parent.Left {
		return node.Parent.Right
	}
	return node.Parent.Left
}

func (node *ByNumNode) isLeaf() bool {
	if node == nil {
		return true
	}
	if node.Right == nil && node.Left == nil {
		return true
	}
	return false
}

func (tree *ByNum) rotateLeft(node *ByNumNode) {
	right := node.Right
	tree.replaceNode(node, right)
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Left = node
	node.Parent = right
}

func (tree *ByNum) rotateRight(node *ByNumNode) {
	left := node.Left
	tree.replaceNode(node, left)
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Right = node
	node.Parent = left
}

func (tree *ByNum) replaceNode(old *ByNumNode, new *ByNumNode) {
	if old.Parent == nil {
		tree.Root = new
	} else {
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	if new != nil {
		new.Parent = old.Parent
	}
}

func (tree *ByNum) insertCase1(node *ByNumNode) {
	if node.Parent == nil {
		node.color = blackByNum
	} else {
		tree.insertCase2(node)
	}
}

func (tree *ByNum) insertCase2(node *ByNumNode) {
	if nodeColorByNum(node.Parent) == blackByNum {
		return
	}
	tree.insertCase3(node)
}

func (tree *ByNum) insertCase3(node *ByNumNode) {
	uncle := node.uncle()
	if nodeColorByNum(uncle) == redByNum {
		node.Parent.color = blackByNum
		uncle.color = blackByNum
		node.grandparent().color = redByNum
		tree.insertCase1(node.grandparent())
	} else {
		tree.insertCase4(node)
	}
}

func (tree *ByNum) insertCase4(node *ByNumNode) {
	grandparent := node.grandparent()
	if node == node.Parent.Right && node.Parent == grandparent.Left {
		tree.rotateLeft(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == grandparent.Right {
		tree.rotateRight(node.Parent)
		node = node.Right
	}
	tree.insertCase5(node)
}

func (tree *ByNum) insertCase5(node *ByNumNode) {
	node.Parent.color = blackByNum
	grandparent := node.grandparent()
	grandparent.color = redByNum
	if node == node.Parent.Left && node.Parent == grandparent.Left {
		tree.rotateRight(grandparent)
	} else if node == node.Parent.Right && node.Parent == grandparent.Right {
		tree.rotateLeft(grandparent)
	}
}

func (node *ByNumNode) maximumNode() *ByNumNode {
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (tree *ByNum) deleteCase1(node *ByNumNode) {
	if node.Parent == nil {
		return
	}
	tree.deleteCase2(node)
}

func (tree *ByNum) deleteCase2(node *ByNumNode) {
	sibling := node.sibling()
	if nodeColorByNum(sibling) == redByNum {
		node.Parent.color = redByNum
		sibling.color = blackByNum
		if node == node.Parent.Left {
			tree.rotateLeft(node.Parent)
		} else {
			tree.rotateRight(node.Parent)
		}
	}
	tree.deleteCase3(node)
}

func (tree *ByNum) deleteCase3(node *ByNumNode) {
	sibling := node.sibling()
	if nodeColorByNum(node.Parent) == blackByNum &&
		nodeColorByNum(sibling) == blackByNum &&
		nodeColorByNum(sibling.Left) == blackByNum &&
		nodeColorByNum(sibling.Right) == blackByNum {
		sibling.color = redByNum
		tree.deleteCase1(node.Parent)
	} else {
		tree.deleteCase4(node)
	}
}

func (tree *ByNum) deleteCase4(node *ByNumNode) {
	sibling := node.sibling()
	if nodeColorByNum(node.Parent) == redByNum &&
		nodeColorByNum(sibling) == blackByNum &&
		nodeColorByNum(sibling.Left) == blackByNum &&
		nodeColorByNum(sibling.Right) == blackByNum {
		sibling.color = redByNum
		node.Parent.color = blackByNum
	} else {
		tree.deleteCase5(node)
	}
}

func (tree *ByNum) deleteCase5(node *ByNumNode) {
	sibling := node.sibling()
	if node == node.Parent.Left &&
		nodeColorByNum(sibling) == blackByNum &&
		nodeColorByNum(sibling.Left) == redByNum &&
		nodeColorByNum(sibling.Right) == blackByNum {
		sibling.color = redByNum
		sibling.Left.color = blackByNum
		tree.rotateRight(sibling)
	} else if node == node.Parent.Right &&
		nodeColorByNum(sibling) == blackByNum &&
		nodeColorByNum(sibling.Right) == redByNum &&
		nodeColorByNum(sibling.Left) == blackByNum {
		sibling.color = redByNum
		sibling.Right.color = blackByNum
		tree.rotateLeft(sibling)
	}
	tree.deleteCase6(node)
}

func (tree *ByNum) deleteCase6(node *ByNumNode) {
	sibling := node.sibling()
	sibling.color = nodeColorByNum(node.Parent)
	node.Parent.color = blackByNum
	if node == node.Parent.Left && nodeColorByNum(sibling.Right) == redByNum {
		sibling.Right.color = blackByNum
		tree.rotateLeft(node.Parent)
	} else if nodeColorByNum(sibling.Left) == redByNum {
		sibling.Left.color = blackByNum
		tree.rotateRight(node.Parent)
	}
}

func nodeColorByNum(node *ByNumNode) colorByNum {
	if node == nil {
		return blackByNum
	}
	return node.color
}

//////////////iterator////////////////

func (tree *ByNum) makeIterator(fn *TestIndexNode) IteratorByNum {
	node := fn.GetSuperNode()
	for {
		if node == nil {
			panic("Wrong index node type!")

		} else if n, ok := node.(*ByNumNode); ok {
			return IteratorByNum{tree: tree, node: n, position: betweenByNum}
		} else {
			node = node.(multiindex.NodeType).GetSuperNode()
		}
	}
}

// Iterator holding the iterator's state
type IteratorByNum struct {
	tree     *ByNum
	node     *ByNumNode
	position positionByNum
}

type positionByNum byte

const (
	beginByNum, betweenByNum, endByNum positionByNum = 0, 1, 2
)

// Iterator returns a stateful iterator whose elements are key/value pairs.
func (tree *ByNum) Iterator() IteratorByNum {
	return IteratorByNum{tree: tree, node: nil, position: beginByNum}
}

func (tree *ByNum) Begin() IteratorByNum {
	itr := IteratorByNum{tree: tree, node: nil, position: beginByNum}
	itr.Next()
	return itr
}

func (tree *ByNum) End() IteratorByNum {
	return IteratorByNum{tree: tree, node: nil, position: endByNum}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's key and value can be retrieved by Key() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *IteratorByNum) Next() bool {
	if iterator.position == endByNum {
		goto end
	}
	if iterator.position == beginByNum {
		left := iterator.tree.Left()
		if left == nil {
			goto end
		}
		iterator.node = left
		goto between
	}
	if iterator.node.Right != nil {
		iterator.node = iterator.node.Right
		for iterator.node.Left != nil {
			iterator.node = iterator.node.Left
		}
		goto between
	}
	if iterator.node.Parent != nil {
		node := iterator.node
		for iterator.node.Parent != nil {
			iterator.node = iterator.node.Parent
			if node == iterator.node.Left {
				goto between
			}
			node = iterator.node
		}
	}

end:
	iterator.node = nil
	iterator.position = endByNum
	return false

between:
	iterator.position = betweenByNum
	return true
}

// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
// If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *IteratorByNum) Prev() bool {
	if iterator.position == beginByNum {
		goto begin
	}
	if iterator.position == endByNum {
		right := iterator.tree.Right()
		if right == nil {
			goto begin
		}
		iterator.node = right
		goto between
	}
	if iterator.node.Left != nil {
		iterator.node = iterator.node.Left
		for iterator.node.Right != nil {
			iterator.node = iterator.node.Right
		}
		goto between
	}
	if iterator.node.Parent != nil {
		node := iterator.node
		for iterator.node.Parent != nil {
			iterator.node = iterator.node.Parent
			if node == iterator.node.Right {
				goto between
			}
			node = iterator.node
			//if iterator.tree.Comparator(node.Key, iterator.node.Key) >= 0 {
			//	goto between
			//}
		}
	}

begin:
	iterator.node = nil
	iterator.position = beginByNum
	return false

between:
	iterator.position = betweenByNum
	return true
}

func (iterator IteratorByNum) HasNext() bool {
	return iterator.position != endByNum
}

func (iterator *IteratorByNum) HasPrev() bool {
	return iterator.position != beginByNum
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator IteratorByNum) Value() ValueType {
	return *iterator.node.value()
}

// Key returns the current element's key.
// Does not modify the state of the iterator.
func (iterator IteratorByNum) Key() ByNumComposite {
	return iterator.node.Key
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iterator *IteratorByNum) Begin() {
	iterator.node = nil
	iterator.position = beginByNum
}

func (iterator IteratorByNum) IsBegin() bool {
	return iterator.position == beginByNum
}

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (iterator *IteratorByNum) End() {
	iterator.node = nil
	iterator.position = endByNum
}

func (iterator IteratorByNum) IsEnd() bool {
	return iterator.position == endByNum
}

// Delete remove the node which pointed by the iterator
// Modifies the state of the iterator.
func (iterator *IteratorByNum) Delete() {
	node := iterator.node
	//iterator.Prev()
	iterator.tree.remove(node)
}

func (tree *ByNum) inPlace(n *ByNumNode) bool {
	prev := IteratorByNum{tree, n, betweenByNum}
	next := IteratorByNum{tree, n, betweenByNum}
	prev.Prev()
	next.Next()

	var (
		prevResult int
		nextResult int
	)

	if prev.IsBegin() {
		prevResult = 1
	} else {
		prevResult = ByNumCompare(n.Key, prev.Key())
	}

	if next.IsEnd() {
		nextResult = -1
	} else {
		nextResult = ByNumCompare(n.Key, next.Key())
	}

	return (true && prevResult >= 0 && nextResult <= 0) ||
		(!true && prevResult > 0 && nextResult < 0)
}
