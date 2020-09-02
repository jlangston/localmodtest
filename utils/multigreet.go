package utils

import (
	"github.com/jlangston/localmodtest/hello"
	"strconv"
)

func AddAndGreet(name string, a, b int) string {
	return hello.Hello(name) + " " + strconv.Itoa(a + b)
}
