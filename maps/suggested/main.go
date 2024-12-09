package main

import "slices"
func findSuggestedFriends(username string, friendships map[string][]string) []string {
	
	suggestedfriends := []string{} 
	if len(friendships[username]) == 0 { // Caso de usuario sin amigos
		return nil
	}
	
		if fr, ok := friendships[username]; ok  {
					for j := 0; j < len(fr); j++ {
						names := friendships[fr[j]]
						for i := 0; i < len(names); i++ {
							friend := names[i]
							if friend != username && !slices.Contains(suggestedfriends, friend) &&!slices.Contains(fr, friend) {
								suggestedfriends = append(suggestedfriends, friend)
							}
						}
						
					}
				return suggestedfriends
		}
	return nil	
}