package sets

type Set[T comparable] struct {
	propertyMap map[T]bool
}

func New[T comparable]() *Set[T] {
	s := Set[T]{}
	s.propertyMap = make(map[T]bool)
	return &s
}

// func (s *Set[T]) New(item T) {
// 	s.propertyMap = make(map[T]bool)
// }

// func (s *Set[T]) propertyType() T {
// 	var nullVal T
// 	return nullVal
// }

//ContainsElement - check if element exist in property map
func (s *Set[T]) ContainsElement(e T) bool {
	_, exist := s.propertyMap[e]
	return exist
}

//AddElement - adds element to set
func (s *Set[T]) AddElement(e T) {
	if !s.ContainsElement(e) {
		s.propertyMap[e] = true
	}
}

//DeleteElement - deletes element from propertyMap
func (s *Set[T]) DeleteElement(e T) {
	delete(s.propertyMap, e)
}

//Intersect - return an intersectionSet that consist of
//intersection of s and s2. original set is traversed
//through and checked if s2 contains its elements
func (s *Set[T]) Intersect(s2 *Set[T]) *Set[T] {
	intersectionSet := New[T]()
	for val := range s.propertyMap {
		if s2.ContainsElement(val) {
			intersectionSet.AddElement(val)
		}
	}
	return intersectionSet
}

//Union - returns a unionSet that consists of a union of
//original set and s2
func (s *Set[T]) Union(s2 *Set[T]) *Set[T] {
	unionSet := New[T]()
	for value := range s.propertyMap {
		unionSet.AddElement(value)
	}
	for value := range s2.propertyMap {
		unionSet.AddElement(value)
	}
	return unionSet
}
