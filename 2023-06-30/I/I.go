// https://codeforces.com/problemset/problem/1165/B

// CodeForces 1165B - Polycarp Training
// Polycarp wants to train before another programming competition. During the first day of his training he should solve exactly 1 problem, during the second day — exactly 2 problems, during the third day — exactly 3 problems, and so on. During the k-th day he should solve k problems.
// Polycarp has a list of n contests, the i-th contest consists of ai problems. During each day Polycarp has to choose exactly one of the contests he didn't solve yet and solve it. He solves exactly k problems from this contest. Other problems are discarded from it. If there are no contests consisting of at least k problems that Polycarp didn't solve yet during the k-th day, then Polycarp stops his training.
// How many days Polycarp can train if he chooses the contests optimally?

// Input
// The first line of the input contains one integer n (1≤n≤2⋅10^5) — the number of contests.
// The second line of the input contains n integers a1,a2,…,an (1≤ai≤2⋅10^5) — the number of problems in the i-th contest.

// Output
// Print one integer — the maximum number of days Polycarp can train if he chooses the contests optimally.

// Examples

// input
// 4
// 3 1 4 1
// output
// 3

// input
// 3
// 1 1 1
// output
// 1

// input
// 5
// 1 1 1 2 2
// output
// 2

// Tutorial
// After sorting the array, we can maintain the last day Polycarp can train, in the variable pos. Initially it is 1. Let's iterate over all elements of the sorted array in non-decreasing order and if the current element ai≥pos then let's increase pos by one. The answer will be pos−1.

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
	var t int
	var n, i, pos uint
	var a []uint

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		a = make([]uint, n)

		for i = 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}

		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

		pos = 1

		for i = 0; i < n; i++ {
			if a[i] >= pos {
				pos++
			}
		}

		fmt.Fprintln(writer, pos-1)
	}

	writer.Flush()
}
