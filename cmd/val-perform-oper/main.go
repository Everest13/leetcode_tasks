package main

func main() {

}

/*
	2011. Final Value of Variable After Performing Operations
	Easy
	https://leetcode.com/problems/final-value-of-variable-after-performing-operations/
	Runtime 0ms 100%, Memory 3.38MB 18.94%
 */
func finalValueAfterOperations(operations []string) int {
	x := 0

	opFuncs := map[string]func(){
		"++X": func() { x = x + 1 },
		"X++": func() { x = x + 1 },
		"--X": func() { x = x - 1 },
		"X--": func() { x = x - 1 },
	}

	for i := 0; i < len(operations); i++ {
		f := opFuncs[operations[i]]
		f()

	}

	return x
}
