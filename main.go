package main

import (
	"fmt"
	"git-rebase-example/controller"
)

func main() {
	var newController controller.Controller
	newController = controller.Square
	fmt.Println(newController)
}
