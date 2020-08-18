package whiteflag

// GetInt fetches the value of an int flag. It prints an error and exits
// the program if flag is missing or no int value is specified.
func GetInt(flag string) int {

	value, isInt := getValueOf(flag).(int)

	if !isInt {
		friendlyPanic("integer flag " + hyphenate(flag) + " missing or no integer value given")
	}

	return value
}

// GetString fetches the value of a string flag. It prints an error and exits
// the program if flag is missing or no string value is specified.
func GetString(flag string) string {

	value, isString := getValueOf(flag).(string)

	if !isString {
		friendlyPanic("string flag " + hyphenate(flag) + " missing or no string value given")
	}

	return value
}

// GetBool checks if flag is present on the command line. It prints an error and exits
// the program if flag is not used in a boolean context (i.e. is followed by a value).
func GetBool(flag string) bool {

	value := getValueOf(flag)

	if value == nil {
		return false
	}

	if _, isBool := value.(bool); !isBool {
		friendlyPanic("flag " + hyphenate(flag) + " is followed by a non-bool value")
	}

	return true
}

func getValueOf(flag string) interface{} {

	parseCommandLine()

	for _, flagsOfType := range flags {
		flag, present := flagsOfType[resolve(flag)]
		if present {
			return flag
		}
	}

	for _, defaultsOfType := range defaults {
		dflt, present := defaultsOfType[resolve(flag)]
		if present {
			return dflt
		}
	}

	return nil
}
