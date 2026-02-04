package main

import (
	"fmt"

	"github.com/ktsoator/omnidraw/demos/master/internal/router"
)

func main() {
	r := router.SetRoute()
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("run http server error:%v\n", err)
	}
}
