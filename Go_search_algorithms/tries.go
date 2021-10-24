	    //  Data Structure
		// Tries ==> structure used to store words
		// each name res a word or part of a word
		// google search uses tries

		// cons  ==> takes a lot of space and time complexity


package tries

import "fmt"

// no of possible characters
const AlphabetSize = 26

// Node
type Node struct {
	children [AlphabetSize] *Node
	isEnd    bool
}

// Trie
type Trie struct {
	root *Node
}

// create new Trie
func InitTrie() *Trie {
	result := &Trie{ root: &Node{}}
	return result
}

// Insert
func (t *Trie) Insert(w string) {
	// get length of the param letter
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'     // 'a'   = 97  
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	} 

	currentNode.isEnd = true
}

// Search    ==> will return true if word is included 
func (t *Trie) Search(w string)  bool{
		// get length of the param letter
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'     // 'a'   = 97  
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	} 

	if currentNode.isEnd == true{
		return true
	}

	return false
}




func tries(){
	myTrie := InitTrie()
	toAdd := []string{ "repert", "repo", "rep", "repoter",}
	for _, v:= range toAdd{
		myTrie.Insert(v)
	}
	fmt.Println(myTrie.Search("rep"))

}