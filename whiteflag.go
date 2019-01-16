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
	flags   = make(map[string]map[string]interface{})
	aliases = make(map[string]flagAliasing)
)

// Alias associates a long flag with a given short flag which functions identically
// to the short version. Also, a description for these flags can be specified which
// will be shown in --help/-h output. Aliasing must happen before calling ParseCommandLine().
func Alias(short, long, description string) {
	if resolve(long) != long {
		panic(hyphenate(long) + " is already aliased to another flag")
	}

	if _, ok := aliases[short]; ok {
		panic(hyphenate(short) + " already has an associated long flag")
	}

	aliases[short] = flagAliasing{long, description}
}

func resolve(long string) string {
	for short, flagDesc := range aliases {
		if long == flagDesc.long {
			return short
		}
	}

	return long
}

// ParseCommandLine scans the command line for supplied flags
// and builds the internal structures to Check/Get from. Must
// be called before any Check/Get.
func ParseCommandLine() {
	flags["bool"] = make(map[string]interface{})
	flags["int"] = make(map[string]interface{})
	flags["string"] = make(map[string]interface{})

	var (
		flag  string
		value interface{}
	)

	for index, token := range os.Args {
		if !isFlag(token) || index == 0 {
			continue
		}

		if len(aliases) > 0 && (token == "--help" || token == "-h") {
			printUsage()
			os.Exit(0)
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
			panic(hyphenate(flag) + " specified more than once")
		}

		switch value.(type) {
		case bool:
			flags["bool"][resolve(flag)] = true
		case string:
			if intVal, err := strconv.Atoi(value.(string)); err == nil {
				flags["int"][resolve(flag)] = intVal
			} else {
				flags["string"][resolve(flag)] = value.(string)
			}
		}
	}
}

// CheckBool returns true when flag is present on the command line.
func CheckBool(flag string) bool {
	_, present := flags["bool"][resolve(flag)]
	return present
}

// CheckInt returns true when flag is present on the command line and
// is followed by an integer value.
func CheckInt(flag string) bool {
	_, present := flags["int"][resolve(flag)]
	return present
}

// CheckString returns true when flag is present on the command line and
// is followed by a string value.
func CheckString(flag string) bool {
	_, present := flags["string"][resolve(flag)]
	return present
}

// GetBool is equivalent to CheckBool() but panics when flag is not set.
func GetBool(flag string) bool {
	if !CheckBool(flag) {
		panic("boolean flag " + hyphenate(flag) + " missing or no boolean value given")
	}

	return true
}

// GetInt fetches the value of an integer flag, panics if
// flag is missing or no integer value is specified.
func GetInt(flag string) int {
	if !CheckInt(flag) {
		panic("integer flag " + hyphenate(flag) + " missing or no integer value given")
	}

	return flags["int"][resolve(flag)].(int)
}

// GetString fetches the value of a string flag, panics if
// flag is missing or no string value is specified.
func GetString(flag string) string {
	if !CheckString(flag) {
		panic("string flag " + hyphenate(flag) + " missing or no string value given")
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

func panic(reason string) {
	fmt.Println(reason)
	os.Exit(1)
}

func printUsage() {
	fmt.Printf("Usage: ./%s <flags>\n\nFlags:\n", os.Args[0])

	for short, flagDesc := range aliases {
		fmt.Printf("  %s  %s\t\t%s\n", hyphenate(short), hyphenate(flagDesc.long), flagDesc.description)
	}
}