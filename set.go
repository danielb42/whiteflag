package whiteflag

import "reflect"

// Alias associates one long flag to a short flag. Also, a description for that flag pair
// can be specified which will be included in --help/-h output. All aliases must be declared
// before any call to FlagPresent(), GetBool(), GetInt() or GetString().
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

// SetIntDefault sets a default value for an int type flag. Defaults must be declared
// before any call to FlagPresent(), GetBool(), GetInt() or GetString().
func SetIntDefault(flag string, value int) {
	setDefault(flag, value)
}

// SetStringDefault sets a default value for a string type flag. Defaults must be declared
// before any call to FlagPresent(), GetBool(), GetInt() or GetString().
func SetStringDefault(flag, value string) {
	setDefault(flag, value)
}

func setDefault(flag string, value any) {
	if checkDefaults(flag) {
		friendlyPanic("default value for " + hyphenate(flag) + " already set")
	}

	t := reflect.TypeOf(value).String()
	defaults[t][resolve(flag)] = value
}
