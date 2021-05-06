// https://www.hackerrank.com/challenges/caesar-cipher-1/problem
package main

func caesarCipher(s string, k int32, n int) string {
	bytes := []byte(s)

	for i := 0; i < n; i++ {
		if isAlphaNumeric(bytes[i]) {
			tempChar := byte('a')
			if isUpperCase(bytes[i]) {
				tempChar = 'A'
			}
			// map a-zA-Z in 26 char range if exceed then do mod 26
			// so char can start from beginning again
			// eg: char z and k is 2 then output should be b (ascii of z + 2 mod 26 => char b)
			pos := (bytes[i] - tempChar + byte(k)) % 26
			bytes[i] = pos + tempChar
		}
	}
	return string(bytes)
}

func isAlphaNumeric(ch byte) bool {
	return ch >= 65 && ch <= 90 || ch >= 97 && ch <= 122
}

func isUpperCase(ch byte) bool {
	return ch >= 65 && ch <= 90
}
