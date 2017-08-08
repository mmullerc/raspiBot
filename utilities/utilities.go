package utilities

//checks if error string has a value and panics
func CheckForStringErr(s string) {
	if len(s) > 0 {
		panic(s)
	}
}
