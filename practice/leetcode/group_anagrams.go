package main

import (
	"fmt"
)

// https://leetcode.com/problems/group-anagrams/
func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ans := groupAnagrams(input)

	fmt.Println(ans)
}

func groupAnagrams(strs []string) [][]string {
	keyMap := getKeys(strs)

	dict := make(map[string][]string)
	ans := make([][]string, 0)

	for _, v := range strs {
		key := keyMap[v]
		if _, ok := dict[key]; ok {
			slice := dict[key]
			slice = append(slice, v)
			dict[key] = slice
		} else {
			slice := make([]string, 0)
			slice = append(slice, v)
			dict[key] = slice
		}
	}

	for _, value := range dict {
		ans = append(ans, value)
	}

	return ans
}

/*
	give a map with [actual][generatedKey]
*/

func getKeys(strs []string) map[string]string {
	dict := make(map[string]string)

	keys := make([]string, 0)
	for _, v := range strs {
		key := generateKey(v)
		keys = append(keys, key)
	}

	for i, v := range strs {
		dict[v] = keys[i]
	}
	return dict
}

/*
	s - input string
	keyformat - freq + char
*/

func generateKey(s string) string {
	var freq [26]int
	var key string
	for _, v := range s {
		freq[v-'a']++
	}

	for i := 0; i < 26; i++ {
		if freq[i] > 0 {
			key += fmt.Sprint(freq[i]) + string(i+97)
		}
	}

	return key
}
