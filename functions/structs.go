package functions

// Graph - the graph itself
type Graph struct {
	StartVertice, EndVertice string
	VerticeEdges             map[string][]string
}

// Ant - for every ant
type Ant struct {
	Path            []string
	CurrentVertice  int
	PreviousVertice string
	Finished        bool
}

// GroupOfPathGroups - to store path groups
var GroupOfPathGroups = [][][]string{}

// KnownPaths - to store all paths available
var KnownPaths = [][]string{}

// FileText - to store the text from the file for printing
var FileText string
