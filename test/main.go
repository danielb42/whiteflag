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

	if wf.CheckInt("cflag") {
		fmt.Println("cflag =", wf.GetInt("cflag"))
	}

	if wf.CheckString("notint") {
		wf.GetInt("notint")
	}

	if wf.CheckBool("notstring") {
		wf.GetString("notstring")
	}

	if wf.CheckInt("notbool") {
		wf.GetBool("notbool")
	}

	if wf.CheckBool("testredefineh") {
		wf.Alias("h", "x", "y")
	}

	if wf.CheckBool("testredefinehelp") {
		wf.Alias("x", "help", "y")
	}

	if wf.CheckBool("testshorttoolong") {
		wf.Alias("xy", "xy", "xy")
	}

	if wf.CheckBool("testlongtooshort") {
		wf.Alias("x", "y", "z")
	}

	if wf.CheckBool("testlongalreadyaliased") {
		wf.Alias("x", "xx", "z")
		wf.Alias("y", "xx", "z")
	}

	if wf.CheckBool("testshortalreadyaliased") {
		wf.Alias("x", "xx", "z")
		wf.Alias("x", "yy", "z")
	}
}
