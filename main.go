package main

import (
	"fmt"
	"gin-skeleton/app/lib"
	"gin-skeleton/router"
)

func main() {
	defer func() {
		fmt.Println(recover())
	}()
	lib.InitMysql("defualt")
	r := router.Router()
	err := r.Run()
	if err != nil {
		fmt.Println("server start err", err)
		return
	}
}
