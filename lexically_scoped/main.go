package main

import "fmt"

func fib() func() int {
	x1, x2 := 0, 1

	return func() int {
		x1, x2 = x2, x1+x2
		return x1
	}
}
func main(){
	f := fib()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}