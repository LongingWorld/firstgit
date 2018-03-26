package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h *Human) SayHi() {
	fmt.Printf("Hi,i am  %s you can call me on %s \n", h.name, h.phone)
}

/*func (h *Human) SayHi() {
    fmt.Printf("Hi,i am  %s you can call me on %s \n", h.name, h.phone)
}*/

func (h *Human) Sing(lyrics string) {
	fmt.Println("Do Ri Mi Fa Su La Xi ...", lyrics)
}

func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}

func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi,i am  %s you can call me on %s \n", e.name, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungeChap interface {
	SayHi()
	Sing(lyrics string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(lyrics string)
	SpendSalary(amount float32)
}

func (h Human) String() string { //实现fmt.Stringer.String interface
	return "◇" + h.name + " - " + strconv.Itoa(h.age) + "years - " + h.phone + "℡◇"
}

func main() {
	mike := Human{"Mike", 18, "16602100121"}
	tom := Student{Human{"Tom", 12, "16601210210"}, "Henan", 100}
	lucy := Employee{Human{"Lucy", 20, "18802100121"}, "IBM", 10000}

	fmt.Println("~~~~~~~值类型T调用引用类型*T的Method~~~~~~~")
	mike.SayHi()
	tom.SayHi()
	lucy.SayHi()
	fmt.Println("~~~~~~~值类型T调用引用类型*T的Method~~~~~~~")

	fmt.Println("Tom is :", tom)   //interface 作为函数的参数
	fmt.Println("Mike is :", mike) //

	var i Men

	i = &mike

	fmt.Println("This is Mike:")
	i.SayHi()
	i.Sing("My love")

	i = &tom
	fmt.Println("This is Tom,a student:")
	i.SayHi()
	i.Sing("Good Good study ,day day up")

	fmt.Println("Let's use a slice of Men and see what happens")
	slice := make([]Men, 3)
	slice[0], slice[1], slice[2] = &mike, &tom, &lucy

	for _, value := range slice {
		value.SayHi()
	}

	fmt.Println("Hello World!")

}
