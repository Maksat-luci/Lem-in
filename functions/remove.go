package functions

// PathValidityCheck - checks if the path group is valid
func PathValidityCheck(groups [][][]string, targetgroup [][]string) bool {
	for _, group := range groups {
		if len(targetgroup) == len(group) {
			var pathInvalidity = make([]bool, len(targetgroup))
			for i, targetPath := range targetgroup {
				groupPath := group[i]
				if len(targetPath) >= len(groupPath) {
					pathInvalidity = append(pathInvalidity, true)
				} else {
					pathInvalidity = append(pathInvalidity, false)
				}
			}
			counter := 0
			for _, value := range pathInvalidity {
				if value {
					counter++
				}
			}
			if counter == len(targetgroup) {
				return true
			}
		}
	}
	return false
}

// RemoveGroups - removes the unneccesary path groups
func (g *Graph) RemoveGroups() {
	var NewGroupOfPathGroups = [][][]string{}
	for index, targetgroup := range GroupOfPathGroups {
		if len(targetgroup) == 1 && index != 0 {
			continue
		}
		if !PathValidityCheck(NewGroupOfPathGroups, targetgroup) {
			NewGroupOfPathGroups = append(NewGroupOfPathGroups, targetgroup)
		}
	}
	GroupOfPathGroups = NewGroupOfPathGroups
}
