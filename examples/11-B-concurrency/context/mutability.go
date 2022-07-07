package main

import (
	"context"
	"fmt"
	"time"
)

var (
	key   = "Harish"
	value = "Sakhineti"
)

func setKeyValue(ctx context.Context) context.Context {
	return context.WithValue(ctx, key, value)
}

func printValue(ctx context.Context) {
	val := ctx.Value(key)

	fmt.Println(val)
}

func main() {
	ctx, _ := context.WithCancel(context.Background())

	ctxWithVal := setKeyValue(ctx)

	go printValue(ctxWithVal)
	go printValue(ctxWithVal)

	go func() {
		time.Sleep(1 * time.Second)
		printValue(ctxWithVal) // ?
	}()

	value = "S"
	updatedCtx := setKeyValue(ctxWithVal)

	printValue(updatedCtx)
	time.Sleep(2 * time.Second)
}
