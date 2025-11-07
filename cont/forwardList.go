package cont /* Containers */

// TODO docs
type ForwardListNode[ValueType any] struct {
	Value ValueType
	Next  *ForwardListNode[ValueType]
}

// TODO docs
type ForwardList[ValueType any] struct {
	head *ForwardListNode[ValueType]
	len int
}

// TODO docs
func NewForwardList[V any]() ForwardList[V] {
	return ForwardList[V]{nil, 0}
}

// TODO docs
func (fl ForwardList[V]) IsEmpty() bool {
	return fl.head == nil
}

// TODO docs
func (fl ForwardList[V]) Len() int {
	return fl.len
}

// TODO docs
func (fl ForwardList[V]) Front() *ForwardListNode[V] {
	return fl.head
}

// TODO docs
func (fl ForwardList[V]) Push(value V) {
	newHead := ForwardListNode{value, fl.head}
	fl.head = newHead
	fl.len++
}

// TODO docs
func (fl ForwardList[V]) Pop() {
	if fl.head == nil {
		// TODO
		return
	}
	fl.head = fl.head.Next
	fl.len--
}

// TODO docs
func (fl ForwardList[V]) Walk[CallableType any](callback CallableType) {
	for node := fl.Front(); node != nil; node = node.Next {
		callback(node.Value)
	}
}
