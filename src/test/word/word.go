package main

import "unicode"

func main() {
	IsPalindrome("A man, a plan, a canal: Panama")
}

// 内存优化比较
// 折中比较

func IsPalindrome(s string) bool {
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	//n := len(letters) / 2
	//for i := 0; i < n; i++ {
	//	if letters[i] != letters[len(letters)-1-i] {
	//		return false
	//	}
	//}
	return true

}

// 折中比较

//func IsPalindrome(letters string) bool {
//	n := len(letters) / 2
//	for i := 0; i < n; i++ {
//		if letters[i] != letters[len(letters)-1-i] {
//			return false
//		}
//	}
//	return true
//
//}

// 原始比较
//func IsPalindrome(s string) bool {
//	for i := range s {
//		if s[i] != s[len(s)-1-i] {
//			return false
//		}
//	}
//	return true
//}
