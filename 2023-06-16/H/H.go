// https://codeforces.com/problemset/problem/1689/B

// CodeForces 1689B - Mystic Permutation

// Monocarp is a little boy who lives in Byteland and he loves programming.
// Recently, he found a permutation of length n. He has to come up with a mystic permutation. It has to be a new permutation such that it differs from the old one in each position.
// More formally, if the old permutation is p1,p2,…,pn and the new one is q1,q2,…,qn it must hold that
// p1≠q1,p2≠q2,…,pn≠qn.
// Monocarp is afraid of lexicographically large permutations. Can you please help him to find the lexicographically minimal mystic permutation?

// Input
// There are several test cases in the input data. The first line contains a single integer t (1≤t≤200) — the number of test cases. This is followed by the test cases description.
// The first line of each test case contains a positive integer n (1≤n≤1000) — the length of the permutation.
// The second line of each test case contains n distinct positive integers p1,p2,…,pn (1≤pi≤n). It's guaranteed that p is a permutation, i. e. pi≠pj for all i≠j.
// It is guaranteed that the sum of n does not exceed 1000 over all test cases.

// Output
// For each test case, output n positive integers — the lexicographically minimal mystic permutations. If such a permutation does not exist, output −1 instead.

// Example
// input
// 4
// 3
// 1 2 3
// 5
// 2 3 4 5 1
// 4
// 2 3 1 4
// 1
// 1
// output
// 2 3 1
// 1 2 3 4 5
// 1 2 4 3
// -1

// Note
// In the first test case possible permutations that are mystic are [2,3,1] and [3,1,2]. Lexicographically smaller of the two is [2,3,1].
// In the second test case, [1,2,3,4,5] is the lexicographically minimal permutation and it is also mystic.
// In third test case possible mystic permutations are [1,2,4,3], [1,4,2,3], [1,4,3,2], [3,1,4,2], [3,2,4,1], [3,4,2,1], [4,1,2,3], [4,1,3,2] and [4,3,2,1]. The smallest one is [1,2,4,3].

// Hint 1
// When is it impossible to find such permutation?

// Hint 2
// It is impossible only for n=1. In every other case, iterate over each position in order from 1to nand take the smallest available number. What's the exception to this?

// Tutorial
// The exception is the last two elements. We can always take the smallest available number for each qi satisfying i<n−1. To do this we maintain an array of bools of already taken numbers, and then iterate over it to find the smallest available number satisfying pi≠qi which is also not checked in the array, and then check it (we took it).
// Now consider (pn−1,pn) and we want (qn−1,qn) to be lexicographically minimal while satisfying pn−1≠qn−1 and pn≠qn. Let a and b be the last two unused numbers in the array of bools with a<b. We try to take (qn−1,qn)=(a,b). If a=pn−1 or b=pn, then we take (qn−1,qn)=(b,a). If (a,b) isn't valid, then (b,a) is. The proof is left as an exercise to the reader.
// This solution runs in O(n2) and can be optimized to O(n log n).

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
	var p, q []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			p = make([]uint, n)
			q = make([]uint, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &p[i])
				q[i] = i + 1
			}

			if n == 1 {
				fmt.Fprintln(writer, -1)
			} else {
				for i = 0; i+1 < n; i++ {
					if p[i] == q[i] {
						q[i], q[i+1] = q[i+1], q[i]
					}
				}

				if p[n-1] == q[n-1] {
					q[n-1], q[n-2] = q[n-2], q[n-1]
				}

				for i = 0; i+1 < n; i++ {
					fmt.Fprint(writer, q[i], " ")
				}

				fmt.Fprintln(writer, q[n-1])
			}
		}
	}

	writer.Flush()
}
