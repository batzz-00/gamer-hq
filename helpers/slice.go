package helpers

type Comparable interface {
	Comparator() interface{}
}

func SliceHasStruct(haystack []Comparable, needle Comparable) Comparable {
	for _, item := range haystack {
		if item.Comparator() == needle.Comparator() {
			return item
		}
	}
	return nil
}

func SliceHas(haystack []Comparable, needle interface{}) Comparable {
	for _, item := range haystack {
		if item.Comparator() == needle {
			return item
		}
	}
	return nil
}
