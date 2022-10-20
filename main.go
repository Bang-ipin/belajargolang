package main

import (
	"belajargolang/arraydata"
	"fmt"
)

func main() {
	// db, err := gorm.Open(mysql.Open("root:<>@tcp(127.0.0.1:3306)/belajargolang?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer db.Close()
	// 	fmt.Println("Connected Database Successfully!")
	x := 20
	fmt.Println(x)
	fmt.Println(Calculation(10, 4))
	arraydata.ArrayData()
}

func Calculation(number int, numberTwo int) int {
	return number * numberTwo
}
