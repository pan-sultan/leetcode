/*
Description:
Given a string s, return the longest palindromic substring in s.

https://leetcode.com/problems/longest-palindromic-substring/

Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Palindromic Substring.
Memory Usage: 3.6 MB, less than 34.09% of Go online submissions for Longest Palindromic Substring.

solved with Manacher's Algorithm
*/
package main

import (
	"log"
)

func main() {
	leftTests := map[string]int{
		"abababababa":  11,
		"xxyvavavyazt": 7,
		"a":            1,
		"":             0,
		"cbabcdcbaxy":  7,
	}

	rightTests := map[string]int{
		"aabbaabbaabbaa": 14,
		"yxaaxy":         6,
		"aa":             2,
		"":               0,
		"cbabcddcbaxy":   8,
		"accaaccaaa":     8,
	}

	for palindrom, expectedMax := range leftTests {
		str := longestPalindrome(palindrom)
		if len(str) != expectedMax {
			log.Fatalf("'%s': %d != %d\n", palindrom, len(str), expectedMax)
		}
	}

	for palindrom, expectedMax := range rightTests {
		str := longestPalindrome(palindrom)
		if len(str) != expectedMax {
			log.Fatalf("'%s': %d != %d\n", palindrom, len(str), expectedMax)
		}
	}
}

// Manacher's algorithm
func longestPalindrome(s string) string {

	// even palindrome length
	longestPalindromeEven := func(s string) <-chan string {
		ch := make(chan string)
		go func() {
			defer close(ch)
			palinLens := make([]int, len(s))
			var i, currRight, palinLen, palinEnd, right, left, curLen, max, maxLeft, maxRight int
			right = -1

			for i = 1; i < len(s); i++ {
				if i < right {
					palinLen = palinLens[left+right-i+1]
					palinEnd = right - i
					if palinLen < palinEnd {
						curLen = palinLen
					} else {
						curLen = palinEnd
					}
				} else {
					curLen = 0
				}

				for i+curLen < len(s) && i-curLen-1 >= 0 && s[i-curLen-1] == s[i+curLen] {
					curLen++
				}

				palinLens[i] = curLen
				currRight = i + curLen - 1
				if currRight > right {
					right = currRight
					left = i - curLen
				}

				curLen = right - left + 1
				if max < curLen {
					max = curLen
					maxLeft = left
					maxRight = right
				}
			}

			if max == 0 {
				return
			}

			ch <- s[maxLeft : maxRight+1]
			return
		}()
		return ch
	}

	// odd palindrome length
	longestPalindromeOdd := func(s string) <-chan string {
		ch := make(chan string)
		go func() {
			defer close(ch)
			palinLens := make([]int, len(s))
			var i, currRight, palinLen, palinEnd, right, left, curLen, max, maxRight, maxLeft int
			right = -1

			for i = 0; i < len(s); i++ {
				if i < right {
					palinLen = palinLens[left+right-i]
					palinEnd = right - i
					if palinLen < palinEnd {
						curLen = palinLen
					} else {
						curLen = palinEnd
					}
				} else {
					curLen = 1
				}

				for ; i+curLen < len(s) && i-curLen >= 0 && s[i-curLen] == s[i+curLen]; curLen++ {
				}

				palinLens[i] = curLen
				currRight = i + curLen - 1
				if currRight > right {
					right = currRight
					left = i - curLen + 1
				}

				curLen = right - left + 1
				if max < curLen {
					max = curLen
					maxRight = right
					maxLeft = left
				}
			}

			if max == 0 {
				return
			}

			ch <- s[maxLeft : maxRight+1]
		}()
		return ch
	}

	palindromEvent := <-longestPalindromeEven(s)
	palindromOdd := <-longestPalindromeOdd(s)

	if len(palindromEvent) < len(palindromOdd) {
		return palindromOdd
	}
	return palindromEvent
}
