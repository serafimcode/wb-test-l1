package main

import (
	"fmt"
	"time"
)

/*Секция офиса, доступ в которую ограничен, допустим, по электронному ключу*/
type PrivateSection struct {
	records map[time.Time]string
}

/*Посетители, которые могут посещать приватную секцию*/
type AllowedVisitor interface {
	hasAccess() bool
}

/*
У посетителя есть доступ в приватную секцию. Доступ могут обнулить,
но электронный ключ все еще может быть у этого посетителя(допусти сотрудик уволился
и забыл вернуть ключ)
*/
func (j *PrivateSection) visit(vi AllowedVisitor, name string) {
	if vi.hasAccess() {
		j.records[time.Now()] = name
		fmt.Println("Visitor entered")
	} else {
		fmt.Println("This user can't enter this section")
	}
}

/*У сотрудника есть доступ приватной секции*/
type Worker struct {
	name      string
	isAllowed bool
}

func (w Worker) hasAccess() bool {
	return w.isAllowed
}

/*
Приглашенный в офис человек, допустим, пришел на собес.
В приватную секцию у него нет доступа, он вообще про нее не знает.
*/
type Guest struct {
	name string
}

/*Человек прошел собес, взяли на испытательный срок, выдадим доступ*/
type GuestAdapter struct {
	guest *Guest
}

func (ga GuestAdapter) hasAccess() bool {
	return true
}

func main() {
	section := PrivateSection{records: map[time.Time]string{}}
	w := Worker{name: "Pavel", isAllowed: true}
	section.visit(w, w.name)

	gAdapter := GuestAdapter{guest: &Guest{name: "Anna"}}
	section.visit(gAdapter, w.name)

	for key, value := range section.records {
		fmt.Printf("%v entered at %v\n", value, key)
	}
}
