package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
  hi := Hello("bob")
  if hi != "hello bob" {
   t.Fail()
  }
}
