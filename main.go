package main

import (
	"github.com/jlangston/localmodtest/hello"
	"github.com/jlangston/localmodtest/utils"
	"fmt"
)

func main() {
	fmt.Println(hello.Hello("bob"))
	fmt.Println(utils.AddAndGreet("bob", 2, 3))
}
