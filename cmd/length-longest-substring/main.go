package main


func main() {
	inputInt := "abba"
	result := lengthOfLongestSubstring(inputInt)
	_=result
}

/*
	3. Longest Substring Without Repeating Characters
	https://leetcode.com/problems/longest-substring-without-repeating-characters/
 */
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	charIDxMap := map[byte]int{}
	length := 0
	j:=0
	for i:=0; i<len(s); i++ {
		char := s[i]
		if charIDx, ok := charIDxMap[char]; ok && charIDx >=j {
			j = charIDx+1
		}

		charIDxMap[char] = i

		lenSub := i-j+1
		if lenSub > length {
			length = lenSub
		}
	}

	return length
}
