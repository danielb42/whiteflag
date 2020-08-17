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

// GetBool is equivalent to FlagPresent() on a bool flag name.
func GetBool(flag string) bool {
	return FlagPresent(flag)
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
