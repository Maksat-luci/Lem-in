package functions

import "fmt"

// Ants - sends the ants
func (g *Graph) Ants(group [][]string, n int) {
	TickNumber := n / len(group)
	if n%len(group) != 0 {
		TickNumber++
	}
	var ants = make([]Ant, n+1)
	// ignore zero because ants start from 1
	ants[0].Finished = true
	number := 0
	for j := 0; j < TickNumber; j++ {
		for _, path := range group {
			number++
			ants[number].Path = path
			ants[number].CurrentVertice = 0
			ants[number].Finished = false
			if number == n {
				break
			}
		}
	}
	EndOfAnts := false
	var occupiedVertices = make(map[string]bool)
	for !EndOfAnts {
		for number, ant := range ants {
			if ant.Finished {
				continue
			}
			vertice := ant.Path[ant.CurrentVertice]
			if occupiedVertices[vertice] {
				fmt.Println()
				break
			}
			fmt.Print("L", number, "-", vertice, " ")
			if number == n {
				fmt.Println()
				if vertice == g.EndVertice {
					ants[number].Finished = true
					for i, ant1 := range ants {
						if ant1.Finished != true {
							fmt.Print("L", i, "-", g.EndVertice, " ")
							fmt.Println()
						}
					}
					EndOfAnts = true
				}
			}
			ants[number].CurrentVertice++
			occupiedVertices[ants[number].PreviousVertice] = false
			if vertice != g.EndVertice {
				occupiedVertices[vertice] = true
				ants[number].PreviousVertice = vertice
			}
			if vertice == g.EndVertice {
				ants[number].Finished = true
			}
		}
	}
}
