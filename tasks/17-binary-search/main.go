package main

import "fmt"

/*Бинарный поиск для отсортированной последовательности*/
func binarySearch(arr []int, searchElem int) int {
	lIdx := 0
	rIdx := len(arr) - 1

	for lIdx <= rIdx {
		/*Выбираем элемент примерно в середине массива*/
		mid := (lIdx + rIdx) / 2
		midElem := arr[mid]

		/*Если искомый элемент равен элменту в середине- поиск завершен,
		возвращаем индекс элемента*/
		if searchElem == midElem {
			return mid
		}
		/*Если искомый элемент меньше элемента в середине,
		изменяем правую границу поиска на индекс, предшествующий
		середине*/
		if searchElem < midElem {
			rIdx = mid - 1
		} else {
			/*Если искомый элемент больше элемента в середине,
			изменяем левую границу поиска на индекс, следующий за
			серединой*/
			lIdx = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{-4, -1, 0, 3, 6, 10}
	/*Ищем элемент который есть*/
	fmt.Println(binarySearch(arr, 6))
	/*Ищем элемент которого нет*/
	fmt.Println(binarySearch(arr, 11))
}
