package main

import "fmt"

func join(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		res += sep
		res += strs[i]
	}
	return res
}

func main() {
	test_strs := []string{"Macro", "Reus"}
	fmt.Println(join("$", test_strs...))
}
