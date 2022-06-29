package functions

// AddEdge - adds edges between the vertices of the graph
func (graph *Graph) AddEdge(firstVertice, secondVertice string) {
	if graph.VerticeEdges == nil {
		graph.VerticeEdges = make(map[string][]string)
	}
	graph.VerticeEdges[firstVertice] = append(graph.VerticeEdges[firstVertice], secondVertice)
	graph.VerticeEdges[secondVertice] = append(graph.VerticeEdges[secondVertice], firstVertice)
}
