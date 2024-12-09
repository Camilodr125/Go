package main

func bulkSend(numMessages int) float64 {
	total:= 0.0
	for i := 0; i < numMessages; i++ {
		total +=   float64(1.0+ float64(i)/float64(100))
	}
	return total
}