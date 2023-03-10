// https://codeforces.com/problemset/problem/1760/C

// CodeForces 1760C - Advantage

// There are n participants in a competition, participant i having a strength of si.
// Every participant wonders how much of an advantage they have over the other best participant. In other words, each participant i wants to know the difference between si and sj, where j is the strongest participant in the competition, not counting i (a difference can be negative).
// So, they ask you for your help! For each i (1≤i≤n) output the difference between si and the maximum strength of any participant other than participant i.

// Input
// The input consists of multiple test cases. The first line contains an integer t (1≤t≤1000) — the number of test cases. The descriptions of the test cases follow.
// The first line of each test case contains an integer n (2≤n≤2⋅10^5) — the length of the array.
// The following line contains n space-separated positive integers s1, s2, ..., sn (1≤si≤10^9) — the strengths of the participants.
// It is guaranteed that the sum of n over all test cases does not exceed 2⋅10^5.

// Output
// For each test case, output n space-separated integers. For each i (1≤i≤n) output the difference between si and the maximum strength of any other participant.

// Example
// input
// 5
// 4
// 4 7 3 5
// 2
// 1 2
// 5
// 1 2 3 4 5
// 3
// 4 9 4
// 4
// 4 4 4 4
// output
// -3 2 -4 -2
// -1 1
// -4 -3 -2 -1 1
// -5 5 -5
// 0 0 0 0

// Note
// For the first test case:
//  • The first participant has a strength of 4 and the largest strength of a participant different from the first one is 7, so the answer for the first participant is 4−7=−3.
//  • The second participant has a strength of 7 and the largest strength of a participant different from the second one is 5, so the answer for the second participant is 7−5=2.
//  • The third participant has a strength of 3 and the largest strength of a participant different from the third one is 7, so the answer for the third participant is 3−7=−4.
//  • The fourth participant has a strength of 5 and the largest strength of a participant different from the fourth one is 7, so the answer for the fourth participant is 5−7=−2.

// Tutorial

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, t, n, i int
	var s, v []int

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			s = make([]int, n)
			v = make([]int, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &s[i])
				v[i] = s[i]
			}

			sort.Ints(v)

			for i = 0; i < n; i++ {
				if s[i] != v[n-1] {
					fmt.Fprint(writer, s[i]-v[n-1], " ")
				} else {
					fmt.Fprint(writer, s[i]-v[n-2], " ")
				}
			}

			fmt.Fprintln(writer)
		}
	}

	writer.Flush()
}