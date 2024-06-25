package server

import (
	"strings"
)

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if strings.Contains(s, str) {
			return true
		}
	}
	return false
}

func removeDuplicates(id []int) []int {
	encountered := map[int]bool{}
	var result []int

	for _, v := range id {
		if encountered[v] == false {
			encountered[v] = true
			result = append(result, v)
		}
	}
	return result
}

func findDuplicates(input []int) []int {
	var duplicates []int
	seen := make(map[int]bool)

	for _, value := range input {
		if seen[value] {
			duplicates = append(duplicates, value)
		} else {
			seen[value] = true
		}
	}

	return duplicates
}

func mapsEqual(map1, map2 map[string][]string) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, valA := range map1 {
		valB, ok := map2[key]
		if !ok || len(valA) != len(valB) {
			return false
		}

		for i, v := range valA {
			if v != valB[i] {
				return false
			}
		}
	}

	return true
}
