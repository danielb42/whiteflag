package main

import (
	"fmt"

	wf "github.com/danielb42/whiteflag"
)

func main() {
	wf.Alias("b", "bool", "An alias.")
	wf.Alias("c", "cflag", "Another alias.")

	wf.SetIntDefault("defint", 123)
	wf.SetStringDefault("defstr", "foobar")

	if wf.FlagPresent("testdefaults") {
		fmt.Println("defint =", wf.GetInt("defint"))
		fmt.Println("defstr =", wf.GetString("defstr"))
	}

	if wf.FlagPresent("testrequired") {
		wf.GetInt("required")
	}

	// Unaliased short
	if wf.FlagPresent("a") {
		fmt.Println("bool =", wf.GetBool("a"))
	}

	// Aliased short
	if wf.FlagPresent("b") {
		fmt.Println("bool =", wf.GetBool("b"))
	}

	if wf.FlagPresent("testboolfalse") {
		fmt.Println("bool =", wf.GetBool("foobarfoobar"))
	}

	if wf.FlagPresent("int") {
		fmt.Println("integer =", wf.GetInt("int"))
	}

	// Unaliased long
	if wf.FlagPresent("string") {
		fmt.Println("string =", wf.GetString("string"))
	}

	// Aliased long
	if wf.FlagPresent("cflag") {
		fmt.Println("cflag =", wf.GetInt("cflag"))
	}

	if wf.FlagPresent("notint") {
		wf.GetInt("notint")
	}

	if wf.FlagPresent("notstring") {
		wf.GetString("notstring")
	}

	if wf.FlagPresent("testredefineh") {
		wf.Alias("h", "x", "y")
	}

	if wf.FlagPresent("testredefinehelp") {
		wf.Alias("x", "help", "y")
	}

	if wf.FlagPresent("testshorttoolong") {
		wf.Alias("xy", "xy", "xy")
	}

	if wf.FlagPresent("testlongtooshort") {
		wf.Alias("x", "y", "z")
	}

	if wf.FlagPresent("testlongalreadyaliased") {
		wf.Alias("x", "xx", "z")
		wf.Alias("y", "xx", "z")
	}

	if wf.FlagPresent("testshortalreadyaliased") {
		wf.Alias("x", "xx", "z")
		wf.Alias("x", "yy", "z")
	}

	// same type
	if wf.FlagPresent("testdefaultalreadyset1") {
		wf.SetIntDefault("defint", 456)
	}

	// across types
	if wf.FlagPresent("testdefaultalreadyset2") {
		wf.SetStringDefault("defint", "barfoo")
	}
}
