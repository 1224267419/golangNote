package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	start, count := 0, 0
	hashtable := map[uint8]int{}
	for i := 0; i < len(s); i++ {

		_, ok := hashtable[s[i]]
		if ok {
			count = max(i-start, count)
			start = hashtable[s[i]] + 1
		}
		hashtable[s[i]] = i

	}
	count = max(count, len(s)-start-1)
	return count
}

func main() {
	//	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	a := 10
	a -= 1
	fmt.Println(a)
}
