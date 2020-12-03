package int

// Connects second slice to first slice
func SliceConnect(s1 []int, s2 []int) []int {
	for i := 0; i < len(s2); i++ {
		s1 = append(s1, s2[i])
	}
	return s1
}

// returns duplicates in s1 and s2
func FindDuplicates(s1 []int, s2 []int) []int {
	var i3 = 0
	var dup []int
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

// Print all int in int slice with new line
func PrintlnAll(s1 []int) {
	for i := 0; i < len(s1); i++ {
		println(s1[i])
	}
}

// Print all int in int slice
func PrintAll(s1 []int) {
	for i := 0; i < len(s1); i++ {
		print(s1[i])
	}
}
