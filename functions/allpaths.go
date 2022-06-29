package functions

var visitedVertices = make(map[string]bool)
var path = []string{}

// AllPaths - Finds all possible paths from the start vertice to the end vertice and stores them in a global variable KnownPaths
func (graph *Graph) AllPaths(vertice string) {
	if visitedVertices[vertice] {
		return
	}
	visitedVertices[vertice] = true
	path = append(path, vertice)
	if vertice != graph.EndVertice {
		for _, adjacentVertice := range graph.VerticeEdges[vertice] {
			graph.AllPaths(adjacentVertice)
		}
	} else {
		var additionalPath = make([]string, len(path))
		for index, pathVertice := range path {
			additionalPath[index] = pathVertice
		}
		KnownPaths = append(KnownPaths, additionalPath)
	}
	path = path[:len(path)-1]
	visitedVertices[vertice] = false
}
