	    //  Data Structure
		//   Linked list helps the implementation of adding data to a list of data in nodes in both the front node(single) and last node(doubly linked list)
		// each node contains the address of the next node // ==> single and doubly linked list
		// each node contains the address of the next node and previous node // ==> doubly linked list
	

package linkedList

import "fmt"

// Node
type node struct{
	data int   // node data
	next *node  //pointer to the node address
}

type linkedList struct {
	head *node  // first node of the list
	length int
}

// Insert node to the front  // ==> func prepend
func (l *linkedList) prepend (n *node){
	// current head of list
	second := l.head
	// data becomes the head
	l.head = n
	// old head turns to next node in the list
	l.head.next = second
	l.length++   // increase list length since a new node is added
}

// print out the value
func (l  linkedList) printListData(){
	// head node
	toPrint := l.head
	for l.length != 0 {
		// head is in the toPrint
		fmt.Printf("%d ",toPrint.data)
		// update the loop
		toPrint = toPrint.next
		l.length--
	}
	// line break
	fmt.Printf("\n")
}

// method to delete node
func (l *linkedList) deleteWithValue( value int){
	// Runtime error fixes
	if l.length == 0 {      // empty list fix
		return
	}

	if l.head.data == value {      // when we delete the head ==> we need to replace it with the next node 
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value{
		// if value is not in the list
		if previousToDelete.next.next == nil{
			return 
		}
		previousToDelete = previousToDelete.next
	}

	previousToDelete.next = previousToDelete.next.next
	l.length--
}


func linkedList(){
	myList := linkedList{}
	// add node to list 
	node1 := &node{data: 20}
	node2 := &node{data: 11}
	node3 := &node{data: 24}
	node4 := &node{data: 23}

	// prepend them to list
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)

	myList.printListData()
	myList.deleteWithValue(24)
	myList.printListData()

	// test for error  ==> NB: handled, should not throw error
	myList.deleteWithValue(50)     // 50 does not exist in list
	emptyList := linkedList{}      // deleting from an
	emptyList.deleteWithValue(10)
}