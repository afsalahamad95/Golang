package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// ctx := context.Background()
	// // fmt.Println(ctx)
	// // add a key value pair to the context
	// ctx = context.WithValue(ctx, "key", "value")
	// // retrieve the value from the context
	// fmt.Println(ctx.Value("key"))
	// ctx = context.WithValue(ctx, "key", "12345")
	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		fmt.Println("context ends")
	// 		return
	// 	default:
	// 		fmt.Println("working")
	// 	}
	// 	<-time.After(500 * time.Millisecond)
	// }
	// // the loop runs indefinitely
	// to prevent that, use timeout
	cont, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // cancel after time limit
	// this cancel can be used to terminate the contex
	for {
		select {
		case <-cont.Done():
			fmt.Println("time out")
			return
		default:
			fmt.Println("working")
		}
		time.Sleep(1 * time.Second)
	}
}
