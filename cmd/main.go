package main

import (
	"fmt"

	godyn "github.com/vedadiyan/godyn/pkg"
)

const TEST = `let(test, ld(x, x < 1))`

// const TEST = `fn(value_a && false)`

func main() {
	fns := make(map[string]godyn.Expression)
	fns["fn"] = fn
	ctx := godyn.New(fns)
	value, err := ctx.Invoke(TEST)
	if err != nil {
		panic(err)
	}
	_ = value
	fmt.Println(value)
	fmt.Println(("s" == "s") && true || false)
}

func fn(args []any) (any, error) {
	return args[0], nil
}
