package mySort

import "sort"

func sortStr(strs []string) []string {
	sort.Strings(strs)
	return strs
}
