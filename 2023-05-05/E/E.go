// https://codeforces.com/problemset/problem/1559/B

// CodeForces 1559B - Mocha and Red and Blue

// As their story unravels, a timeless tale is told once again...
// Shirahime, a friend of Mocha's, is keen on playing the music game Arcaea and sharing Mocha interesting puzzles to solve. This day, Shirahime comes up with a new simple puzzle and wants Mocha to solve them. However, these puzzles are too easy for Mocha to solve, so she wants you to solve them and tell her the answers. The puzzles are described as follow.
// There are n squares arranged in a row, and each of them can be painted either red or blue.
// Among these squares, some of them have been painted already, and the others are blank. You can decide which color to paint on each blank square.
// Some pairs of adjacent squares may have the same color, which is imperfect. We define the imperfectness as the number of pairs of adjacent squares that share the same color.
// For example, the imperfectness of "BRRRBBR" is 3, with "BB" occurred once and "RR" occurred twice.
// Your goal is to minimize the imperfectness and print out the colors of the squares after painting.

// Input
// Each test contains multiple test cases.
// The first line contains a single integer t (1≤t≤100) — the number of test cases. Each test case consists of two lines.
// The first line of each test case contains an integer n (1≤n≤100) — the length of the squares row.
// The second line of each test case contains a string s with length n, containing characters 'B', 'R' and '?'. Here 'B' stands for a blue square, 'R' for a red square, and '?' for a blank square.

// Output
// For each test case, print a line with a string only containing 'B' and 'R', the colors of the squares after painting, which imperfectness is minimized. If there are multiple solutions, print any of them.

// Example
// input
// 5
// 7
// ?R???BR
// 7
// ???R???
// 1
// ?
// 1
// B
// 10
// ?R??RB??B?
// output
// BRRBRBR
// BRBRBRB
// B
// B
// BRRBRBBRBR

// Note
// In the first test case, if the squares are painted "BRRBRBR", the imperfectness is 1 (since squares 2 and 3 have the same color), which is the minimum possible imperfectness.

// Tutorial
// For a longest period of "?", it is optimized to paint either "RBRB..." or "BRBR...", so the imperfectness it made is only related to the colors of both sides of it.
// Choose the one with the lower imperfectness for each longest period of "?" in O(n) is acceptable.
// More elegantly, if the square on the left or on the right of a "?" is painted, simply paint "?" with the different color from it. This can be proved to reach the minimum imperfectness by considering the parity.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, n, cnt, i uint
	var s string
	var v []byte

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &s)
			v = []byte(s)
			cnt = 0

			for i = 0; i < n; i++ {
				if v[i] != '?' {
					cnt++
				}
			}

			if cnt == 0 {
				v[0] = 'B'
			}

			for i = 1; i < n; i++ {
				if v[i] == '?' && v[i-1] != '?' {
					v[i] = v[i-1] ^ ('B' ^ 'R')
				}
			}

			for i = n - 1; i > 0; i-- {
				if v[i-1] == '?' && v[i] != '?' {
					v[i-1] = v[i] ^ ('B' ^ 'R')
				}
			}

			fmt.Fprintln(writer, string(v))
		}
	}

	writer.Flush()
}
