package utils

// StringPointerToNULL converts a string pointer to NULL if the string is empty
func StringPointerToNULL(s *string) *string {
	if s != nil && *s == "" {
		return nil
	}
	return s
}

// IntPointerToNULL converts an int pointer to NULL if the int is 0
func IntPointerToNULL(i *int) *int {
	if i != nil && *i == 0 {
		return nil
	}
	return i
}
