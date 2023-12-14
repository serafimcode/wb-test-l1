package main

import "fmt"

/*
Слайс это структура которая содержит указатель на массив,
длину это массива, и длину самого слайса. Когда мы делаем слайс
от массива, мы по сути манипулируем указателем и длиной слайса.
slice[:i] - мы все еще указываем на первый элемент массива,
но изменяем длину слайса.
slice[i+1:] - а здесь мы сдвигаем указатель с начала массива на i+1.
Но при этом мы не меняем сам массив. Чтобы его изменить мы предаем указатель
*/
func deleteElem(slice *[]int, i int) {
	*slice = append((*slice)[:i], (*slice)[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	deleteElem(&slice, 2)
	fmt.Println(slice)
}
