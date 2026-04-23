package main

import "fmt"

func main() {
	var smsSendingLimit int = 1
	var costPerSms float64 = 0.00000
	var hasPermission bool = false

	var username string = ""

	fmt.Printf("%v    %f %v %q \n", smsSendingLimit, costPerSms, hasPermission, username)
}