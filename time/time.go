package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format("01-02-2006")) // this date is the standard format
	fmt.Println(t.Format("01-02-2006 Monday"))
	fmt.Println(t.Format("01-02-2006 15:04:05 Monday"))
	// above date time and day are standard, remember them
	fmt.Println(runtime.NumCPU()) // prints number of cpus for execution
}
