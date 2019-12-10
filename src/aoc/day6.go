package aoc

import "strings"

func _createOrbitMap(input []string) map[string][]string {
	orbitMap := make(map[string][]string)
	for _, orbitDetails := range input {
		parts := strings.Split(orbitDetails, ")")
		existingValues, ok := orbitMap[parts[1]]
		if ok {
			existingValues = append(existingValues, parts[0])
			orbitMap[parts[1]] = existingValues
		} else {
			orbitMap[parts[1]] = []string{parts[0]}
		}
	}
	return orbitMap
}

func _findTarget(orbitMap map[string][]string, orbit string, path map[string]int, distance int) (bool, int) {
	innerOrbits, ok := orbitMap[orbit]
	found := false
	if ok && len(innerOrbits) > 0 {
		for _, inner := range innerOrbits {
			pathDistance, exists := path[inner]
			if exists {
				return exists, pathDistance + distance
			} else {
				path[inner] = distance + 1
				found, distance = _findTarget(orbitMap, inner, path, path[inner])
			}
		}
	}
	return found, distance
}

func solveDay6A(input []string) int {
	orbitMap := _createOrbitMap(input)
	count := 0
	for inner, _ := range orbitMap {
		_, tmp := _findTarget(orbitMap, inner, make(map[string]int), 0)
		count += tmp
	}
	return count
}

func solveDay6B(input []string) int {
	/* Instead of the usual BFS which would require a bi-directional graph we just use the one-directional graph and
	   do DFS twice. */
	orbitMap := _createOrbitMap(input)
	path := make(map[string]int)
	found, shortestPath := _findTarget(orbitMap, "YOU", path, 0)
	if !found {
		found, shortestPath = _findTarget(orbitMap, "SAN", path, 0)
	}
	if !found {
		panic("I am lost! Where is Santa?! :'(")
	}
	return shortestPath - 1
}
