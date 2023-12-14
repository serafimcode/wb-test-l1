package main

import "fmt"

var callCounter uint

type human struct {
	age  uint
	name string
}

func (h human) sayHello() {
	callCounter++
	fmt.Printf("Hello! My name is %v, and I'am %v years old(call count: %v) \n", h.name, h.age, callCounter)
}

type action struct {
	human
	someField string
}

func main() {
	var a action
	a = action{human: human{
		age: 25, name: "Test human",
	}}
	/*Встраивание методов - это особенность языка GO,
	позволяющая нам из "встраивающей" структуры вызывать методы "встроенной" структуры*/
	a.sayHello()

	/*Этот вызов анологичен предыдущему*/
	a.human.sayHello()

	/*По аналогии с методами, мы можем обращаться к полям встроенной структуры*/
	fmt.Printf("Action name: %v, Action age: %v", a.name, a.age)
}
