package pkgcsv

import "strings"

const CsvBoolTrue = "TRUE"
const CsvBoolFalse = "FALSE"

// ParseCSVBool parses a string value from a CSV into a boolean.
// It returns two boolean values:
// - The first return value is the parsed boolean (true or false).
// - The second return value indicates whether parsing was successful.
//
// The function handles the following cases:
// - "TRUE" returns (true, true) - true value, successfully parsed
// - "FALSE" or empty string returns (false, true) - false value, successfully parsed
// - Any other string returns (false, false) - indicating parsing failure
func ParseCSVBool(value string) (bool, bool) {
	value = strings.TrimSpace(value)

	switch value {
	case CsvBoolTrue:
		return true, true
	case CsvBoolFalse, "":
		return false, true
	default:
		return false, false
	}
}
