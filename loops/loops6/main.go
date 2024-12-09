package main

func countConnections(groupSize int) int {
	totalRelations:= 0
	for i := 1; i <= groupSize; i++ {
		relation:= groupSize - i 
		totalRelations += relation
	}
	return totalRelations
}