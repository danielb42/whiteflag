package main

import (
	"fmt"

	wf "github.com/danielb42/whiteflag"
)

func main() {
	wf.Alias("b", "bool", "An alias.")
	wf.Alias("c", "cflag", "Another alias.")

	if wf.CheckBool("b") {
		fmt.Println("bool =", wf.GetBool("b"))
	}

	if wf.CheckInt("int") {
		fmt.Println("integer =", wf.GetInt("int"))
	}

	if wf.CheckString("string") {
		fmt.Println("string =", wf.GetString("string"))
	}

	if wf.CheckString("notint") {
		wf.GetInt("notint")
	}
}
