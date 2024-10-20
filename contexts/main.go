package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	// fmt.Println(ctx)
	// add a key value pair to the context
	ctx = context.WithValue(ctx, "key", "value")
	// retrieve the value from the context
	fmt.Println(ctx.Value("key"))
}
