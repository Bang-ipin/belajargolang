package main

import (
	"fmt"
	"log"
	"github.com/gorilla/mux" 
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var db *gorm.DB
var err error

func main() {
    // Please define your username and password for MySQL.
    db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/belajargolang?charset=utf8&parseTime=True")
    // NOTE: See we’re using = to assign the global var
    // instead of := which would assign it only in this function

    if err!=nil{
    	log.Println("Connection Failed to Open")
    }else{ 
    	log.Println("Connection Established")
    }
}