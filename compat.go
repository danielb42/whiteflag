package whiteflag

// ParseCommandLine is a no-op and only kept for backward compatibility.
func ParseCommandLine() {}

// CheckBool returns true when flag is present on the command line.
// It is deprecated, use FlagPresent().
func CheckBool(flag string) bool {
	return FlagPresent(flag)
}

// CheckInt returns true when flag is present on the command line and
// is followed by an integer value.
// It is deprecated, use FlagPresent().
func CheckInt(flag string) bool {
	return FlagPresent(flag)
}

// CheckString returns true when flag is present on the command line and
// is followed by a string value.
// It is deprecated, use FlagPresent().
func CheckString(flag string) bool {
	return FlagPresent(flag)
}
