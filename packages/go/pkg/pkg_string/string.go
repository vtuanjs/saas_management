package pkgstring

func OrString(val ...string) string {
	for _, v := range val {
		if v != "" {
			return v
		}
	}
	return ""
}

// PointerString returns a pointer to the given string. If the string is empty, it returns nil.
func PointerString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// DerefString returns the value of the given string pointer. If the pointer is nil, it returns an empty string.
func DerefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
