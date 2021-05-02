/*
Description:
Given a string s, find the length of the longest substring without repeating characters.

https://leetcode.com/problems/longest-substring-without-repeating-characters/submissions/

It seems to be the fastest solution
Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Substring Without Repeating Characters.
Memory Usage: 2.6 MB, less than 99.94% of Go online submissions for Longest Substring Without Repeating Characters.

*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		start := time.Now()
		max := lengthOfLongestSubstring(str)
		duration := time.Since(start)
		fmt.Printf("length of longest substring: %d\n", max)
		fmt.Printf("%s: runtime: %v\n", str, duration)
	}
}

func lengthOfLongestSubstring(s string) int {
	var c, j, i, k, max, count int
	hash := make([]int, 127)

	for i = 0; i < len(s); {
		c = int(s[i])
		j = hash[c]
		if j >= k {
			count = i - k
			if count > max {
				max = count
			}
			k = j
		}
		i++
		hash[c] = i
	}

	count = i - k
	if count > max {
		max = count
	}

	return max
}
