// Third problem

package main

import "fmt"

func main() {
	for x := 1; x <= 100; x++ {
		// Is divisible by 3 and 5, print FizzBuzz
		if x%3 == 0 && x%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if x%3 == 0 { // Is divisible by only 3, print Fizz
			fmt.Println("Fizz")
		} else if x%5 == 0 { // Is divisible by only 5, print Buzz
			fmt.Println("Buzz")
		}
	}
}
