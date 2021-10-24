	    //  Data Structure
	// heap algorithm in Go (HEAPSORT) ==> priority keys [heaps help for priority(highest priority)]
	// every parent must have a higher value than the children
	// nodes in heap are arranged as array

	// ==> MAX HEAP

package heap

import "fmt"

// MaxHeap ==> struct has a slice that holds the array
type MaxHeap struct{
	array []int
}

// Insert ==> Add an aelement to the heap
func (h *MaxHeap) Insert(key int){
	// takes the param and add to the array
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract ==> returns the largest key, removes it from the heap
func (h *MaxHeap) Extract() int {
	// get the root index
	extract := h.array[0]
	// last index
	lastIndex := len(h.array) - 1

	// check for emty array
	if len(h.array) == 0 {
		fmt.Println("Extraction from empty array not possible")
		// closes the method
		return -1
	}
	//replace the first index with the last
	h.array[0] = h.array[lastIndex]
	// slice to prevent repitition
	h.array = h.array[:lastIndex]

	return extract
}


func (h *MaxHeap) maxHeapifyUp(index int){
	// when parent index is smaller than index of current array
	for h.array[parent(index)] < h.array[index]{
		// swap parent with current index
		h.swap(parent(index), index)
		// update index
		index = parent(index)
	}
}

// Heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {

	lastIndex := len(h.array) - 1
	l,r := left(index), right(index)
	childToCompare := 0

	// as long as index still has a child
	for l <= lastIndex{
		// when left index is the only child
		if l ==lastIndex{
			childToCompare = l
		} else if h.array[l] > h.array[r] {     // when left child is larger
			childToCompare = l
		}else {                                 // when right child is larger
			childToCompare = r
		}

		// compare array value of current index to larger child and swap if smaller child to compare
		if h.array[index] < h.array[childToCompare]{
			h.swap(index, childToCompare)
			index = childToCompare
			l,r = left(index), right(index)
		}else{
			return 
		}

		
	}
	 
}

// get parent index
func parent(i int) int {
	return (i-1) / 2
}

// left and right child formular
//  parent index x 2 = 1 = left child index
//  parent index x 2 = 2 = right child index

// get child left child index
func left(i int) int {
	return 2 * i + 1
}

// get right child index
func right(i int) int {
	return 2 * i + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1,i2 int){
	//  ==> GO supports swap =>  1,2 = 2,1
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}


// Heapify
func heap(){
	
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 35, 40, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap{
		m.Insert(v)
		fmt.Println(m)
	}
	//  Use breakpoint and debugger tool to check each swap process 
	// Run the loop 7times
	for i := 0; i < 7; i++ {
		m.Extract()
		fmt.Println(m)
	}
}