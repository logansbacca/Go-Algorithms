package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	longestword := "" 
	for scanner.Scan() {
		line := scanner.Text()
		wordsInline := strings.Split(line, " ")
		for i:=0; i<len(wordsInline); i++ {
			if len(wordsInline[i]) > len(longestword){
				longestword = wordsInline[i]
			}
		}
		}
		fmt.Println(longestword)
	}

