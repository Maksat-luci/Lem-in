package functions

// IntersectionCheck - check if the current path intersects with the other to avoid jams
func IntersectionCheck(group [][]string, path []string) bool {
	for index, vertice := range path {
		for _, groupPath := range group {
			for _, groupVertice := range groupPath {
				if vertice == groupVertice && index != len(path)-1 {
					return true
				}
			}
		}
	}
	return false
}
