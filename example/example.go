package main

import (
	wf "github.com/danielb42/whiteflag"
)

func main() {

	// optionally associate long flags with short flags and supply
	// a description which will be included in -h/--help output.
	//
	// ! aliasing must happen before any call to FlagPresent(), GetBool(), GetInt() or GetString()
	//
	wf.Alias("b", "mybool", "This is a flag.")
	wf.Alias("s", "mystring", "A string to print.")

	// check if a boolean flag -b or --mybool (as aliased above) is specified (=true).
	//
	if wf.FlagPresent("mybool") {
		println("The bool flag is set")
	}

	// check if -s or --mystring (as aliased above) is present. if so, print its string value. Errors
	// and quits if mystring value is not a string.
	//
	// ! we check FlagPresent() first as GetString() on an absent value would
	// ! trigger an error (thus making mystring a required flag)
	//
	if wf.FlagPresent("mystring") {
		println(wf.GetString("mystring"))
	}

	// if -x and -y are present, add their values. Errors and quits if x or y values are not integers.
	//
	// ! again, we check FlagPresent() first as GetInt() on an absent value would
	// ! trigger an error (thus making x and y required flags)
	//
	if wf.FlagPresent("x") && wf.FlagPresent("y") {
		x := wf.GetInt("x")
		y := wf.GetInt("y")
		sum := x + y
		println("sum of x and y:", sum)
	}

	// optionally assign default values for flags
	//
	// ! defaults must be set before any call to FlagPresent(), GetBool(), GetInt() or GetString()
	//
	wf.SetIntDefault("x", 2)
	wf.SetIntDefault("y", 3)
	x := wf.GetInt("x") // does not trigger an error now if x is not specified on command line
	y := wf.GetInt("y") // does not trigger an error now if y is not specified on command line
	sum := x + y
	println("sum of x and y:", sum)

	// you can check if an int/string flag is explicitly holding its types' zero value (0/"")
	// or if the flag is completely absent from the command line.
	//
	if !wf.FlagPresent("missing") {
		println("integer flag --missing is not present on command line")
	}
	// vs.
	if wf.FlagPresent("missing") && wf.GetInt("missing") == 0 {
		println("--missing has value 0")
	}
	// and
	if wf.FlagPresent("missing") && len(wf.GetString("missing")) == 0 {
		println("--missing holds the empty string")
	}
}
