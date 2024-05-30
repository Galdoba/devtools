package slice

func HasElement[T comparable](sl []T, elem ...T) bool {
	for _, el1 := range sl {
		for _, el2 := range elem {
			if el1 == el2 {
				return true
			}
		}
	}
	return false
}
