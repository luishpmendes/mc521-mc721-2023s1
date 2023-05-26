// https://codeforces.com/contest/1054/problem/B

// CodeForces 1054B - Appending Mex
// Initially Ildar has an empty array. He performs n steps. On each step he takes a subset of integers already added to the array and appends the mex of this subset to the array.
// The mex of an multiset of integers is the smallest non-negative integer not presented in the multiset. For example, the mex of the multiset [0,2,3] is 1, while the mex of the multiset [1,2,1] is 0.
// More formally, on the step m, when Ildar already has an array a1,a2,…,am−1, he chooses some subset of indices 1≤i1<i2<…<ik<m (possibly, empty), where 0≤k<m, and appends the mex(ai1,ai2,…aik) to the end of the array.
// After performing all the steps Ildar thinks that he might have made a mistake somewhere. He asks you to determine for a given array a1,a2,…,an the minimum step t such that he has definitely made a mistake on at least one of the steps 1,2,…,t, or determine that he could have obtained this array without mistakes.

// Input
// The first line contains a single integer n (1≤n≤100000) — the number of steps Ildar made.
// The second line contains n integers a1,a2,…,an (0≤ai≤10^9) — the array Ildar obtained.

// Output
// If Ildar could have chosen the subsets on each step in such a way that the resulting array is a1,a2,…,an, print −1.
// Otherwise print a single integer t — the smallest index of a step such that a mistake was made on at least one step among steps 1,2,…,t.

// Examples

// input
// 4
// 0 1 2 1
// output
// -1

// input
// 3
// 1 0 1
// output
// 1

// input
// 4
// 0 1 2 239
// output
// 4

// Note
// In the first example it is possible that Ildar made no mistakes. Here is the process he could have followed.
//  • 1-st step. The initial array is empty. He can choose an empty subset and obtain 0, because the mex of an empty set is 0. Appending this value to the end he gets the array [0].
//  • 2-nd step. The current array is [0]. He can choose a subset [0] and obtain an integer 1, because mex(0)=1. Appending this value to the end he gets the array [0,1].
//  • 3-rd step. The current array is [0,1]. He can choose a subset [0,1] and obtain an integer 2, because mex(0,1)=2. Appending this value to the end he gets the array [0,1,2].
//  • 4-th step. The current array is [0,1,2]. He can choose a subset [0] and obtain an integer 1, because mex(0)=1. Appending this value to the end he gets the array [0,1,2,1].
// Thus, he can get the array without mistakes, so the answer is −1.
// In the second example he has definitely made a mistake on the very first step, because he could not have obtained anything different from 0.
// In the third example he could have obtained [0,1,2] without mistakes, but 239 is definitely wrong.

// Tutorial
// Ildar has written an array correctly if there exists numbers 0,1,…,ai−1 to the left of ai for all i. This is equivalent to ai≤max(−1,a1,a2,…,ai−1)+1 for all i. If this condition is false for some i we made a mistake.
// So the solution is to check that ai≤max(−1,a1,a2,…,ai−1)+1 in the increasing order of i. If it is false i is answer. If it is always true answer is −1.
// Time complexity: O(n).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t, n, a, ans, m, i int

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		ans = -1
		m = 0

		for i = 0; i < n && ans < 0; i++ {
			fmt.Fscan(reader, &a)

			if a == m {
				m++
			} else if a > m {
				ans = i + 1
			}
		}

		for ; i < n; i++ {
			fmt.Fscan(reader, &a)
		}

		fmt.Fprintln(writer, ans)
	}

	writer.Flush()
}
