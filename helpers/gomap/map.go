package gomap

func MapHasAnotherMap[K, V comparable](m, sub map[K]V) bool {
	for k, vsub := range sub {
		if vm, found := m[k]; !found || vm != vsub {
			return false
		}
	}
	return true
}
