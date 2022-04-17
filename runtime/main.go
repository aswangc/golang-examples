package main

import (
	"runtime/debug"
)

func main() {
	debug.SetMaxThreads(3)

}
