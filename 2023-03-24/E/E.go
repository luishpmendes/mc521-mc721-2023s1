// https://codeforces.com/problemset/problem/1722/C

// CodeForces 1722C - Word Game

// Three guys play a game: first, each person writes down n distinct words of length 3. Then, they total up the number of points as follows:
//  • if a word was written by one person — that person gets 3 points,
//  • if a word was written by two people — each of the two gets 1 point,
//  • if a word was written by all — nobody gets any points.
// In the end, how many points does each player have?

// Input
// The input consists of multiple test cases. The first line contains an integer t (1≤t≤100) — the number of test cases. The description of the test cases follows.
// The first line of each test case contains an integer n (1≤n≤1000) — the number of words written by each person.
// The following three lines each contain n distinct strings — the words written by each person. Each string consists of 3 lowercase English characters.

// Output
// For each test case, output three space-separated integers — the number of points each of the three guys earned. You should output the answers in the same order as the input; the i-th integer should be the number of points earned by the i-th guy.

// Example
// input
// 3
// 1
// abc
// def
// abc
// 3
// orz for qaq
// qaq orz for
// cod for ces
// 5
// iat roc hem ica lly
// bac ter iol ogi sts
// bac roc lly iol iat
// output
// 1 3 1
// 2 2 6
// 9 11 5

// Note
// In the first test case:
//  • The word abc was written by the first and third guys — they each get 1 point.
//  • The word def was written by the second guy only — he gets 3 points.

// Tutorial
// You need to implement what is written in the statement. To quickly check if a word is written by another guy, you should store some map<string, int> or Python dictionary, and increment every time you see a new string in the input. Then, you should iterate through each guy, find the number of times their word appears, and update their score.
// The complexity is O(nlogn) per testcase.

package main

import (
	"bufio"
	"fmt"
	"os"
)

type set map[string]struct{}

func (s set) add(word string) {
	s[word] = struct{}{}
}

func (s set) has(word string) bool {
	_, ok := s[word]
	return ok
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, n, i, points uint
	var word string
	var written_by_A, written_by_B, written_by_C set
	var count map[string]uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			written_by_A = make(set)
			written_by_B = make(set)
			written_by_C = make(set)
			count = make(map[string]uint)

			fmt.Fscan(reader, &n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &word)
				written_by_A.add(word)
				count[word]++
			}

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &word)
				written_by_B.add(word)
				count[word]++
			}

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &word)
				written_by_C.add(word)
				count[word]++
			}

			points = 0

			for word = range written_by_A {
				if count[word] == 1 {
					points += 3
				} else if count[word] == 2 {
					points++
				}
			}

			fmt.Fprint(writer, points, " ")

			points = 0

			for word = range written_by_B {
				if count[word] == 1 {
					points += 3
				} else if count[word] == 2 {
					points++
				}
			}

			fmt.Fprint(writer, points, " ")

			points = 0

			for word = range written_by_C {
				if count[word] == 1 {
					points += 3
				} else if count[word] == 2 {
					points++
				}
			}

			fmt.Fprintln(writer, points)
		}
	}

	writer.Flush()
}
