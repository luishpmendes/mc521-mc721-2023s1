// https://codeforces.com/problemset/problem/1277/D

// CodeForces 1277D - Let's Play the Words?

// Polycarp has n different binary words. A word called binary if it contains only characters '0' and '1'. For example, these words are binary: "0001", "11", "0" and "0011100".
// Polycarp wants to offer his set of n binary words to play a game "words". In this game, players name words and each next word (starting from the second) must start with the last character of the previous word. The first word can be any. For example, these sequence of words can be named during the game: "0101", "1", "10", "00", "00001".
// Word reversal is the operation of reversing the order of the characters. For example, the word "0111" after the reversal becomes "1110", the word "11010" after the reversal becomes "01011".
// Probably, Polycarp has such a set of words that there is no way to put them in the order correspondent to the game rules. In this situation, he wants to reverse some words from his set so that:
//  • the final set of n words still contains different words (i.e. all words are unique);
//  • there is a way to put all words of the final set of words in the order so that the final sequence of n words is consistent with the game rules.
// Polycarp wants to reverse minimal number of words. Please, help him.

// Input
// The first line of the input contains one integer t (1≤t≤10^4) — the number of test cases in the input. Then t test cases follow.
// The first line of a test case contains one integer n (1≤n≤2⋅10^5) — the number of words in the Polycarp's set. Next n lines contain these words. All of n words aren't empty and contains only characters '0' and '1'. The sum of word lengths doesn't exceed 4⋅106. All words are different.
// Guaranteed, that the sum of n for all test cases in the input doesn't exceed 2⋅105. Also, guaranteed that the sum of word lengths for all test cases in the input doesn't exceed 4⋅10^6.

// Output
// Print answer for all of t test cases in the order they appear.
// If there is no answer for the test case, print -1. Otherwise, the first line of the output should contain k (0≤k≤n) — the minimal number of words in the set which should be reversed. The second line of the output should contain k distinct integers — the indexes of the words in the set which should be reversed. Words are numerated from 1 to n in the order they appear. If k=0 you can skip this line (or you can print an empty line). If there are many answers you can print any of them.

// Example
// input
// 4
// 4
// 0001
// 1000
// 0011
// 0111
// 3
// 010
// 101
// 0
// 2
// 00000
// 00001
// 4
// 01
// 001
// 0001
// 00001
// output
// 1
// 3
// -1
// 0
//
// 2
// 1 2

// Tutorial
// For a concrete set of words, it's not hard to find a criteria for checking if there is a correct order of arrangement of words for playing a game. Let's call such sets of words correct. Firstly the set of words is correct if the number of words like 0...1 and the number of words like 1...0 differ by no more than 1. Secondly it's correct if the number of words like 0...0 or like 1...1 is zero, because they have the same characters at the beginning and at the ending, and we can insert them in any position. And finally if words of both kinds 0...0 and 1...1 are present and there is at least one word like 0...1 or 1...0.
// It can be easily proved if we note that this problem is equivalent to the Euler traversal of a directed graph with two nodes. But let's prove it without resorting to graph theory:
//  • if there are words of both kinds 0...0 and 1...1, but there is no words of kinds 0...1 and 1...0, starting from a word of one kind you can't go to a word of another kind. Consequently, if words of both kinds 0...0 and 1...1 are present, there should be at least one word like 0...1 or 1...0 — is a necessary condition of the problem;
//  • if the number of words like 0...1 and the number of words like 1...0 differ by no more than 1, we can call them alternately starting with a kind that is larger. If these numbers are equal, we can start with any kind. And we can insert words of kind 0...0 and 1...1 at any suitable moment.
// Reversals only affect the mutual number of lines of the kind 0...1 and 1...0. Therefore, immediately while reading the input data, we can check the necessary condition (first item above).
// Without loss of generality we may assume that the number of words like 0...1 equals n01 and like 1...0 equals n10. Also we assume that n01>n10+1. Remember that all words in the current set are unique. Let's prove that we can always choose some words of kind 0...1 and reverse them so that n01=n10+1 (and at the result all words would still be unique).
// In fact, the set of words of kind n10 has no more than n10 such words that after the reversing, the word will turn into an existing one (because it will become of type 1...0 and there are only n10 such words). And it means that there is no less than n01−n10 words which we can reverse and get still unique word. So, we can choose any n01−n10−1 of them.
// Thus, after checking of the necessary condition (first item above), we need to reverse just n01−n10−1 words of kind that is larger, which reversals aren't duplicates.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, n, i, ans uint
	var s []string
	var s01, s10 map[string]struct{}
	var u []bool
	var rev []uint
	var ss string

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)

			s = make([]string, n)
			s01 = make(map[string]struct{})
			s10 = make(map[string]struct{})
			u = make([]bool, 2)

			u[0] = false
			u[1] = false

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &s[i])

				if s[i][0] == '0' && s[i][len(s[i])-1] == '1' {
					s01[s[i]] = struct{}{}
				}

				if s[i][0] == '1' && s[i][len(s[i])-1] == '0' {
					s10[s[i]] = struct{}{}
				}

				if s[i][0] == '0' {
					u[0] = true
				} else {
					u[1] = true
				}

				if s[i][len(s[i])-1] == '0' {
					u[0] = true
				} else {
					u[1] = true
				}
			}

			if u[0] && u[1] && len(s01) == 0 && len(s10) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				rev = make([]uint, 0, n)

				if len(s01) > len(s10)+1 {
					for i = 0; i < n; i++ {
						if s[i][0] == '0' && s[i][len(s[i])-1] == '1' {
							ss = reverse(s[i])

							if _, ok := s10[ss]; !ok {
								rev = append(rev, i)
							}
						}
					}
				} else if len(s10) > len(s01)+1 {
					for i = 0; i < n; i++ {
						if s[i][0] == '1' && s[i][len(s[i])-1] == '0' {
							ss = reverse(s[i])

							if _, ok := s01[ss]; !ok {
								rev = append(rev, i)
							}
						}
					}
				}

				if len(s01) > len(s10) {
					ans = uint(len(s01)-len(s10)) / 2
				} else {
					ans = uint(len(s10)-len(s01)) / 2
				}

				fmt.Fprintln(writer, ans)

				for i = 0; i < ans; i++ {
					fmt.Fprint(writer, rev[i]+1, " ")
				}

				fmt.Fprintln(writer)
			}
		}
	}

	writer.Flush()
}
