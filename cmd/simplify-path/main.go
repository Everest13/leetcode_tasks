package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/home/user/Documents/../Pictures"

	path = "/../"

	path = "/a/./b/../../c/" // /c

	res := simplifyPath(path)
	fmt.Println(res)
}


/*
	71. Simplify Path
	Meduim
	https://leetcode.com/problems/simplify-path/description/
	Runtime 3ms Beats 53.96%, Memory 4.55MB Beats 19.31%
 */
func simplifyPath(path string) string {

	parts := strings.Split(path, "/")
	res := []string{}

	n := len(parts)
	for i:=1; i<n; i++ {
		part := parts[i]
		if part == "" || part == "." {
			continue
		}

		if parts[i] == ".." {
			if len(res)>0  {
				res = res[0:len(res)-1]
			}
			continue
		}

		res = append(res, "/" + part)
	}

	if len(res) == 0 {
		return "/"
	}

	resStr := ""
	for i:=0; i<len(res); i++ {
		resStr+=res[i]
	}

	return resStr
}

func simplifyPath2(path string) string {

	parts := strings.Split(path, "/")
	res := []string{}

	n := len(parts)
	for i:=1; i<n; i++ {
		part := parts[i]
		if part == "" || part == "." {
			continue
		}

		if parts[i] == ".." {
			if len(res)>0  {
				res = res[0:len(res)-1]
			}
			continue
		}

		res = append(res, "/" + part)
	}

	if len(res) == 0 {
		return "/"
	}

	resStr := strings.Join(res, "")
	return resStr
}