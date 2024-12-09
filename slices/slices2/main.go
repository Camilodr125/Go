package main

func createMatrix(rows, columns int) [][]int {
	matrix := [][]int{}
	for i := 0; i < rows; i++ {
		values := []int{}
		for j := 0; j < columns; j++ {
			values = append(values, i*j)
			
		}
		matrix = append(matrix, values)
	}
	return matrix
}