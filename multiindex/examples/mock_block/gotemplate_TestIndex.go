// Code generated by gotemplate. DO NOT EDIT.

package mock_block

import (
	"foundation/container"
	"foundation/multiindex"
)

// template type MultiIndex(SuperIndex,SuperNode,Value)

type TestIndex struct {
	super *ById
	count int
}

func NewTestIndex() *TestIndex {
	m := &TestIndex{}
	m.super = &ById{}
	m.super.init(m)
	return m
}

/*generic class*/

type TestIndexNode struct {
	super *ByIdNode
}

/*generic class*/

//method for MultiIndex
func (m *TestIndex) GetSuperIndex() interface{} { return m.super }
func (m *TestIndex) GetFinalIndex() interface{} { return nil }

func (m *TestIndex) GetIndex() interface{} {
	return nil
}

func (m *TestIndex) Size() int {
	return m.count
}

func (m *TestIndex) Clear() {
	m.super.clear()
	m.count = 0
}

func (m *TestIndex) Insert(v ValueType) bool {
	_, res := m.insert(v)
	return res
}

func (m *TestIndex) insert(v ValueType) (*TestIndexNode, bool) {
	fn := &TestIndexNode{}
	n, res := m.super.insert(v, fn)
	if res {
		fn.super = n
		m.count++
		return fn, true
	}
	return nil, false
}

func (m *TestIndex) Erase(iter multiindex.IteratorType) {
	m.super.erase_(iter)
}

func (m *TestIndex) erase(n *TestIndexNode) {
	m.super.erase(n.super)
	m.count--
}

func (m *TestIndex) Modify(iter multiindex.IteratorType, mod func(*ValueType)) bool {
	return m.super.modify_(iter, mod)
}

func (m *TestIndex) modify(mod func(*ValueType), n *TestIndexNode) (*TestIndexNode, bool) {
	defer func() {
		if e := recover(); e != nil {
			container.Logger.Error("#multi modify failed: %v", e)
			m.erase(n)
			m.count--
			panic(e)
		}
	}()
	mod(n.value())
	if sn, res := m.super.modify(n.super); !res {
		m.count--
		return nil, false
	} else {
		n.super = sn
		return n, true
	}
}

func (n *TestIndexNode) GetSuperNode() interface{} { return n.super }
func (n *TestIndexNode) GetFinalNode() interface{} { return nil }

func (n *TestIndexNode) value() *ValueType {
	return n.super.value()
}

/// IndexBase
type TestIndexBase struct {
	final *TestIndex
}

type TestIndexBaseNode struct {
	final *TestIndexNode
	pv    *ValueType
}

func (i *TestIndexBase) init(final *TestIndex) {
	i.final = final
}

func (i *TestIndexBase) clear() {}

func (i *TestIndexBase) GetSuperIndex() interface{} { return nil }

func (i *TestIndexBase) GetFinalIndex() interface{} { return i.final }

func (i *TestIndexBase) insert(v ValueType, fn *TestIndexNode) (*TestIndexBaseNode, bool) {
	return &TestIndexBaseNode{fn, &v}, true
}

func (i *TestIndexBase) erase(n *TestIndexBaseNode) {
	n.pv = nil
}

func (i *TestIndexBase) erase_(iter multiindex.IteratorType) {
	container.Logger.Warn("erase iterator doesn't match all index")
}

func (i *TestIndexBase) modify(n *TestIndexBaseNode) (*TestIndexBaseNode, bool) {
	return n, true
}

func (i *TestIndexBase) modify_(iter multiindex.IteratorType, mod func(*ValueType)) bool {
	container.Logger.Warn("modify iterator doesn't match all index")
	return false
}

func (n *TestIndexBaseNode) value() *ValueType {
	return n.pv
}
