# whiteflag

[![pipeline status](https://gitlab.com/danielb42/whiteflag/badges/master/pipeline.svg)](https://gitlab.com/danielb42/whiteflag/commits/master)
[![GoDoc](https://godoc.org/github.com/danielb42/whiteflag?status.svg)](https://godoc.org/github.com/danielb42/whiteflag) 
[![Go Report Card](https://goreportcard.com/badge/github.com/danielb42/whiteflag)](https://goreportcard.com/report/github.com/danielb42/whiteflag) 
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)  
![Whiteflag Gopher](whiteflag.png)

A minimal but sane flag-package for people who just need some CLI flags in Golang projects, not command/argument/option-parsing frameworks for space ships. If you, too, waved a white flag on those, whiteflag is here to assist.

## What it does
- provides a `Check` and a `Get` method for Boolean, Integer and String type flags  
These can be utilized directly without further setup of each flag. 
- allows you to distinguish between absent and zero-valued flags
- `-h/--help` prints basic Usage/Help text (see [example](#with-long-flags))

## What it does not
whiteflag just doesn't try to be overly smart or versatile. For instance, if you need to define "required" flags or default values for absent flags, just use a few well-placed Get/Check calls yourself.

## Examples
Please have a look at the comprehensive [example source file](example/example.go).  

### Basic
The following snippet would print "gopher" when called with `-p gopher`.
```golang
package main

import wf "github.com/danielb42/whiteflag"

func main() {
    wf.ParseCommandLine()
    
    if wf.CheckString("p") {
        println(wf.GetString("p"))
    }
}
```

### With long flags
The next snippet will print the sum of two integers given through `-x` and `-y`.  
Let's also associate long flags to the short flags so we could equivalently run the snippet with `--first` and `--second`.  
Aliasing flags makes them known to the Usage/Help text generation.

```golang
package main

import wf "github.com/danielb42/whiteflag"

func main() {
    wf.Alias("x", "first",  "The first number.")
    wf.Alias("y", "second", "The second number.")

    wf.ParseCommandLine()

    if wf.CheckInt("x") && wf.CheckInt("y") {
        x := wf.GetInt("x")
        y := wf.GetInt("y")
        sum := x + y
        println("sum of x and y:", sum)
    }
}
```

Usage/Help text would be available through `-h/--help`:

```
Usage: ./example <flags>

Flags:
  -x  --first    The first number.
  -y  --second   The second number.
```

## License
MIT