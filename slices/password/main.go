package main
import "unicode"
func isValidPassword(password string) bool  {
	valid := false 
	if len(password) >= 5 && len(password) <= 12 {
		upper := 0
		digit := 0
		for _, ch := range password {
			
			if unicode.IsUpper(ch)  {
				upper += 1
			} 
			if unicode.IsDigit(ch) {
				digit += 1
			}

			
	}
	if upper > 0 && digit > 0 {
		valid = true
	}
}
	return valid
}