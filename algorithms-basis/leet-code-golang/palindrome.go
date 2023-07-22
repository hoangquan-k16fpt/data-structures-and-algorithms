package main
import (  
    "fmt"
)

func longestPalindrome(s string) string {
    str := ""
	for i:=0; i<len(s)-1; i++{
		for j:=len(s); j>0; j++{
			if s[:i] == s[j-1:]{
				result := s[i:j]
			} else {

			}
		}
	}
	return str
}



func main() {
	s := "Hello World"
	longestPalindrome(s)
}