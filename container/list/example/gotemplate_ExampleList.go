// Code generated by gotemplate. DO NOT EDIT.

package example

import (
	"unsafe"

	"foundation/offsetptr"
)

// template type List(Value,Allocator)

type NodeExampleList struct {
	next  offsetptr.Pointer `*Node`
	prev  offsetptr.Pointer `*Node`
	list  offsetptr.Pointer `*List`
	Value Item
}

const _SizeofNodeExampleList = unsafe.Sizeof(NodeExampleList{})

func NewNodeExampleList(value Item) *NodeExampleList {
	var node *NodeExampleList
	if defaultAlloc == nil {
		node = new(NodeExampleList)
	} else {
		node = (*NodeExampleList)(defaultAlloc.Allocate(_SizeofNodeExampleList))
	}

	node.next = *offsetptr.NewNil()
	node.prev = *offsetptr.NewNil()
	node.list = *offsetptr.NewNil()
	node.Value = value

	return node
}

func (n *NodeExampleList) Free() {
	if n != nil {
		defaultAlloc.DeAllocate(unsafe.Pointer(n))
	}
}

func (n *NodeExampleList) Next() *NodeExampleList {
	if p := (*NodeExampleList)(n.next.Get()); p != &(*ExampleList)(n.list.Get()).root {
		return p
	}
	return nil
}

type ExampleList struct {
	root NodeExampleList
	len  int
}

const _SizeofExampleList = unsafe.Sizeof(ExampleList{})

func NewExampleList() *ExampleList {
	if defaultAlloc == nil {
		return new(ExampleList).Init()
	} else {
		return (*ExampleList)(defaultAlloc.Allocate(_SizeofExampleList)).Init()
	}
}

func (l *ExampleList) Free() {
	if l != nil {
		defaultAlloc.DeAllocate(unsafe.Pointer(l))
	}
}

func (l *ExampleList) Init() *ExampleList {
	l.root.next.Set(unsafe.Pointer(&l.root))
	l.root.prev.Set(unsafe.Pointer(&l.root))
	l.len = 0
	return l
}

func (l *ExampleList) lazyInit() {
	if l.len == 0 {
		l.Init()
	}
}

func (l *ExampleList) Front() *NodeExampleList {
	if l.len == 0 {
		return nil
	}

	return (*NodeExampleList)(l.root.next.Get())
}

func (l *ExampleList) Back() *NodeExampleList {
	if l.len == 0 {
		return nil
	}

	return (*NodeExampleList)(l.root.prev.Get())
}

func (l *ExampleList) insert(e, at *NodeExampleList) *NodeExampleList {
	n := offsetptr.NewPointer(at.next.Get())

	at.next.Set(unsafe.Pointer(e))
	e.prev.Set(unsafe.Pointer(at))
	e.next.Set(n.Get())
	(*NodeExampleList)(n.Get()).prev.Set(unsafe.Pointer(e))
	e.list.Set(unsafe.Pointer(l))

	l.len++
	return e
}

func (l *ExampleList) PushFront(value Item) *NodeExampleList {
	l.lazyInit()
	return l.insert(NewNodeExampleList(value), &l.root)
}

func (l *ExampleList) PushBack(value Item) *NodeExampleList {
	l.lazyInit()
	return l.insert(NewNodeExampleList(value), (*NodeExampleList)(l.root.prev.Get()))
}

func (l *ExampleList) Values() []Item {
	if l.len == 0 {
		return nil
	}

	values := make([]Item, 0, l.len)

	for n := l.Front(); n != nil; n = n.Next() {
		values = append(values, n.Value)
	}

	return values
}

func (l *ExampleList) remove(n *NodeExampleList) {
	(*NodeExampleList)(n.prev.Get()).next.Forward(&n.next)
	(*NodeExampleList)(n.next.Get()).prev.Forward(&n.prev)
	(*NodeExampleList)(n.next.Get()).Free()
	(*NodeExampleList)(n.prev.Get()).Free()
	n.next.Set(nil)
	n.prev.Set(nil)
	n.list.Set(nil)
	l.len--
}

func (l *ExampleList) Remove(n *NodeExampleList) {
	if (*ExampleList)(n.list.Get()) == l {
		l.remove(n)
	}
}
