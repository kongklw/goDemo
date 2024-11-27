package hello_test

import (
	"example/goDemo/hello"
	"testing"
)

func TestHello(t *testing.T) {
	name := "your name is xiaohong"
	if hello.GetName("xiaohong") != name {
		t.Fatal("wrong getname func")
	}
}
