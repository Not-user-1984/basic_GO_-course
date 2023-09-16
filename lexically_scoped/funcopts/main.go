package main

import (
	"fmt"
)


type Person struct {
	Name string
	Surname string
	Age int
}




func main(){
    join := Person{Name: "Дима", Surname: "Dima", Age: 21}
	join.Age = 27
	fmt.Print(join)

}

type Item struct {
	NoOption string
	Parametr1 string
	Parametr2 int
}

func NewItem(opts ...option) *Item {
	i := &Item{
		NoOption: "usual",
		Parametr1: "default",
		Parametr2: 42,
	}
	for _, opt := range opts{
		opt(i)
	}
	return i
}
type option func(*Item)