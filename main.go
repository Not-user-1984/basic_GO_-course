package main

import (
	"fmt"
	p "yandex_go/pointer"
	s "yandex_go/slice"
)
var weekTeemp =[3]string{"sss", "aaaa","dddd"}


func main() {
	my_int := "строка"
	p.Show_memory(my_int)
	fmt.Println(my_int)
	p.Strit(weekTeemp)
	x := p.Person{
		Name: "alex",
		Age: 25,
	}
	p.UpdatePersonWithLastVisited(&x)
	fmt.Println(x)
	s.DemoSlice()
	s.InsreaseSlice()
	s.Append_and_reverce()
	s.String_reverce()
} 