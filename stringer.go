package main

import(
	"fmt"
	
)

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)

}

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " Coffee Port"
}




func main(){
	//The error type is a predeclared identifier, like int/string
	//it is part of the "universe block". It is available wherever u r
	var err error
	err = ComedyError("The action is not allowed")
	fmt.Println(err)

	//Stringer interface. It is teidious to type so many printf statments
	//fmt.Printf("The volume is %0.2f litres", 3.141952)
	//instead of having so many Printf, you can use Stringer inteface
	//as long as your defined type has a String method returning a string.
	//it satisfies fmt.Stringer interface
	coffeePot := CoffeePot("LuxBrew")
	fmt.Println(coffeePot.String())

	//many functions in fmt package checks whether value passed satisfies
	//the Stringer interface, and call their String methods if so. This
	//includes Print, Println, Printf, and more. So, we can do:
	fmt.Println(coffeePot)
	fmt.Print(coffeePot, "\n")
	fmt.Printf("%s", coffeePot)

	//Empty interface will take any value, empty interface will accept all
	///types. Println has an empty interface as its argument.

	//func AcceptAnything(thing interface{}) {
	//}

	//func main(){
	//AcceptAnything(3.14159)
	//AcceptAnything("A string")
	//AcceptAnything(true)
	//AcceptAnything(Whistle("Toyco Canary'))
	
}
