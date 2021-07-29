package mySort

import "sort"

func SortStr(strs []string) []string {
	sort.Strings(strs)
	return strs
}
