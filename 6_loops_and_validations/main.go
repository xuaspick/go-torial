package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	cicleSum(50)
	doubleUntil(1000)
	// infiniteLoop() //? uncomment to see infinite loop (CRTL + C to stop 🔥)

	fmt.Printf("The square root of 4 is %v\n", sqrt(4))
	fmt.Printf("The square root of -4 is %v\n", sqrt(-4))

	for i := 2; i <= 3; i++ {
		limit := float64((i - 1) * 10)
		potencia := pow(3, float64(i), limit)

		if potencia == limit {
			fmt.Printf("The power of 3^%v is hiher than %v\n", i, limit)
		} else {
			fmt.Printf("The power of 3^%v is %v and lower than %v\n", i, potencia, limit)
		}
	}

	for i := 1; i < 5; i++ {
		rootableFloat := float64(i * i)
		fmt.Printf("The square root of %v is %v\n", rootableFloat, Task1Sqrt(rootableFloat))
	}

	switchExample()
	switchWithTime()
	switchWithNoCondition()

}

func cicleSum(intCicles int) {
	sum := 0
	for i := 0; i < intCicles; i++ { //? full for loop syntax
		sum += i
	}
	fmt.Printf("The sum of the first %v numbers is %v\n", intCicles, sum)
}

func doubleUntil(intSalida int) {
	sum := 1
	prevSum := 0
	for sum < intSalida { //?  for loop with only condition ()
		prevSum = sum
		sum += sum
	}
	fmt.Printf("The value before %v is %v\n", sum, prevSum)
}

func infiniteLoop() {
	for { //? infinite loop
		fmt.Println("Infinite loop")
	}
}

// ? if statement
func sqrt(x float64) string {
	if x < 0 { //? if statement
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// ? if statement with short statement
func pow(x, n, lim float64) (potencia float64) {
	if v := math.Pow(x, n); v < lim { //? short statement
		return v
	}
	return lim
	//? v is not visible here
}

func switchExample() {
	//? switch statement
	switch os := runtime.GOOS; os { //? only executes the first true case
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		//? freebsd, openbsd,
		//? plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func switchWithTime() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	tomorrow := today + 1
	fmt.Println("Today is", today)
	fmt.Println("Tomorrow is", tomorrow)
	fmt.Print("Saturday is: ")
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func switchWithNoCondition() {
	//? switch with no condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Print("Good morning!")
	case t.Hour() < 17:
		fmt.Print("Good afternoon.")
	default:
		fmt.Print("Good evening.")
	}
	fmt.Printf(" The time is %v hrs", t.Hour())

}
