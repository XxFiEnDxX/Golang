package main

import "fmt"

var PublicVar = "Amit" //public variable start with capital letters in go

func main() {
	var username string = "Amit Kumar Singh"
	fmt.Println(username)
	fmt.Printf("Variable type : %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("Variable type : %T \n", isLoggedin)

	var unsignInt uint8 = 255 //unsigned 8-bit integers. Range: 0-255
	fmt.Println(unsignInt)
	fmt.Printf("Variable type : %T \n", unsignInt)

	var signInt int8 = 127 //signed 8-bit integers. Range: -128-127
	fmt.Println(signInt)
	fmt.Printf("Variable type : %T \n", signInt)

	//byte - alias for uint8
	//rune - alias for int32

	var smallFloat float64 = 123.12323523442634645845754
	fmt.Println(smallFloat)
	fmt.Printf("Variable type : %T \n", smallFloat)

	//default value and some alias

	var newInt int
	fmt.Println(newInt) //default value for int is 0(just like in java).

	var newStr string
	fmt.Println("x" + newStr + "x") //empty string

	// implicit type

	var newVar = "String type" //it will assign newVar string datatype
	fmt.Println(newVar)

	//no "var" keyword
	number := 123 // := is a walrus oprator. only work inside a method ...not work for global variables
	fmt.Println(number)
	fmt.Println(PublicVar)
}
