package main

import (
    "fmt"
    "os"
    "path/filepath"
)

// func fact(n int) int {
//     res := 1
//     for n > 0 {
//         res *= n
//         n--
//     }
//     return res
// }

func PrintAllFiles(path string){
    files, err := os.ReadDir(path)
    if err != nil {
        fmt.Println("unable to get list of files", err)
        return
    }
    for _, f := range files{
        filename := filepath.Join(f.Name())
        fmt.Println(filename)
        if f.IsDir() {
            PrintAllFiles(filename)
        }
    }
}


func main(){
	// fac := fact(45)
	// print(fac)
    PrintAllFiles("map/")
    
}