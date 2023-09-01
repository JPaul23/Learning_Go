package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Println("Doing something====>: myKey's value is", ctx.Value("key"))
	anotherCtx := context.WithValue(ctx, "key", "my second context")
	doAnotherThing(anotherCtx)

	fmt.Println("Doing something====>: myKey's value is", ctx.Value("key"))
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "my first context")
	doSomething(ctx)
}

func doAnotherThing(ctx context.Context) {
	fmt.Printf("Doing Another===> myKey's value is %s\n", ctx.Value("key"))
}
