package main

import(
	"fmt"

)


type Number int

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}


func main(){
	ten := Number(10)
	ten.Add(4)
	
	ten.Subtract(5)
	
	four := Number(4)
	four.Add(3)
	
	four.Subtract(2)
	//you can't add carFuel + busFuel, even if their underlying data type is the same

	//I can write functions to convert from Gallons to type Number int

	//But the code becomes too tedious with different coversions. Say I have milliliters
	//then i will have miliToGallons, GallonsToLitres, LitresToMillion. It is even annoying

	//Solution: Methods, so, I can have ToGallons method and it doesn't matter what the from-unit is
	//all it knows is that it will convert it to Gallons.

	//Define Methods:
	//func (m MyType) sayHi() {
	//fmt.Println("Hi")
	//}
	//
	//Call the method:
	//
	//value := MyType("...")
	//value.sayHi()
	

	
	
}
