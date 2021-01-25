Varaible declartions in Go:

Arrays:
Size of the array is set at declaration
var myNumberArray [5]int = [5]int{, 2, 3 ,4, 5}
var myStringArray [5]string = [5]sting{"Pam", "Jim", "Dwight", "Michael", "Angela"}

myStringArray[4] = "Ryan"


----------------

Slices:
Note: you have to use append function, as the size of the slice is not, it grows as items are appended.
my mySlice []string = []string{"Pam", "Jim", "Dwight", "Michael", "Angela"}

var mySlice []string

mySlice = append(mySlice, "Pam")

you can create a slice from an existing array for example:

var myStringArray [5]string = [5]string{"Pam", "Jim", "Dwight", "Michael", "Angela"}

my myStringSlice := myStringArray[0:3]

myStringSlice will be {"Pam", "Jim, "Dwight"}

If you modify any element of the slice, it will modify the original array:

myStringSlice[0] = "Karen"

Array myStringSlice will become {"Karen", "Jim", "Dwight", "Michael", "Angela"}



---------------

Maps: Key/Value pairs

var myMaps := make(map[string]int)

//you can crete map literal
var myMaps = map[string]int{"Toronto": 1, "Montreal": 2, "Vancouver": 3}

for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}


//or you can create map and then assign values


myMap := make(map[string]int)
	
myMap["Vancouver"] = 3
myMap["Toronto"] = 1
myMap["Montreal"] = 2



Note: just declaring map does not create it, so, the following code will give you error


var myMap map[string]int
	
myMap["Vancouver"] = 3
myMap["Toronto"] = 1
myMap["Montreal"] = 2

You will get: panic: assignment to entry in nil map

after declaring it, you need to create the map:

myMap = make(map[string]int)


---------------------------
Structs

You can define multiple types as a group using struct

type Employee struct {
    Name     string
    Id	     float64
    Title    string
    Address  
{

type Customer struct {
     Name     string
     Rate     float64
     Status   bool
     Address
}

type Address struct {
     Street  	string
     City	string
     State	string
     PostalCode	int
     Country	string
}


Note, instead of having address fields in Employee and Customer struct, we can seperate it out and use the reference Address.
We can modify it as such

//main.go file
//import(
	"github.com/fsayed/office"
)
.
.
.
var employee := office.Employee{ Name: "Jim Helpert" }
employee.Id = 524859554
employee.Title = "Assistant General Manager, Scranton Branch"
employee.Address.Street = "123 Oakwood Street"
employee.Address.City = "Scranton"
employee.Address.State = "PA"
employee.Address.PostalCode = 12345
employee.Address.Country = "USA"


-----------------------------
var name string

name = "Jim"

var name string = "Jim"

var name = "Jim"

name := "Jim"

You can't use := assignment in package level, it will give you: syntax error: non-declaration statement outside function body


var number int

number := 5

var number int = 5

var number = 5

------------------------------
Custom Types:

You can create your own typss. We already looked at structs, those are custom types

ex: type Customer struct {
    ...
    ...
}


type Id float64
type ZipCode float64


var idNum := Id(12345)
var zcNum := ZipCode(12345)

Now, idNum is of type Id and zcNum is of type ZipCode. Their both underlying type is float64.
You can create methods for your types as well:

func (i Id) printId(){
     fmt.Println("Employee ID:", i)
}

//method call

idNum.printId()

You can also update your value by changing the receive to pointer

func (i *Id) updateID(){
     *i += 100000000
}


Note, the method call is still:

idNum.updateID()

Go will implicitly use pointer in updateID method. So, the above code will chaneg the value of idNum

----------------------------------------
Panic:
You can call panic function to end your progam. 


func main() {
 // some condition does not match
 panic("exiting the program")
}

However calling panic should be reserved when you encountered an unmanageable error. You should defer a recover() function


func calmDown(){
fmt.Println(recover())
}

func main() {
defer calmDown()
panic("error")

}


Note that reover() takes in interface {} and returns a interface {} type value. So, you can't calll function on your return value. For Ex:

func calmDown() {
 p := recover()
 //the code below will throw an error
 fmt.Println(p.Error())

 //user type assertion to get the underlying type back

  err, ok := p.(error)

if ok {
       fmt.Println(err)
 }
}

func main() {
 defer calmDown()
 err := fmt.Errorf("there's an error")
 panic(err)
}
