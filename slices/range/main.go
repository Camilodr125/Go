package main

func indexOfFirstBadWord(msg []string, badWords []string) int  {
	
	
		for i, ms := range msg {
			for  _, ba := range badWords {
				if ms == ba {
				return i
			} 
		}
	}
		return -1
}
	
	
