package main
type cost struct {
	day int
	value float64
}
func getCostsByDay(costs []cost) []float64 {
	costByday := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i] 
		day := cost.day
		for  day >= len(costByday){
			costByday = append(costByday, 0.0)
			
		}

		costByday[day] += cost.value
	}
	return costByday
}