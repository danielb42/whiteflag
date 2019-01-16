package main

import (
	"fmt"

	wf "github.com/danielb42/whiteflag"
)

func main() {
	wf.Alias("b", "bool", "An alias.")
	wf.Alias("c", "cflag", "Another alias.")

	wf.ParseCommandLine()

	if wf.CheckBool("b") && wf.GetBool("bool") {
		fmt.Println("bool set")
	}

	if wf.CheckInt("int") {
		fmt.Println("integer =", wf.GetInt("int"))
	}

	if wf.CheckString("notint") {
		wf.GetInt("notint")
	}

	if wf.CheckString("string") {
		fmt.Println("string =", wf.GetString("string"))
	}
}
