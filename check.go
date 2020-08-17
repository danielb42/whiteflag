package whiteflag

// FlagPresent checks if flag was specified on the command line.
func FlagPresent(flag string) bool {

	parseCommandLine()

	for _, flagsOfType := range flags {
		if _, present := flagsOfType[resolve(flag)]; present {
			return true
		}
	}

	return false
}

func checkDefaults(flag string) bool {

	for _, defaultsOfType := range defaults {
		if _, present := defaultsOfType[resolve(flag)]; present {
			return true
		}
	}

	return false
}
