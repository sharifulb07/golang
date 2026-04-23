// Something weird is happening in this code.
// What should be happening is that we create 2 separate constants:
// premiumPlanName and basicPlanName. Right now it looks like
// we're trying to overwrite one of them.

package main 

import "fmt"


func main(){
	const premiumPlanName="Premium Plan"
	const basicPlanName="Basic Plan"

	fmt.Println("plan: ", premiumPlanName)
	fmt.Println("plan: ", basicPlanName)




}
