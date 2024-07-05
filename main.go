package main

import "fmt"

/*
Task 1 :
Create a program to provide functionality to create employees and provide
Getter and Setter functions. Employees should have basic fields like id, name, email, etc.

Task 2 :
Create a program to create orders(id, amount, shipping fee) and recalculate the order's shipping
fees depending on the order amount(amount > 255 shipping charges = 0) Display the order info before
and after recalculation.
*/

type Employee struct {
	id    int
	name  string
	email string
}

type Order struct {
	id          int
	amount      int
	shippingFee float64
}

func NewOrder(id int, amount int, shippingFee float64) *Order {
	return &Order{
		id:          id,
		amount:      amount,
		shippingFee: shippingFee,
	}
}

func (order *Order) recalculateOrder() {
	if order.amount < 255 {
		order.shippingFee = 0
	}
}

func NewEmployee(id int, name string, email string) *Employee {
	return &Employee{
		id: id, name: name, email: email,
	}
}
func (e *Employee) GetID() int {
	return e.id
}

func (e *Employee) SetID(id int) {
	e.id = id
}

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetEmail() string {
	return e.email
}

func (e *Employee) SetEmail(email string) {
	e.email = email
}

func main() {
	/*
		fmt.Println("Hello World")
			fmt.Println("Perimeter of rectangle is :", PerimeterRectangle(2, 3))
			fmt.Println("Circumference of Circle is:", PerimeterCircle(4))
			fmt.Println("Bob Replied : ", Hey("WHAT IS THE STATUS ?"))
			fmt.Println("Number of Prime Numbers: ", PrimeNumbers(0))

	*/
	e := NewEmployee(100234, "Darshana", "darshana@gmail.com")
	fmt.Println("Employee Details")
	fmt.Println("id: ", e.GetID())
	fmt.Println("name: ", e.GetName())
	fmt.Println("email: ", e.GetEmail())
	fmt.Println()
	//Setting new value
	e.SetID(2)
	e.SetName("Divya")
	e.SetEmail("divya@zopsmart.com")
	fmt.Println()
	fmt.Println("Modified employee details")
	fmt.Println("Employee details")
	fmt.Println("id: ", e.GetID())
	fmt.Println("name: ", e.GetName())
	fmt.Println("email: ", e.GetEmail())
	fmt.Println()
	//Creating Order
	o := NewOrder(101, 240, 1200)
	//Initial Order Details
	fmt.Println("Order details")
	fmt.Println("id: ", o.id)
	fmt.Println("amount: ", o.amount)
	fmt.Println("shippingFee: ", o.shippingFee)
	fmt.Println()
	o.recalculateOrder()
	fmt.Println("Modifies Shipping Charges")
	fmt.Println("id: ", o.id)
	fmt.Println("amount: ", o.amount)
	fmt.Println("shippingFee: ", o.shippingFee)
}
