	    //  Data Structure
	//    Queue

		// QUEUE     (FIFO) First in first out   ==> queua and dequeue operations
		// ******      ******      ******      ******   
		// *  3 *      *  2 *      *  1 *      *  0 *
		// ******      ******      ******      ******

					//  STACK                                           
					// ******
					// *  3 *
					// ******

					// ******
					// *  2 *
					// ******

					// ******
					// *  1 *
					// ******

					// ******
					// *  0 *
					// ******


package queue

import "fmt"

// Queue
type Queue struct {
	items []int
}
// Queue enqueue
func(q *Queue) Enqueue(i int){
	// append value in param to the Queue slice
	q.items = append(q.items, i)
}

// Dequeuing  ==> value is removed from the front
func(q *Queue) Dequeue() int {
	// first value added
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}




func queue(){
	myQueue := Queue{}
	fmt.Println(myQueue)

	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)
	fmt.Println(myQueue)

	myQueue.Dequeue()
	fmt.Println(myQueue)

}