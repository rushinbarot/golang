/* Implement Data Structure to determine if string has uniq 
carecture without additional data structure 
*/

package main 

import (
	"fmt"
)

func main() {
	var input = []string{"abcd","aaa"}
	for _,str := range input {
		fmt.Println(str,"-->",isUnique(str), "00>",isUnique2(str))
	}
}

func isUnique(str string) bool {
	if len(str) > 128 {
		return false

	}
	var char_set [128]bool
	for _,val := range str{
		if char_set[val] {
			return false
		}
		char_set[val] = true
	}
	return true
}

func isUnique2(input string) bool {
	seen := make(map[rune]bool)
	for _,val := range input{
		if _, ok := seen[val]; ok {
			return false
		} else {
			seen[val] = true
		}
	}
	return true
}