package main

func getNameCounts(names []string) map[rune]map[string]int {
	value := make(map[rune]map[string]int)
	nameCount := make(map[string]int)
	for _, name := range names {
		runes := []rune(name)
		nameCount[name] += 1
		value[runes[0]] = nameCount
		
		
	}
	return value
}