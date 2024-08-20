package string

// removeElementByValue removes the first occurrence of the specified value from the string slice.
func RemoveElementByValue(slice []string, value string) []string {
	// Find the index of the value
	index := -1
	for i, v := range slice {
		if v == value {
			index = i
			break
		}
	}

	// If the value was not found, return the original slice
	if index == -1 {
		return slice
	}

	// Create a new slice with one less element
	newSlice := make([]string, 0, len(slice)-1)

	// Copy elements except the one at the found index
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)

	return newSlice
}

// removeElement removes an element from the string slice at the given index.
func RemoveElement(slice []string, index int) []string {
	// Check if the index is valid
	if index < 0 || index >= len(slice) {
		return slice
	}

	// Create a new slice with one less element
	newSlice := make([]string, 0, len(slice)-1)

	// Copy elements except the one at the given index
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)

	return newSlice
}
