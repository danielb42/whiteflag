# whiteflag

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/danielb42/whiteflag)
![Tests](https://github.com/danielb42/whiteflag/workflows/Tests/badge.svg)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/danielb42/whiteflag)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/danielb42/whiteflag)](https://pkg.go.dev/github.com/danielb42/whiteflag)
[![Go Report Card](https://goreportcard.com/badge/github.com/danielb42/whiteflag)](https://goreportcard.com/report/github.com/danielb42/whiteflag)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)  
![Whiteflag Gopher](https://raw.githubusercontent.com/danielb42/whiteflag/master/whiteflag.png)

A sane flag-package for gophers who just need some CLI flags in Golang projects, not command-structuring frameworks for space ships. If you waved a white flag on the usual whoppers, `whiteflag` is here to assist.

## Features

- provides a method `FlagPresent` to check for specified flags, and methods `Get(Bool|Int|String)` to access their values (these methods can be utilized directly without further setup of each flag)
- allows you to distinguish between absent and zero-valued flags
- `-h/--help` prints basic generated Usage/Help text (see examples)
- Default values for flags can be specified
- Required flags can be achieved implicitly (see examples)

## Examples

Please have a look at the comprehensive [example source file](example/example.go).  

### Basic

The following snippet would print "gopher" when called with `-p gopher`.

```golang
package main

import wf "github.com/danielb42/whiteflag"

func main() {
    if wf.FlagPresent("p") {
        println(wf.GetString("p"))
    }
}
```

### With long+required+default flags and nice 'Usage' output

The next snippet will print the sum of two integers given through `-x` and `-y`. For `y` we specify a default value. Let's also associate long flags to the short flags so we could equivalently run the snippet with `--first` and `--second`. Aliasing flags makes them known to the Usage/Help text generation.

```golang
package main

import wf "github.com/danielb42/whiteflag"

func main() {
    wf.Alias("x", "first",  "The first number.")
    wf.Alias("y", "second", "The second number.")
    wf.SetIntDefault("y", 42)

    // we don't do a FlagPresent() check on x und y before Get'ting them so
    // the program will exit if x is not specified, thus making x 'required'.
    // For a missing y flag, the default value of 42 would be used.

    x := wf.GetInt("x")
    y := wf.GetInt("y")
    sum := x + y
    println("sum of x and y:", sum)
}
```

For the snippet above the following Usage/Help text would be available through `-h/--help`:

```golang
Usage: ./example <flags>

Flags:
  -x  --first    The first number.
  -y  --second   The second number.
```

## License

MIT
