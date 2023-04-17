// https://codeforces.com/problemset/problem/1056/A

// CodeForces 1056A - Determine Line

// Arkady's morning seemed to be straight of his nightmare. He overslept through the whole morning and, still half-asleep, got into the tram that arrived the first. Some time after, leaving the tram, he realized that he was not sure about the line number of the tram he was in.
// During his ride, Arkady woke up several times and each time he saw the tram stopping at some stop. For each stop he knows which lines of tram stop there. Given this information, can you help Arkady determine what are the possible lines of the tram he was in?

// Input
// The first line contains a single integer n (2≤n≤100) — the number of stops Arkady saw.
// The next n lines describe the stops. Each of them starts with a single integer r (1≤r≤100) — the number of tram lines that stop there. r distinct integers follow, each one between 1 and 100, inclusive, — the line numbers. They can be in arbitrary order.
// It is guaranteed that Arkady's information is consistent, i.e. there is at least one tram line that Arkady could take.

// Output
// Print all tram lines that Arkady could be in, in arbitrary order.

// Examples

// input
// 3
// 3 1 4 6
// 2 1 4
// 5 10 5 6 4 1
// output
// 1 4

// input
// 5
// 1 1
// 10 10 9 8 7 100 5 4 3 99 1
// 5 1 2 3 4 5
// 5 4 1 3 2 5
// 4 10 1 5 3
// output
// 1

// Note
// Consider the first example. Arkady woke up three times. The first time he saw a stop with lines 1, 4, 6. The second time he saw a stop with lines 1, 4. The third time he saw a stop with lines 10, 5, 6, 4 and 1. He can be in a tram of one of two lines: 1 or 4.

// Tutorial
// This is a simple implementation problem. A tram line is suitable if it appears at all stops, i.e. exactly n times. Make an array of 100 integers and count how many times each integer appears. Then just output each index where the array hits n.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var n, r, i, j, k uint
	var v [100]uint

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		v = [100]uint{}

		for i = 0; i < n; i++ {
			fmt.Fscan(reader, &r)

			for j = 0; j < r; j++ {
				fmt.Fscan(reader, &k)
				v[k-1]++
			}
		}

		for i = 0; i < 100; i++ {
			if v[i] == n {
				fmt.Fprint(writer, i+1, " ")
			}
		}

		fmt.Fprintln(writer)
	}

	writer.Flush()
}
