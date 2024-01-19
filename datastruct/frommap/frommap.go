package frommap

import "sort"

// Keys_MSS_Sorted - возвращает сортированые ключи map[string]string
// если reverse = true, задает обратный порядок
func Keys_MSS_Sorted(stringStringMap map[string]string, reverse bool) []string {
	keys := []string{}
	for k := range stringStringMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	if reverse {
		keys = reverseStringSlice(keys)
	}
	return keys
}

func reverseStringSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// Keys_MSS_Sorted - возвращает сортированые ключи map[string]string
// если reverse = true, задает обратный порядок
func KeysUsed(stringMap interface{}) map[string]bool {
	keys := make(map[string]bool)
	switch stringMap := stringMap.(type) {
	default:
		return nil
	case map[string]string:
		for k := range stringMap {
			keys[k] = false
		}
	case map[string]int:
		for k := range stringMap {
			keys[k] = false
		}
	case map[string]bool:
		for k := range stringMap {
			keys[k] = false
		}
	case map[string]float64:
		for k := range stringMap {
			keys[k] = false
		}

	}
	return keys
}
