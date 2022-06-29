package functions

// IntersectedPathsSeparation - separates paths which don't intersect with the others
func (g *Graph) IntersectedPathsSeparation() {
	for index, path := range KnownPaths {
		KnownPaths[index] = path[1:]
	}
	for index, path := range KnownPaths {
		var group = [][]string{}
		group = append(group, path)
		for index2 := index + 1; index2 < len(KnownPaths); index2++ {
			path2 := KnownPaths[index2]
			if !IntersectionCheck(group, path2) {
				group = append(group, path2)
			}
		}
		GroupOfPathGroups = append(GroupOfPathGroups, group)
	}
}
