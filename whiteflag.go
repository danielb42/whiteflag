package whiteflag

import (
	"fmt"
	"os"
	"strconv"
)

type flagAliasing struct {
	long        string
	description string
}

var (
	cliAlreadyParsed bool

	flags    = make(map[string]map[string]interface{})
	defaults = make(map[string]map[string]interface{})
	aliases  = make(map[string]flagAliasing)
)

func init() {
	flags["bool"] = make(map[string]interface{})
	flags["int"] = make(map[string]interface{})
	flags["string"] = make(map[string]interface{})

	defaults["bool"] = make(map[string]interface{})
	defaults["int"] = make(map[string]interface{})
	defaults["string"] = make(map[string]interface{})

	aliases["h"] = flagAliasing{"help", "show this help text"}
}

func resolve(long string) string {
	for short, flagAliasing := range aliases {
		if long == flagAliasing.long {
			return short
		}
	}

	return long
}

func parseCommandLine() {
	if cliAlreadyParsed {
		return
	} else {
		cliAlreadyParsed = true
	}

	var (
		flag  string
		value interface{}
	)

	for index, token := range os.Args {
		if !isFlag(token) || index == 0 {
			continue
		}

		if token == "--help" || token == "-h" {
			printUsage()
		}

		if index < len(os.Args)-1 && !isFlag(os.Args[index+1]) {
			if isLongFlag(token) {
				flag, value = token[2:], os.Args[index+1]
			}
			if isShortFlag(token) {
				flag, value = string(token[1]), os.Args[index+1]
			}
		} else {
			if isLongFlag(token) {
				flag, value = token[2:], true
			}
			if isShortFlag(token) {
				flag, value = string(token[1]), true
			}
		}

		if FlagPresent(flag) {
			friendlyPanic(hyphenate(flag) + " specified more than once")
		}

		switch v := value.(type) {
		case bool:
			flags["bool"][resolve(flag)] = true
		case string:
			if intVal, err := strconv.Atoi(v); err == nil {
				flags["int"][resolve(flag)] = intVal
			} else {
				flags["string"][resolve(flag)] = v
			}
		}
	}
}

func hyphenate(flag string) string {
	if len(flag) == 0 {
		return ""
	} else if len(flag) == 1 {
		return "-" + flag
	}

	return "--" + flag
}

func isLongFlag(token string) bool {
	return len(token) > 2 && token[0:2] == "--"
}

func isShortFlag(token string) bool {
	return len(token) == 2 && token[0] == '-' && token[1] != '-'
}

func isFlag(token string) bool {
	return isLongFlag(token) || isShortFlag(token)
}

func friendlyPanic(reason string) {
	fmt.Println(reason)
	os.Exit(1)
}

func printUsage() {
	fmt.Printf("Usage: %s <flags>\n\nFlags:\n", os.Args[0])

	for short, flagDesc := range aliases {
		fmt.Printf("  %s  %s\t\t%s\n", hyphenate(short), hyphenate(flagDesc.long), flagDesc.description)
	}

	os.Exit(1)
}
