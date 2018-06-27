package utils

import "fmt"

func SliceString(s []interface{}) []string {
	ss := make([]string, 0)
	for i, v := range s {
		ss[i] = fmt.Sprintf("%v", v)
	}
	return ss
}
func SliceInterface(t []string) []interface{} {
	s := make([]interface{}, len(t))
	for i, v := range t {
		s[i] = v
	}
	return s
}

func StringInSlice(a string, list ...string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
