package whiteflag

// (DEPRECATED) This call can be removed.
func ParseCommandLine() {}

// (DEPRECATED) Use FlagPresent().
func CheckBool(flag string) bool {
	return FlagPresent(flag)
}

// (DEPRECATED) Use FlagPresent().
func CheckInt(flag string) bool {
	return FlagPresent(flag)
}

// (DEPRECATED) Use FlagPresent().
func CheckString(flag string) bool {
	return FlagPresent(flag)
}
