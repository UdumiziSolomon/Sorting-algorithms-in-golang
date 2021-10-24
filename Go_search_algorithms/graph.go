	    //  Data Structure
		// Graphs  ==> Abstract data type   [ dense(many edges), sparse(little edges) ]
		// garaph type ==> directed, undirected, cyclic, weighted and unweighted
		//  Vertex ==> node
		// edges ==> line that connect parent to children, vise versa and other vertexes

		// ways to represent graph   ==> Adjacency list (looks like list) and Adjacency matrix(2D Array)


package graph

import "fmt"

// Graph rep adjacency list graph
type Graph struct {
	vertices [] *Vertex
}

// Vertex
type Vertex struct {
	key int
	adjacent [] *Vertex
}

// Add Vertex
func (g *Graph) AddVertex (k int){
	if contains(g.vertices,k){
		err := fmt.Errorf("Vertex %v not added because it already exists", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// Add Edge  ==> adding edge to the graph  ==> Directed graph
func (g *Graph) AddEdge (from, to int){
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	// check error
	if fromVertex == nil || toVertex == nil {
		// check if edge is in range of the graph
		err := fmt.Errorf("Invalid edge (%v --> %v)", from , to)
		fmt.Println(err.Error())
	}else if contains(fromVertex.adjacent, to){     //if vertex contains edge already
		err := fmt.Errorf("Existing edge (%v --> %v)", from , to)
		fmt.Println(err.Error())
	} else{
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
	
}

// getVertex
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k{
			return g.vertices[i]
		}
	}
	return nil
}


// contains  ==> check if vertex already exist in the graph
func contains(s []*Vertex, k int) bool {
	for _, v := range s{
		if k == v.key {
			return true
		}
	}
	return false
}


// print adjacent list for each vertex in the graph
func (g *Graph) Print(){
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
			for _, v := range v.adjacent {
				fmt.Printf(" %v ", v.key)
			}
	}
}

func graph(){
	myGraph := &Graph{}

	for i:=0; i<5; i++ {
		myGraph.AddVertex(i)
	}
	fmt.Println(myGraph)

	// adding a new vertex
	myGraph.AddVertex(5)

	myGraph.AddEdge(1,4)
	//  show error for existing edge
	myGraph.AddEdge(1,4)
	// show error for 6 ==> sice it exceeds the range of edge in the vertex
	myGraph.AddEdge(2,6)
	myGraph.Print()
}