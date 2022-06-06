package main

import "fmt"

func ArrayDiff(array1 []string, array2 []string) []string {
	for index, v1 := range array1 {
		for _, v2 := range array2 {
			if v1 == v2 {
				array1 = RemoveIndex(array1, index)
				ArrayDiff(array1, array2)
			}
		}
	}
	return unique(array1)
}
func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
func main() {
	test := []string{
		"test",
		"rest",
		"nest",
		"guest",
		"pest",
	}
	rest := []string{
		"best",
		"west",
		"nest",
		"guest",
	}
	fmt.Println(ArrayDiff(test, rest))
}
