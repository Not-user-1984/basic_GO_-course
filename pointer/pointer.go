package pointer

import (
    "fmt"
	"time"
)

type Person struct {
	Name string 
	Age int
	lastVisited time.Time
}
func UpdatePersonWithLastVisited(p *Person) {
    p.lastVisited = time.Now()
}

func Show_memory(a string) {
    p := &a
    fmt.Println(a, p)
    a = "ihiiuhji"
}


func Strit(weekTeemp [3]string){
    for i,temp := range &weekTeemp{
        fmt.Println(i,temp)
    }
}
