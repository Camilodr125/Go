package main

func maxMessages(thres int) int  {
	totalCost := 0
	for i := 0; ; i++ {
		totalCost += (100 + i)
		if thres < totalCost  {
			return i
		} 
	}
	
}