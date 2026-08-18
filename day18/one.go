package main


import (
	"fmt"
)

// interface declaration
type PaymentProcessor interface{
	Pay(amount float64) error
}

// stripe
type Stripe struct{}

func (s Stripe)Pay(amount float64)error{

	fmt.Println("Payment completed via stripe: ", amount)
	return nil
}

// paypal
type Paypal struct{}

func (s Paypal) Pay(amount float64) error{
	fmt.Println("Payment completed via Payment: ", amount)
	return nil
}


func ProcessPayment(p PaymentProcessor, amount float64) error{

	err:=p.Pay(amount)

	if err!=nil{
		fmt.Println("Payment failed: ", err)
		return err 
	}

	fmt.Println("Payment successfully completed")
	return nil 


}


func main(){
	stripe:=Stripe{}
	paypal:=Paypal{}

	ProcessPayment(stripe, 2000)
	ProcessPayment(paypal, 300)
}