package main

import (
	"context"
	"fmt"
)

func updateKeyInCtx(ctx context.Context) {
	// val := ctx.Value("data")

	// val.(map[string]string)["Gaurav"] = "A"

	val := ctx.Value("Gaurav")

	lastName, _ := val.(string)

	lastName = "A"
	fmt.Println("Updated last name:", lastName)
}

func main() {
	parentCtx := context.Background()

	m := map[string]string{"Gaurav": "Agarwal", "Gautham": "Anand"}

	// ctx := context.WithValue(parentCtx, "data", m)

	ctx, _ := context.WithCancel(parentCtx)
	for key, val := range m {
		ctx = context.WithValue(ctx, key, val)
	}

	updateKeyInCtx(ctx)

	val := ctx.Value("Gaurav")

	fmt.Println(val)
}
