package main

import (
	"fmt"
	"math/cmplx"
)

//? variable declarations may be "factored" into blocks, as with import statements.

var (
	ToBE   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// ? Make a function that returns the values of the variables declared below.
func ExecuteTask() {
	//? get variables here.
	smsSendingLimit, costPerSms, hasPermission, userName := getVariables()
	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		costPerSms,
		hasPermission,
		userName,
	)
}

// ? Function that returns the values of the variables declared below.
func getVariables() (smsSendingLimit int, costPerSms float32, hasPermission bool, userName string) {
	smsSendingLimit = 0
	costPerSms = 3.14
	hasPermission = true
	userName = "John Doe"
	return
}

func ExecuteTask2() {
	congrats := "Happy birthday!"
	fmt.Println(congrats)
}

func ExecuteTask3() { //? Change the value of the variables below so the type is float64
	penniesPerText := 5
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)          //? %T is used to print the type of a variable
	fmt.Printf("The type of penniesPerText is %T\n", float64(penniesPerText)) //? %T is used to print the type of a variable
	penniesPerText = 5.0                                                      //! won't change to INT, gets 'ignored' and stays FLOAT64
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)
}
