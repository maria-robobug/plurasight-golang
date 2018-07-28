package main

import (
	"fmt"

	"github.com/maria-robobug/plurasight/code-school/packages/model"
)

func main() {
	jumperList := model.GetList()
	for _, jumper := range jumperList {
		fmt.Println(jumper.Jump())
	}
}
