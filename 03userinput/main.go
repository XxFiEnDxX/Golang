package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the Pizze:")

	// coma ok || err ok

	input, _ := reader.ReadString('\n') //working as a try catch block .. input will store in "input" variable.
	//and if we face any type of error, then it will be store in _ or any other variable.
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type of rating %T", input)
}
