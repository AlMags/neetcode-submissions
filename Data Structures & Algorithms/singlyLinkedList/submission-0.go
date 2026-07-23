type Node struct {
    value int
    next *Node
}

func CreateNewNode(value int, nextNode *Node) *Node {
    return &Node{
        value: value,
        next: nextNode,
    }
}

type LinkedList struct {
    head *Node
    tail *Node
}

func NewLinkedList() *LinkedList {
    head := CreateNewNode(-1, nil)
    return &LinkedList{
        head: head,
        tail: head,
    }
}

func (ll *LinkedList) Get(index int) int {
    currentNode := ll.head.next
    i := 0
    for currentNode != nil {
        if i == index {
            return currentNode.value
        }
        i++
        currentNode = currentNode.next
    }
    return -1
}

func (ll *LinkedList) InsertHead(val int) {
    newNode := CreateNewNode(val, ll.head.next)
    ll.head.next = newNode
    if newNode.next == nil {
        ll.tail = newNode
    }
    
}

func (ll *LinkedList) InsertTail(val int) {
    ll.tail.next = CreateNewNode(val, nil)
    ll.tail = ll.tail.next
}

func (ll *LinkedList) Remove(index int) bool {
    currentNode := ll.head
    i := 0

    for i < index && currentNode != nil {
        i++
        currentNode = currentNode.next
    }

    if currentNode != nil && currentNode.next != nil {
        if currentNode.next == ll.tail{
            ll.tail = currentNode
        }
        currentNode.next = currentNode.next.next
        return true
    }

    return false
}

func (ll *LinkedList) GetValues() []int {
    currentNode := ll.head.next
    var result []int
    for currentNode != nil {
        result = append(result, currentNode.value)
        currentNode = currentNode.next
    }
    return result
}
