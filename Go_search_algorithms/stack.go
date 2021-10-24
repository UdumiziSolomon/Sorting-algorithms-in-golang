	    //  Data Structure
	//    Stack     (LIFO)  Last in first out  ==> push and pop operations

					//  STACK                                           
					// ******
					// *  1 *
					// ******

					// ******
					// *  2 *
					// ******

					// ******
					// *  3 *
					// ******

					// ******
					// *  4 *
					// ******

														// QUEUE     (FIFO) First in first out   ==> queua and dequeue operations
										// ******      ******      ******      ******   
										// *  1 *      *  2 *      *  3 *      *  4 *
										// ******      ******      ******      ******
package stack

import "fmt"

// Stack
type Stack struct {
	items []int
}
// Stack Push
func(s *Stack) Push(i int){
	// append value in param to the Stack slice
	s.items = append(s.items, i)
}

// Stack Pop
func (s *Stack) Pop() int{
	// last node in ths splice
	l := len(s.items) - 1
	toRemove := s.items[l]
	// splice off the last node
	s.items = s.items[:l]
	return toRemove
}



func stack(){
	myStack := Stack{}
	fmt.Println(myStack)

	myStack.Push(3)
	myStack.Push(5)
	myStack.Push(7)
	fmt.Println(myStack)
	
	// pop ops
	myStack.Pop()
	fmt.Println(myStack)

}