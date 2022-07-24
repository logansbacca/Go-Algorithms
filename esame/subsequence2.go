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
	a, _ := regexp.MatchString(`\D`, s)
	if a {
		return
	}
	Sottostringhe(s)

}

func Sottostringhe(s string) map[string]int {

	mp := make(map[string]int)

	var ns []string

	for i := 0; i < len(s)-1; i++ {
		if s[i] < s[i+1] {
			ns = append(ns, string(s[i]))
		} else {
			ns = append(ns, string(s[i]))
			break
		}
	}
	var res string
	for i := 0; i < len(ns)-1; i++ {
		for j := i + 1; j <= len(ns); j++ {
			a := ns[i:j]
			for _, v := range a {
				if len(a) >= 2 {
					res = res + v
				}
			}

			x := strings.Count(s, res)
			if len(res) >= 2 {
				mp[res] = x
				res = ""
			}

		}

	}
	var rs [] string

	for i := range mp {
		rs= append(rs,i)
	}
	// sort slice of strings by length 
	sort.Slice(rs, func(i,j int) bool {
		return len(rs[i]) > len(rs[j])
	})

	for _,v :=range rs {
		fmt.Println(v, mp[v])
	}
	return mp

}
