// find all subsequences of a string

package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	matched, _ := regexp.MatchString(`\D`, s)

	if matched {
		return
	}

	Sottostringhe(s)
}

// example s = 123121212

func Sottostringhe(s string) map[string]int {

	var notrepeted []string
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		if i != len(runes)-1 && runes[i] < runes[i+1] {
			notrepeted = append(notrepeted, string(runes[i]))
		} else {
			notrepeted = append(notrepeted, string(runes[i]))
			break
		}
	}

	var newval string
	for _, v := range notrepeted {
		newval = newval + v
	}

	mp := make(map[string]int)
	var res string
	var count int
	for i := 0; i < len(newval); i++ {
		for j := i + 1; j <= len(newval); j++ {
			res = newval[i:j]
			if len(res) >= 2 {
				count = strings.Count(s, res)
				mp[res] = count
			}
		}
	}

	var keys []string

	for k, _ := range mp {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// sort string by length
	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) > len(keys[j])
	})

	for _, k := range keys {
		fmt.Println(k, mp[k])
	}

	return mp

}
