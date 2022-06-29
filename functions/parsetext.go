package functions

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// ParseText - parses the data from the text file
func ParseText(filename string) (int, Graph, error) {
	var graph Graph
	text, err := ioutil.ReadFile("tests/" + filename)
	ErrorCheck(err)
	Lines := strings.Split(string(text), "\n")
	antsNumber, err := strconv.Atoi(Lines[0])
	ErrorCheck(err)
	roomsAndEdges := Lines[1:]
	lineType := "vertice"
	for _, line := range roomsAndEdges {
		if line == "##start" {
			lineType = "start"
			continue
		} else if line == "##end" {
			lineType = "end"
			continue
		} else if len(line) == 0 || line[:1] == "#" {
			continue
		}
		splitLines := strings.Split(line, " ")
		// fmt.Println(splitLines)
		if len(splitLines) == 1 && len(splitLines[0]) > 1 {
			edges := strings.Split(splitLines[0], "-")
			graph.AddEdge(edges[0], edges[1])
			continue
		}
		if lineType == "start" {
			graph.StartVertice = splitLines[0]
		} else if lineType == "end" {
			graph.EndVertice = splitLines[0]
		}
		lineType = "vertice"
	}
	FileText = string(text)
	return antsNumber, graph, err
}
