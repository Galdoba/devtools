package doublelinkedlist

// Node - is element struct for LinkedList and DoubleLinkedList
type Node[T comparable] struct {
	property T
	nextNode *Node[T]
	prevNode *Node[T]
}

// func NewNode[T comparable](property T) *Node[T] {
// 	n := Node[T]{}
// 	n.property = property
// 	return &n
// }

// func newNode[T comparable]() *Node[T] {
// 	n := Node[T]{}
// 	return &n
// }

// List - has a headNode pointer as it's property.
// By traversing to nextNode from headNode, you can iterate
// through the linked List
type List[T comparable] struct {
	headNode *Node[T]
}

func NewList[T comparable]() List[T] {
	lst := List[T]{}
	return lst
}

// AddToHead - adds node to start of the linked list.
// A new node is instatiated and its property is set to the property
// parameter that is passed. The nextNode points to the current headNode
// of lst, and headNode is set to the pointer of the new node
// that is created.
func (lst *List[T]) AddToHead(property T) {
	var node = &Node[T]{}
	node.property = property
	node.nextNode = nil
	if lst.headNode != nil {
		node.nextNode = lst.headNode
		lst.headNode.prevNode = node
	}
	lst.headNode = node
}

// IterateList - iterates form headNode property and do some action
// with property of current node.
// The iteration happens with the head node moves to nextNode of
// the headNode property untill the current node is no longer 'nil'
func (lst *List[T]) IterateList(action func(T) error) {
	var node *Node[T]
	for node = lst.headNode; node != nil; node = node.nextNode {
		action(node.property)
	}
}

// LastNode - return the node at the end of the list.
// List is traversed to check if nextNode is nil from nextNode of headNode
func (lst *List[T]) LastNode() *Node[T] {
	var node *Node[T]
	var lastNode *Node[T]
	for node = lst.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd - adds node at the end of the list.
func (lst *List[T]) AddToEnd(property T) {
	var node = &Node[T]{}
	node.property = property
	node.nextNode = nil
	var lastNode *Node[T]
	lastNode = lst.LastNode()
	if lastNode != nil {
		node.prevNode = lastNode
		lastNode.nextNode = node
	}
}

// NodeWithValue - returns first node with matched properties
func (lst *List[T]) NodeWithValue(property T) *Node[T] {
	var node *Node[T]
	var nodeWith *Node[T]
	for node = lst.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddAfter - adds the node after a specific node. A node with the nodeProperty value
// is retrived using NodeWithValue() method. A node with property is created and
// added after the 'nodeWith'.
func (lst *List[T]) AddAfter(nodeProperty, property T) {
	var node = &Node[T]{}
	node.property = property
	node.nextNode = nil
	var nodeWith *Node[T]
	nodeWith = lst.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		// node.prevNode = nodeWith
		nodeWith.nextNode = node
	}
}

func (n *Node[T]) hasProperty(property T) bool {
	return n.property == property
}

// NodeBetweenValues - return Node after traverse the list to find Node with
// previousNode and NextNode matching first and second properties
func (lst *List[T]) NodeBetweenValues(first, second T) *Node[T] {
	var node *Node[T]
	var nodeWith *Node[T]
	for node = lst.headNode; node != nil; node = node.nextNode {
		if node.prevNode != nil && node.nextNode != nil {
			if node.prevNode.hasProperty(first) && node.nextNode.hasProperty(second) {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}
