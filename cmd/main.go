package main

import (
	"fmt"

	godyn "github.com/vedadiyan/godyn/pkg"
)

const TEST = `fn((1+2)*3*(8*(8+3*(2+2)))+10+(1+2)*3*(8*(8+3*(2+2)))+10+(1+2)*3*(8*(8+3*(2+2)))+10+(1+2)*3*(8*(8+3*(2+2)))+10+(1+2)*3*(8*(8+3*(2+2)))+10+(1+2)*3*(8*(8+3*(2+2)))+10+(1+2)*3*(8*(8+3*(2+2)))+10)`

// const TEST = `fn(value_a && false)`

func main() {
	fns := make(map[string]godyn.Expression)
	fns["fn"] = fn
	ctx := godyn.New(fns)
	value, err := ctx.Invoke(TEST)
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
}

func fn(args []any) (any, error) {
	return args[0], nil
}
