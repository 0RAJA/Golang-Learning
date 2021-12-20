package main

func main() {

}

func isPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func firstPalindrome(words []string) string {
	for i := range words {
		if isPalindrome(words[i]) {
			return words[i]
		}
	}
	return ""
}
