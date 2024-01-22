package main

import "fmt"

var c, python, java bool //? A var declaration can include initializers, one per variable.
//! foo := 42 //? Short variable declarations can be used inside functions only, this can't be used outside a function, use VAR instead.

//? variables without an explicit initial value are given their zero value.
var (
	defInt   int
	defStr   string
	defBool  bool
	defFloat float64
)

const Pi = 3.14 //? Constants are declared like variables, but with the const keyword.
const (
	//? Create a huge number by shifting a 1 bit left 100 places.
	//? In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	//? Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
	// Small = 1 << 1 //! This is the same as the line above.
)

func passMeAnInt(x int) int {
	return x*10 + 1
}

func passMeAFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	var i, j int = 1, 2       //? A var declaration can include initializers, one per variable.
	tuViejaEsta := "en tanga" //? Short variable declarations can be used inside functions.
	fmt.Println(i, j, c, python, java, tuViejaEsta)
	ExecuteTask()  //? variables returned from a function.
	ExecuteTask2() //? Happy birthday!
	ExecuteTask3() //? It prints different type handlings of a variable.

	//? You can access the variables declared in tasks.go
	fmt.Printf("el ToBe de tasks.go es (%T) %v\n", ToBE, ToBE)
	fmt.Printf("el MaxInt de tasks.go es (%T) %v\n", MaxInt, MaxInt)
	fmt.Printf("el z de tasks.go es (%T) %v\n", z, z)

	//? Zero values
	fmt.Printf("%v %v %v %q\n", defInt, defBool, defFloat, defStr)

	//? infering types by short declaration
	var someInt int = 42
	fmt.Printf("The type of someInt is %T\n", someInt) //? %T is used to print the type of a variable
	anotherInt := 42
	fmt.Printf("The type of anotherInt is %T\n", anotherInt) //? %T is used to print the type of a variable
	imaginaryNumber := 0.867 + 0.5i
	fmt.Printf("The type of imaginaryNumber is %T\n", imaginaryNumber) //? %T is used to print the type of a variable

	//? Numeric Constants
	// Pi = 3.1416 //! Constants can't be reassigned.
	fmt.Println(Pi)
	fmt.Println(passMeAnInt(Small))
	// fmt.Println(passMeAnInt(Big)) //! This will overflow.
	fmt.Println(passMeAFloat(Small))
	fmt.Println(passMeAFloat(Big))

}
