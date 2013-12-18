package main

import "fmt"

func main() {
	fmt.Printf("Starting subset generator\n")
	s := []string { "1", "2", "3" }
	fmt.Printf("Generating 3 bit Gray codes\n")
	gray := generateGrayCode(3)
	prettyPrint(gray)
	fmt.Printf("Generating all possible sub sets for: s = %s\n", s)
	t := generateAllSubSets(s)
	prettyPrint(t)
}

func generateAllSubSets(s []string) ([]string) {
	if len(s) == 0 {
		return s
	}
	gray := generateGrayCode(uint(len(s)))
	subsets := make([]string, len(gray))
	for i:=0; i<len(gray); i++ {
		printed := false
		for j:=0; j<len(gray[i]); j++ {
			if(string(gray[i][j]) == "1") {
				ind := len(gray[i]) - j - 1
				if(j > 0 && printed) {
					subsets[i] = subsets[i] + "," + s[ind]
				} else {
					subsets[i] = subsets[i] + s[ind]
				}
				printed = true
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

func prettyPrint(s []string) {
	if(len(s) == 0) {
		fmt.Printf("[ ]")
	} else {
		fmt.Printf("[ {%s}", s[0])
		for i:=1;i<len(s);i++ {
			fmt.Printf(", {%s}", s[i])
		}
		fmt.Printf(" ]\n")
	}
}