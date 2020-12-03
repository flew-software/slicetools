package string

// Connects second slice to first slice
func SliceConnect(s1 []string, s2 []string) []string {
	for i := 0; i < len(s2); i++ {
		s1 = append(s1, s2[i])
	}
	return s1
}

// returns duplicates in s1 and s2
func FindDuplicates(s1 []string, s2 []string) []string {
	var i3 = 0
	var dup []string
	for i := 0; i < len(s1); i++ {
		for i2 := 0; i2 < len(s2); i2++ {
			if s2[i2] == s1[i] {
				dup = append(dup, s1[i])
				i3++
			}
		}
	}
	return dup
}

// Print all strings in string slice with new line
func PrintlnAll(s1 []string) {
	for i := 0; i < len(s1); i++ {
		println(s1[i])
	}
}

// Print all stings in string slice
func PrintAll(s1 []string) {
	for i := 0; i < len(s1); i++ {
		print(s1[i])
	}
}
