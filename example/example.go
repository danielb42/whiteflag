package main

import (
	wf "github.com/danielb42/whiteflag"
)

func main() {

	// optional - associate long flags with short flags and supply a description which
	// will be shown in -h/--help output.
	//
	// ! if a long flag has been aliased to the short flag, both can be queried equally
	// ! aliasing must happen before calling ParseCommandLine()
	//
	wf.Alias("b", "mybool", "This is a flag.")
	wf.Alias("s", "mystring", "A string to print.")

	// mandatory - build up the internal structures to Check/Get from.
	//
	// ! must happen before any Get/Check call.
	//
	wf.ParseCommandLine()

	// check if a boolean flag -b or --mybool (as aliased above) is specified (=true).
	//
	if wf.CheckBool("mybool") {
		println("The bool flag is set")
	}

	// check if -s or --mystring (as aliased above) is present and is followed by
	// a string value. if so, print the given string value.
	//
	// ! we CheckString() first as GetString() on an absent value would trigger a panic
	//
	if wf.CheckString("mystring") {
		println(wf.GetString("mystring"))
	}

	// if -x and -y are present and are followed by integer values, add their values.
	//
	// ! we CheckInt() first as GetInt() on an absent value would trigger a panic
	//
	if wf.CheckInt("x") && wf.CheckInt("y") {
		x := wf.GetInt("x")
		y := wf.GetInt("y")
		sum := x + y
		println("sum of x and y:", sum)
	}

	// you can check if an int/string flag is explicitly holding its types' zero value (0/"")
	// or if the flag is completely absent from the command line.
	//
	if !wf.CheckInt("missing") {
		println("integer flag --missing is not present on command line")
	} else if wf.GetInt("missing") == 0 {
		println("--missing has value 0")
	} else {
		println("--missing has value", wf.GetInt("missing"))
	}
	// Note: Specifying --missing with an other-than-int value will consume --missing as a
	// flag of that other type, thus CheckInt("missing") will still return false.
	// Same for specifying --missing alone, which would result in it being a bool flag.

}
