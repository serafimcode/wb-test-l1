package main

import "fmt"

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	/*Множество- это набор неповторяющихся элементов,
	чтобы избавиться от дубликатов, создадим мапу,
	в которой ключ- это элементы множества, зн-я, по сути, что угодно,
	возьмем bool, с ним чуть проще реализовать полноценную стуктуру Set*/
	set := map[string]bool{}

	for _, el := range strings {
		set[el] = true
	}

	fmt.Printf("%+v", set)
}
