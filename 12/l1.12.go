package main

import "fmt"

func main() {
	seq := []string{"cat", "cat", "dog", "cat", "tree"}

	set := createSet(seq)

	newSeq := make([]string, 0, len(set))

	for k := range set {
		newSeq = append(newSeq, k)
	}

	fmt.Println(newSeq)
}

func createSet(seq []string) map[string]struct{} {
	set := make(map[string]struct{})

	for _, v := range seq {
		set[v] = struct{}{}
	}

	return set
}
