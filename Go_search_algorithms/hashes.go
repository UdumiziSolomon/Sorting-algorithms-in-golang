	    //  Data Structure
		// Hash  ==> Quick search for words in an array   ==> through a hash function
		// hash code ==> output

		// Hash table ==> table
		// ASCII code
		// collision ==> cons of using a hash function in storing values in the has table (2 value having the same hash code)
		// collision handling (open addressing) && seperate chaining(array + linked list)

		// seperate chaining ==> linked lis [keys, bucket]


package hashes

import "fmt"
const ArraySize = 7

// Hashtable structure
type HashTable struct {
	array [ArraySize] *bucket
}
// Bucket structure  ==> linked list 
type bucket struct {
	head *bucketNode
}

// bucketNode structure
type bucketNode struct{
	key string
	next *bucketNode	
}

// Hash Table
// Insert  ==> take a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}
// Search  ==> return true if key is in the hash table
func (h *HashTable) Search(key string) bool{
	index := hash(key)
	return h.array[index].search(key)	
}

// Delete  ==> take a key and delete it from hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)	
}


// Bucket
// Insert
func (b *bucket) insert(k string){
	if !b.search(k){
		newNode := &bucketNode{ key: k}
		newNode.next = b.head
		b.head = newNode
	}else{
		fmt.Println(k," already exists")
	}
	
}
// Search
func (b *bucket) search(k string) bool{
	currentNode := b.head
	for currentNode != nil{
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
// Delete
func (b *bucket) delete(k string) {

	if b.head.key == k{
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil{
		if previousNode.next.key == k{
			// delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// hash function
func hash(key string) int {
	sum := 0
	for _, v := range key{
		sum += int(v)
	}
	return sum % ArraySize
}

// Init ==> function that initialises the hash table [create a bucket in each slot]
func Init() *HashTable{
	result := &HashTable{}
	for i:= range result.array{
		result.array[i] = &bucket{}
	}
	return result
}

func hashes(){
	myHash := Init()
	list := []string{ "SOLO", "EJIRO", "IFFY", "GODDY", "BRAD", "JOHN", "JOSH",}

	for _, v := range list{
		myHash.Insert(v)
	}
	fmt.Println(myHash)

	myHash.Delete("BRAD")
	fmt.Println("BRAD", myHash.Search("BRAD"))
	fmt.Println("GODDY", myHash.Search("GODDY"))


	
}