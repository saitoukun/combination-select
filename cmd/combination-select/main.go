package main

import (
	"combination-select/internal"
	"fmt"
)

func main() {
	const port = ":8080"
	fmt.Print("Listen" + port)
	internal.Handler(port)
}
