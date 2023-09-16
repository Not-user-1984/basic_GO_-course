package main

import (
	"fmt"
    p "yandex_go/map/product"
    t "yandex_go/map/taskTwoSum"
    d "yandex_go/map/duplicates"

)

// type MyMap map[string]string

type MyMapInt map[int]string

var s = make(map[string]string)
var sentence string =  "Πολύ λίγα πράγματα συμβαίνουν στο σωστό χρόνο, και τα υπόλοιπα δεν συμβαίνουν καθόλου"
var my_slice = []int{}


type NameBit struct{
    ID  int
    Name  string
    Price  float64
}


func count_symbol(sentence string){
    frequency := make(map[rune]int)
    for _, v := range sentence {
        frequency[v]++
    }
    for k, v := range frequency {
        fmt.Printf("Знак %c встречается %d раз \n", k, v)
    } 
    
}

func main(){

    x := make(MyMapInt, 5)
    x[1] = "один"
    x[2] = "два"
    x[3] = "три"
    fmt.Println(x)

	m := make(map[int]int)
    m[323] = 42424     
    m[121] = 35353        
    fmt.Println(m)
    myprice := make(map[int]NameBit)


    myprice[89898] = NameBit{
    ID: 1,
    Name: "ihihihi",
    Price : 5.898,
    }
    println(myprice)


    v, ok := x[2]
    y := m[2]
    y++
    m[2] = y

    fmt.Println(v,ok)
    fmt.Println(y)
    fmt.Println(len(m))

    if s != nil {            
        s["ключ"] = "после иннoцилизации"
    }
    fmt.Println(s)

    for k, v := range x {
        x[k] = "И " + v
        fmt.Println(x, v)
    }
    fmt.Println(x)
    
    count_symbol(sentence)

    fmt.Println(
        p.GetDelecacies(
        p.NewProduct()))

    fmt.Println(
        t.TwoSum(
            17,
            []int{2, 7, 11, 15}))
        
    fmt.Println(d.RemoveDulicates([]string{
        "cat",
        "dog",
        "bird",
        "dog",
        "parrot",
        "cat",
    }))


}

