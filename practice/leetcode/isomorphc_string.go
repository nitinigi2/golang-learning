package practice

// https://leetcode.com/problems/isomorphic-strings/

func isIsomorphic(s string, t string) bool {
	n1, n2 := len(s), len(t)
	var arr, brr [256]int

	if n1 != n2 {
		return false
	}

	for i := 0; i < n1; i++ {
		if arr[s[i]] != brr[t[i]] {
			return false
		}
		arr[s[i]] = i + 1
		brr[t[i]] = i + 1
	}
	return true
}
