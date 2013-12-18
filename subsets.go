package main

import "fmt"

func main() {
	out("Starting subset generator\n")
	s := []string { "1", "2", "3" }
	gray := generateGrayCode(3)
	prettyPrint("gray", gray, true)
	t := generateAllSubSets(s)
	prettyPrint2D("all subsets", t)
}

func generateAllSubSets(s []string) ([][]string) {
	if len(s) == 0 {
		return [][]string {}
	}
	// Generate an n-bit Gray code to determine subsets
	gray := generateGrayCode(uint(len(s)))
	subsets := make([][]string, len(gray))
	for i:=0; i<len(gray); i++ {
		for j:=len(gray[i])-1; j>=0; j-- {
			// Check if the bit is "1". If so, print the corresponding element 
			if(string(gray[i][j]) == "1") {
				ind := len(gray[i]) - j - 1
				subsets[i] = append(subsets[i], s[ind])
			}
		}
	}
	return subsets
}

func generateGrayCode(n uint) (s []string) {
    if n <= 0 {
        return
	}
	code := make([]string, 2)
	code[0], code[1] = "0", "1"

	// Every iteration of this loop generates 2*i codes from previously
	// generated i codes.
	for i:=2; i<(1<<n); i = i<<1 {
		// Enter the prviously generated codes again in arr[] in reverse
		// order. Nor arr[] has double number of codes.
		for j:=i-1; j>=0; j-- {
			code = append(code, code[j])
		}
		// append 0 to the first half
		for j:=0; j<i; j++ {
			code[j] = "0"+code[j]
		}
		// append 1 to the second half
		for j:=i; j<2*i; j++ {
			code[j] = "1"+code[j]
		}
	}
	return code
}

func prettyPrint(title string, s []string, useNewline bool) {
	if title != "" {
		out("%s = ", title)
	}
	if len(s) == 0 {
		out("[ ]")
		if useNewline {
			out("\n")
		}
	} else {
		out("[ %s", s[0])
		for i:=1;i<len(s);i++ {
			out(", %s", s[i])
		}
		out(" ]")
		if useNewline {
			out("\n")
		}
	}
}

func prettyPrint2D(title string,s [][]string) {
	if title != "" {
		out("%s = ", title)
	}
	if len(s) == 0 {
		out("[ ]\n")
	} else {
		out("[ %s", s[0])
		for i:=1;i<len(s);i++ {
			out(", ")
			prettyPrint("", s[i], false)
		}
		out(" ]\n")
	}
}

// Shorthand function for fmt.Printf
func out(format string, v...interface{}) {
	fmt.Printf(format, v...)
}