package main

import (
	"fmt"

	"master/internal/router"
)

func main() {
	server := router.SetRoute()
	if err := server.Run(":8080"); err != nil {
		fmt.Printf("run http server error:%v\n", err)
	}
}
