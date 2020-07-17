package main

import (
	"fmt"
	"git-rebase-example/controller"
)

func main() {
	var buttonArrow controller.Controller
	buttonArrow = controller.Right
	fmt.Println("========  ========")
	fmt.Printf("%+v\n", buttonArrow)
	fmt.Println("=================")
}
