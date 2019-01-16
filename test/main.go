package main

import wf "github.com/danielb42/whiteflag"

func main() {
	if wf.CheckBool("s") {
		println("short bool set")
	}
}
