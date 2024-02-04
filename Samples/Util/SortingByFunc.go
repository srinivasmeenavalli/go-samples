package main

import (
	"fmt"
	"sort"
)

/*
 sort strings by their length instead of alphabetically
*/

/* byLength type that is just an alias for the builtin []string type
 */
type byLength []string

/*
implement sort.Interface - Len, Less, and Swap -
on our type so we can use the sort package’s generic Sort function.
Len and Swap will usually be similar across
types and Less will hold the actual custom sorting logic
*/
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
