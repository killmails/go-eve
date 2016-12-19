package util

import "fmt"

func SliceItoa(i []int) (s []string) {
	fmt.Println(i)
	for _, v := range i {
		s = append(s, fmt.Sprintf("%v", v))
	}
	return
}
