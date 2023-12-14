package main

import "fmt"

func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, -5, 0, 0.5, 7, 5, 24.5, -21.0, 32.5}
	/*Создаем мапу, ключами которой будут нижние интервалы тепературы с шагом в 10,
	а значениями будут искомые подмножества*/
	temperatureMap := make(map[int][]float64)

	for _, v := range data {
		/*Определяем ключ в мапе. Приводим к int, чтобы округлить вниз.
		Домножаем на 10, чтобы получить нужный интервал*/
		key := int(v/10) * 10

		/*Когда слайс только инициализирован, то его зн-е будет nil,
		а длина 0. Проверяем, если длина больше 0, добавляем элемент по ключу,
		иначе создаем слайс с одним значением*/
		if len(temperatureMap[key]) > 0 {
			temperatureMap[key] = append(temperatureMap[key], v)
		} else {
			temperatureMap[key] = []float64{v}
		}
	}

	fmt.Printf("%+v", temperatureMap)
}
