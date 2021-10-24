	    //  Data Structure
		// BINARY TREES ==> Binary search trees  [all nodes have 2 children]  --> Balanced and unbalanced
		// Root ==> top node
		// parent ==> node with children
		// children ==> left[smaller], right[larger]
		// leaves ==> tree no children

		// Advantage ==> speed


package binaryTree

import "fmt"

// check search count
var count int

// Node  ==> rep components of a binary search tree
type Node struct{
	Key int
	// children
	Left *Node
	Right *Node
}

// Insert
	func (n *Node) Insert (k int){
		// add if key is not in the tree
		if n.Key < k{
			// move right
			if n.Right == nil {
				n.Right = &Node{Key: k}
			}else{
				// recursive call on the right node
				n.Right.Insert(k)
			}
		} else if n.Key > k {
			// move left
			if n.Left == nil {
				n.Left = &Node{Key: k}
			}else{
				// recursive call on the right node
				n.Left.Insert(k)
			}
		}
	}

// Search  ==> takes a value and return a boolean  of a node existence
func (n *Node) Search (k int) bool {
	// to check how many count iteration is done in searching for a particular value
	count++

	// when tree is empty and there's no match
	if n == nil {
		return false
	}

	if n.Key < k {
		// move right
		return n.Right.Search(k)
	} else if n.Key > k{
		// move left
		return n.Left.Search(k)
	}

	return true
}



func binaryTree(){
	tree := &Node{Key: 100}
	tree.Insert(50);    // ==> SMALLER  ==> LEFT SIDE
	fmt.Println(tree)

	toSearch := []int{ 101, 39, 29, 90, 11, 49, 5,}

	for _, s := range toSearch{
		tree.Insert(s)
	}

	
	// fmt.Println(tree.Search(200))   // ==> search for available ( 200 is not in tree, should return false) 
	// fmt.Println(count)
	fmt.Println(tree.Search(5))    // ==> returns true since 11 was inserted into the tree
	fmt.Println(count)             //  ==> return 6 iterations


}