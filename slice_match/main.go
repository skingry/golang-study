package main

import "fmt"
import "slices"
import "strings"

func main() {
	// All the various inputs, both examples and others...

	//S := []int{0, 0, 0, 0, 0, 0, 2, 0, 0}
	//S := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}

	// TODO: Edge case... output on this sequence is "1"... should probably be "7"
	//S := []int{7, 6, 5, 4, 3, 2, 1}

	//S := []int{1, 2, 5, 2, 1, 1, 4, 6, 9, 0}

	// TODO: Edge case... output on this sequence is "0"... should probably be "7"
	//S := []int{7, 6, 5, 4, 3, 2, 1, 0}

	// TODO: Not entirely sure about this one... if the first four elements, or the last six, best fit the description of "longest ascending"... technically, the last six aren't "ascending", just constant... but they are a longer sequence
	S := []int{1, 1, 2, 7, 2, 2, 2, 2, 2, 2}

	ss := []int{}
	var n int
	m := make(map[int]string)

	// All the logic to iterate over the sequence and find subsequences that are ascending... storing them in a map to be sorted later
	for i := 0; i < len(S); i++ {
		n = i + 1
		if n != len(S) {
			if S[i] <= S[n] {
				ss = append(ss, S[i])
			} else if S[n] <= S[i]{
				ss = append(ss, S[i])
				subString := arrayToString(ss, ", ")
				m[len(ss)] = subString
				ss = nil
			}
		} else {
			ss = append(ss, S[i])
			subString := arrayToString(ss, ", ")
			m[len(ss)] = subString
		}
	}

	// Extract keys from map, to do a sort finding max later
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
  }

	// Output the longest subsequence
	fmt.Println(m[slices.Max(keys)])
}

// Copy/Pasta'd this function from:
// https://tinyurl.com/3en9u8jy
func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
