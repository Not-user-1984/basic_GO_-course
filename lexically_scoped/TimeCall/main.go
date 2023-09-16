package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func countCall(f func(string)) func(string) {
	cnt := 0
	funcname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()

	return func(s string) {
		cnt++
		fmt.Printf(
			"Функция %s вызова %d раз\n", funcname , cnt)
	}
}

func metricTimeCall(f func(string)) func(string) {
	return func(s string) {
		start := time.Now()
		f(s)
		fmt.Println("Время выполнения", time.Now().Sub(start))
	}
}

func my_print(s string) {
	fmt.Println(s)
}

func PrintAllFilesWithFilterClosure(path string, filter string){

	var walk func(string)

	walk =func(path string) {
		files, err := os.ReadDir(path)
		if err != nil{
			fmt.Println("unable to get list of files", err)
			return
		}
		for _, f :=  range files {
			filename :=  filepath.Join(path, f.Name())
			if strings.Contains(filename, filter) {
				fmt.Println(filename)
			}
			if f.IsDir(){
				walk(filename)
			}
		}
	}

	walk(path)
}
func main (){
	filter := ".go"
	path := "/Users/diplug/go_projeсt/stady/yandex_go/lexically_scoped"
	PrintAllFilesWithFilterClosure(path, filter)
	coutedPrint := countCall(my_print)
	coutedPrint("привет, мир")
	coutedPrint("Привет, лошара")

	countAndMetricPrint := metricTimeCall(coutedPrint)
	countAndMetricPrint("привет, мир")
	countAndMetricPrint("привет, лошара")
}