// https://codeforces.com/problemset/problem/1714/B

// CodeForces 1714B - Remove Prefix

// Polycarp was presented with some sequence of integers a of length n (1≤ai≤n). A sequence can make Polycarp happy only if it consists of different numbers (i.e. distinct numbers).
// In order to make his sequence like this, Polycarp is going to make some (possibly zero) number of moves.
// In one move, he can:
//  • remove the first (leftmost) element of the sequence.
// For example, in one move, the sequence [3,1,4,3] will produce the sequence [1,4,3], which consists of different numbers.
// Determine the minimum number of moves he needs to make so that in the remaining sequence all elements are different. In other words, find the length of the smallest prefix of the given sequence a, after removing which all values in the sequence will be unique.

// Input
// The first line of the input contains a single integer t (1≤t≤10^4) — the number of test cases.
// Each test case consists of two lines.
// The first line contains an integer n (1≤n≤2⋅10^5) — the length of the given sequence a.
// The second line contains n integers a1,a2,…,an (1≤ai≤n) — elements of the given sequence a.
// It is guaranteed that the sum of n values over all test cases does not exceed 2⋅10^5.

// Output
// For each test case print your answer on a separate line — the minimum number of elements that must be removed from the beginning of the sequence so that all remaining elements are different.

// Example
// input
// 5
// 4
// 3 1 4 3
// 5
// 1 1 1 1 1
// 1
// 1
// 6
// 6 5 4 3 2 1
// 7
// 1 2 1 7 1 2 1
// output
// 1
// 4
// 0
// 0
// 5

// Note
// The following are the sequences that will remain after the removal of prefixes:
//  • [1,4,3];
//  • [1];
//  • [1];
//  • [6,5,4,3,2,1];
//  • [2,1].
// It is easy to see that all the remaining sequences contain only distinct elements. In each test case, the shortest matching prefix was removed.

// Tutorial
// Let's turn the problem around: we'll look for the longest suffix that will make Polycarp happy, since it's the same thing.
// Let's create an array cnt, in which we will mark the numbers already encountered. Let's go along a from right to left and check if ai does not occur to the right (in this case it is marked in cnt), if it occurs to the right, then removing any prefix that does not include i, we get an array where ai occurs twice, so we have to delete prefix of length i.

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
	var t, n, i uint
	var a []uint
	var used []bool

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			a = make([]uint, n)
			used = make([]bool, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a[i])
				used[i] = false
			}

			for i = n; i > 0; i-- {
				if used[a[i-1]-1] {
					break
				}
				used[a[i-1]-1] = true
			}

			fmt.Fprintln(writer, i)
		}
	}

	writer.Flush()
}
