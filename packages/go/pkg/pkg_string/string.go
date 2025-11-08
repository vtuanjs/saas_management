package pkgstring

func OrString(val ...string) string {
	for _, v := range val {
		if v != "" {
			return v
		}
	}
	return ""
}
