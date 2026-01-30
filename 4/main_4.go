package main

func SliceSubtracting(minuend, subtractible []string) []string {
	set := make(map[string]bool, len(subtractible))
	for _, value := range subtractible {
		set[value] = true
	}

	difference := make([]string, 0, len(minuend))
	for _, value := range minuend {
		if !set[value] {
			difference = append(difference, value)
		}
	}

	return difference
}