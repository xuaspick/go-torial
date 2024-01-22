package main

import "fmt"

//? Make a function that returns the values of the variables declared below.
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

//? Function that returns the values of the variables declared below.
func getVariables() (smsSendingLimit int, costPerSms float32, hasPermission bool, userName string) {
	smsSendingLimit = 0
	costPerSms = 3.14
	hasPermission = true
	userName = "John Doe"
	return
}
