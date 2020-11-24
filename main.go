package main

import (
	"runtime"

	"github.com/cs3210-fall2020/gsh/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
