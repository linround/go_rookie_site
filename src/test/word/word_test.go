package main

import (
	"fmt"
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

// //	func TestFrenchPalindrome(t *testing.T) {
// //		if !IsPalindrome("été") {
// //			t.Error(`IsPalindrome("été") = false`)
// //		}
// //	}
// //
func ExampleIsPalindrome() {
	fmt.Println("===============ExampleIsPalindrome")
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}

func TestNonPalindrome(t *testing.T) {
	if !IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

//func BenchmarkIsPalindrome(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		IsPalindrome("A man, a plan, a canal: Panama")
//	}
//}
