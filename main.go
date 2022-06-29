package main

import (
	"fmt"
	functions "lem-in/functions"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		filename := args[1]
		antsNumber, graph, err := functions.ParseText(filename)
		if err != nil {
			fmt.Println("ERROR: invalid data format")
			return
		}
		if antsNumber == 0 {
			fmt.Println("ERROR: invalid number of Ants!")
			return
		}
		graph.AllPaths(graph.StartVertice)
		for i := 0; i < len(functions.KnownPaths); i++ {
			for j := 0; j < len(functions.KnownPaths); j++ {
				if len(functions.KnownPaths[i]) < len(functions.KnownPaths[j]) {
					functions.KnownPaths[i], functions.KnownPaths[j] = functions.KnownPaths[j], functions.KnownPaths[i]
				}
			}
		}
		for _, path := range functions.KnownPaths {
			if len(path) == 2 {
				fmt.Println(functions.FileText)
				fmt.Println()
				for i := 1; i <= antsNumber; i++ {
					fmt.Print("L", i, "-", graph.EndVertice, " ")
				}
				fmt.Println("")
				os.Exit(0)
			}
		}
		graph.IntersectedPathsSeparation()
		graph.RemoveGroups()
		var minNumberOfTurns int
		var pathGroupTurns = make(map[int][][]string)
		for _, group := range functions.GroupOfPathGroups {
			groupLength := len(group)
			maxLen := len(group[len(group)-1])
			var remainingAnts = 0
			for _, path := range group {
				remainingAnts += maxLen - len(path)
			}
			antsInLine := antsNumber - remainingAnts
			if antsInLine == 0 {
				continue
			}
			numberOfTurns := maxLen + antsInLine/groupLength
			if antsInLine%groupLength != 0 {
				numberOfTurns++
			}
			pathGroupTurns[numberOfTurns] = group
			minNumberOfTurns = numberOfTurns
		}
		for turns := range pathGroupTurns {
			if turns < minNumberOfTurns {
				minNumberOfTurns = turns
			}
		}
		bestPathGroup := pathGroupTurns[minNumberOfTurns]
		if len(bestPathGroup) == 0 {
			fmt.Println("ERROR: Path or start/end vertice not found. Program aborted.")
			return
		}
		fmt.Println(functions.FileText)
		fmt.Println()
		graph.Ants(bestPathGroup, antsNumber)
	} else {
		fmt.Println("Invalid number of arguments")
	}
}
