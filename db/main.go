package main

import (
	"fmt"
	"financialmanager/db/db"
	"financialmanager/db/web"
)


func main() {
	fmt.Println("hello world!")
	db.Initdb()
	web.Init()
}

