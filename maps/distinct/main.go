package main
import "strings"

func countDistinctWords(messages []string) int {
	words := make(map[string] int)
	
	for i := 0; i < len(messages); i++ {
		
		phrases := strings.Fields(strings.ToLower(messages[i]))
		for j := 0; j < len(phrases); j++ {
			wordOfPhrase := phrases[j]
			if _, ok := words[wordOfPhrase]; ok {
				 words[wordOfPhrase] +=1
				
			} else {
				words[wordOfPhrase] =1
			}
			
		
		}
		
	}

	
	return len(words)
}