package main
import ("fmt"
 "strconv")

func fizzbuzz() string {
	result := ""
	for i := 1; i <= 100; i++ {
		if i%5 == 0 && i % 3 == 0 {
			result += "fizzbuzz\n"
		} else  if  i%5 == 0 {
			result += "buzz\n"

		} else if i%3 == 0 {
			result += "fizz\n"
		} else {
			result += strconv.Itoa(i) +"\n"
		}
	}
	return result
	
}

// don't touch below this line

func main() {
	fmt.Println(fizzbuzz())
}
