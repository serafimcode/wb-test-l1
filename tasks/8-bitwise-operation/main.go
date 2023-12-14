package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number: ")
	var target uint64
	fmt.Scan(&target)

	fmt.Println("Enter 0 or 1: ")
	var changeBit uint64
	fmt.Scan(&changeBit)

	fmt.Println("Enter a bit number to change: ")
	var targetBit int
	fmt.Scan(&targetBit)

	fmt.Println("Your number before:")
	fmt.Println(target)
	fmt.Printf("%b\n", target)

	if changeBit != 1 && changeBit != 0 {
		panic("Provided wrong value")
	}

	/*Создаем битовую маску равную по длине(количеству бит) целевому числу*/
	var mask uint64

	if changeBit == 1 {
		/*
			1. x << n - битовое смещение числа x на n битов влево. Т.е. все биты нашего числа смещаются влево,
			а справа добавляются нули. Смещаем единицу на тот бит, который мы хотим изменить.
		*/
		mask = 1 << targetBit
		/*
			Выполняем битовое ИЛИ, которое для каждого бита результирующего числа даст единицу,
			если хотябы у одного из операндов в этом бите стоит единица. 1011 | 0100 = 1111
		*/
		target |= mask
	}
	if changeBit == 0 {
		/*
			Маска создаемтся аналогичным образом за счет смещения, но теперь мы инвертируем число,
			чтобы в маске 0 был только в том месте, в котором нам нужно, а в остальных битах будет единица
		*/
		mask = ^(1 << targetBit)
		/*
			Выполняем битовое И, которое для каждого бита результирующего числа даст единицу,
			если у обоих операндов в этом бите стоит единица. 1111 | 1011 = 1011
		*/
		target &= mask
	}

	fmt.Println("Your number after:")
	fmt.Println(target)
	fmt.Printf("%b\n", target)
}
