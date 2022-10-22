package main

// #cgo LDFLAGS: -L. -lfancy
// #include "fancy.hpp"
// #cgo pkg-config: python3
// #include <Python.h>
// #include "fib.c"
import "C"
import (
	"fmt"
	"sync"
	"time"
)

func cpu_intensive(n int) int {
	fmt.Println("[go] dispatching cpu_intensive(", n, ")")
	return int(C.CPUIN(C.int(n)))
}

func io_intensive() int {
	fmt.Println("[go] dispatching io_intensive")
	return int(C.IOIN())
}

func dispatch_cpu_intensive(wg *sync.WaitGroup, n int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		cpu_intensive(n)
	}()
}

func dispatch_io_intensive(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		io_intensive()
	}()
}

func main() {
	// A dirty example launching a bunch of C++ calls
	tbeg := time.Now()
	var wg sync.WaitGroup
	//	dispatch_cpu_intensive(&wg, 30)
	//	dispatch_cpu_intensive(&wg, 35)
	//	dispatch_cpu_intensive(&wg, 40)
	//	dispatch_cpu_intensive(&wg, 45)
	//	dispatch_cpu_intensive(&wg, 45)
	//	dispatch_cpu_intensive(&wg, 50)
	//	dispatch_cpu_intensive(&wg, 30)
	//	dispatch_cpu_intensive(&wg, 35)
	//	dispatch_cpu_intensive(&wg, 40)
	//	dispatch_cpu_intensive(&wg, 45)
	//	dispatch_cpu_intensive(&wg, 45)
	//	dispatch_cpu_intensive(&wg, 50)
	//	dispatch_cpu_intensive(&wg, 30)
	//	dispatch_cpu_intensive(&wg, 35)
	dispatch_cpu_intensive(&wg, 40)
	dispatch_cpu_intensive(&wg, 45)
	dispatch_cpu_intensive(&wg, 45)
	dispatch_cpu_intensive(&wg, 50)
	dispatch_io_intensive(&wg)
	C.Py_Initialize()
	C.
		//	dispatch_io_intensive(&wg)
		//	dispatch_io_intensive(&wg)
		fmt.Println("[go] dispatched all tasks")
	wg.Wait()
	fmt.Println("[go] done (elapsed:", time.Since(tbeg).Seconds(), "seconds)")
}
