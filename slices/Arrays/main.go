package main

func getMessageWithRetries(first,second, third string) ([3]string, [3]int) {
	cost1:= len(first)
	cost2 := cost1 + len(second)
	cost3 := cost2 + len(third)
	total := [3]int{cost1, cost2, cost3}
	messages := [3]string{first, second,third}

	return messages, total
	 
}