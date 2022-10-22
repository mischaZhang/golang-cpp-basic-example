package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("/home/zdt/workspace/golang-cpp-basic-example/goroutines/fib.cpython-310-x86_64-linux-gnu.so")
	//p, err := plugin.Open("/home/zdt/workspace/golang-cpp-basic-example/goroutines/libfancy.so")
	if err != nil {
		panic(err)
	}

	i, _ := p.Lookup("fib")
	fmt.Println(i)
	fi := i.(func(int))
	fi(0)

	// c, _ := p.Lookup("IOIN")
	// fc := c.(func())
	// fc()
}
