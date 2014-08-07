package main

import (
	"fmt"
	"time"
)

type Human struct {
	name  string
	age   int
	phone string
}

func main() {
	lxx := Human{"lxx", 20, "10086"}
	fmt.Printf("name %s  ,age %d , phone %s \n", lxx.name, lxx.age, lxx.phone)
	date := time.Now().Format(time.ANSIC)
	fmt.Println(date)
	fmt.Println(time.August)
}
