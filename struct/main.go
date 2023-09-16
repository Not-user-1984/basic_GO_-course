package main

import (
	"encoding/json"
	"os"
	"fmt"
	j "yandex_go/struct/json"
	e  "yandex_go/struct/utils"
	d "yandex_go/struct/db"
)


func main(){

	db, err := d.InitializeDB()
	e.ConnectToDBError(err)
	if db == nil {
        return
    }
	


	file ,err := os.Open("struct/json/person.json")
	e.CheckError(err, "Error opening file:")
	defer file.Close()

	

	var people []j.Person
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&people)
	e.CheckError(err, "Error decoding JSON:")

	for _, person := range people {
        db.Create(&person)
    }
	job := j.Job{
		Name:    "Software Engineer",
		Company: "TechCo",
		Pay:     75000,
	}

	person := j.Person{
		Name:  "dima",
		Age:   "33",
		Email: "dima@example.com",
		Job: j.Job{
			Name:    "Software Engineer",
			Company: "TechCo",
			Pay:     75000,
		},
	}
	
	err = d.CreatePersonWithJob(db, person, job)
	if err != nil {
		fmt.Println(err)
		return
	}




	var result []j.Person
	db.Find(&result)

	// Теперь данные хранятся в переменной result
	fmt.Println("Считанные данные из базы данных:")
	for _, person := range result {
		fmt.Println("Name:", person.Name)
		fmt.Println("Age:", person.Age)
		fmt.Println("Email:", person.Email)
		fmt.Println("Job Name:", person.Job.Name)
		fmt.Println("Job Company:", person.Job.Company)
		fmt.Println("Job Pay:", person.Job.Pay)
		fmt.Println()
	}
}