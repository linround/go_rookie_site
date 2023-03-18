package word

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("dbd") {
		t.Error(`出错了`)
	}
	if !IsPalindrome("kaak") {
		t.Error(`("kaak") = 出错了`)
	}
}
func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}
func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}
func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}
