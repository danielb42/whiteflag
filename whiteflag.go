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
	flags            = make(map[string]map[string]interface{})
	aliases          = make(map[string]flagAliasing)
)

func init() {
	flags["bool"] = make(map[string]interface{})
	flags["int"] = make(map[string]interface{})
	flags["string"] = make(map[string]interface{})

	aliases["h"] = flagAliasing{"help", "show this help text"}
}

// Alias associates one long flag to a short flag. Also, a description for that flag pair can be specified which
// will be included in --help/-h output. All aliases must be declared before any Check or Get function is called.
func Alias(short, long, description string) {
	if short == "h" || long == "help" {
		friendlyPanic("cannot re-define builtin -h or --help")
	}

	if len(short) > 1 {
		friendlyPanic("short flag aliased to " + hyphenate(long) + " must not be longer than 1 char")
	}

	if len(long) < 2 {
		friendlyPanic("long flag aliased to " + hyphenate(short) + " must be longer than 1 char")
	}

	if resolve(long) != long {
		friendlyPanic(hyphenate(long) + " is already aliased to another short flag")
	}

	if _, present := aliases[short]; present {
		friendlyPanic(hyphenate(short) + " already has an associated long flag")
	}

	aliases[short] = flagAliasing{long, description}
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

		if CheckBool(flag) || CheckInt(flag) || CheckString(flag) {
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

// CheckBool returns true when flag is present on the command line.
func CheckBool(flag string) bool {
	parseCommandLine()
	_, present := flags["bool"][resolve(flag)]
	return present
}

// CheckInt returns true when flag is present on the command line and
// is followed by an integer value.
func CheckInt(flag string) bool {
	parseCommandLine()
	_, present := flags["int"][resolve(flag)]
	return present
}

// CheckString returns true when flag is present on the command line and
// is followed by a string value.
func CheckString(flag string) bool {
	parseCommandLine()
	_, present := flags["string"][resolve(flag)]
	return present
}

// GetBool is equivalent to CheckBool()
func GetBool(flag string) bool {
	return CheckBool(flag)
}

// GetInt fetches the value of an integer flag, prints an error and exits
// the program if flag is missing or no integer value is specified.
func GetInt(flag string) int {
	parseCommandLine()

	if !CheckInt(flag) {
		friendlyPanic("integer flag " + hyphenate(flag) + " missing or no integer value given")
	}

	return flags["int"][resolve(flag)].(int)
}

// GetString fetches the value of an string flag, prints an error and exits
// the program if flag is missing or no string value is specified.
func GetString(flag string) string {
	parseCommandLine()

	if !CheckString(flag) {
		friendlyPanic("string flag " + hyphenate(flag) + " missing or no string value given")
	}

	return flags["string"][resolve(flag)].(string)
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
