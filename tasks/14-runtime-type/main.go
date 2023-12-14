package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str interface{} = "Hello"
	var integer interface{} = 5
	var channel interface{} = make(chan string)

	/*Пакет reflect предоставляет средства для работы с данными неизвестного типа.
	TypeOf возвращает тип переданного параметра*/
	fmt.Println(reflect.TypeOf(str))
	fmt.Println(reflect.TypeOf(integer))
	fmt.Println(reflect.TypeOf(channel))
}
