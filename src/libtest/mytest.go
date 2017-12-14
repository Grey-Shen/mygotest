package main

import "fmt"

func main() {
	// reader := strings.NewReader("this is just a test")
	// fmt.Println("len--", reader.Len())
	// reader.Reset("中国")
	// // fmt.Println("reader"s)
	// a, b, _ := reader.ReadRune()
	// fmt.Println("a is ", a, "b is ", b)
	// n, _ := reader.WriteTo(os.Stdout)
	// fmt.Println(" ")
	// fmt.Println("n is ", n)
	fmt.Println(lengthOfLongestSubstring("aab"))
}

// ccdeff

// subStringLen: 1
// subStringMax: 0
// subStringLen: 2
// subStringMax: 0
// subStringLen: 1
// subStringMax: 2
// subStringLen: 2
// subStringMax: 2

func longestSubstring(s string, k int) int {
	if len(s) <= 0 || k > len(s) {
		return 0
	}

	if k == 0 {
		return len(s)
	}

	record := make(map[string]int)
	for _, r := range s {
		if _, ok := record[string(r)]; !ok {
			record[string(r)] = 1
		} else {
			record[string(r)]++
		}
	}

	index := 0
	// var value rune
	// for index, value = range s {
	// 	if record[string(value)] < k {
	// 		break
	// 	}
	// }

	for index < len(s) && record[string(s[index])] >= k {
		index++
	}

	if index == len(s) {
		return len(s)
	}

	left := longestSubstring(s[:index], k)
	right := longestSubstring(s[index+1:], k)

	if left > right {
		return left
	} else {
		return right
	}

}

// aab
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	var (
		longest int
		i       int
		j       int
	)

	for i = 0; i < len(s); i++ {
		record := make(map[byte]int)
		record[s[i]] = 1
		for j = i + 1; j < len(s); j++ {
			if _, ok := record[s[j]]; !ok {
				record[s[j]] = 1
			} else {
				cLen := j - i
				if cLen > longest {
					longest = cLen
				}
				break
			}

		}

		if j == len(s) && j-i > longest {
			longest = j - i
		}
	}

	return longest
}
