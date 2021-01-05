package types

// StrSlice for running comparisions
type StrSlice []string

// Has returns true if slice contains passed value
func (list StrSlice) Has(a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// ArrayContains helper function to check if array contains target value
func ArrayContains(list *[]string, target string) bool {
	for _, item := range *list {
		if item == target {
			return true
		}
	}
	return false
}
