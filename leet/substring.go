package main

func main() {
	isSubsequence("abc", "abcdeg")
}

func isSubsequence(s string, t string) bool {
	var i, j int
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
		
	}
	if i == len(s) {
		return true
	} else {
		return false
	}
}
