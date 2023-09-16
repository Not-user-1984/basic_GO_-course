package utils

import (
	"fmt"
	"os"
)

func CheckError(err error, message string) {
	if err != nil {
		fmt.Println(message, err)
		os.Exit(1)
	}
}

func ConnectToDBError(err error) {
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
}