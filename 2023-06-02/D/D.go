// https://codeforces.com/problemset/problem/1806/B

// CodeForces 1806B - Mex Master

// You are given an array a of length n. The score of a is the MEX of [a1+a2,a2+a3,…,an−1+an]. Find the minimum score of a if you are allowed to rearrange elements of a in any order. Note that you are not required to construct the array a that achieves the minimum score.
// The MEX (minimum excluded) of an array is the smallest non-negative integer that does not belong to the array. For instance:
//  • The MEX of [2,2,1] is 0, because 0 does not belong to the array.
//  • The MEX of [3,1,0,1] is 2, because 0 and 1 belong to the array, but 2 does not.
//  • The MEX of [0,3,1,2] is 4 because 0, 1, 2, and 3 belong to the array, but 4 does not.

// Input
// The first line contains a single integer t (1≤t≤10^4) — the number of test cases. The description of test cases follows.
// The first line of each test case contains a single integer n (2≤n≤2⋅10^5).
// The second line of each test case contains n integers a1,a2,…,an (0≤ai≤2⋅10^5).
// It is guaranteed that the sum of n over all test cases does not exceed 2⋅105.

// Output
// For each test case, output the minimum score of a after rearranging the elements of a in any order.

// Example
// input
// 3
// 2
// 0 0
// 3
// 0 0 1
// 8
// 1 0 0 0 2 0 3 0
// output
// 1
// 0
// 1

// Note
// In the first test case, it is optimal to rearrange a as [0,0], the score of this array is the MEX of [0+0]=[0], which is 1.
// In the second test case, it is optimal to rearrange a as [0,1,0], the score of this array is the MEX of [0+1,1+0]=[1,1], which is 0.

// Tutorial
// Hint 1: ans≤2.
// First, let's determine if ans can be 0. That means we can't place two 0s next to each other. This is achievable when the number of 0s is not greater than ⌈n/2⌉.
// Then determine if ans can be 1. That means we can't place 0 and 1 next to each other. Therefore, if there is no 1 in a or there exist an element x≥2 in a, we can simply rearrange a as [0,0,…,0,x,1,1,…] to make ans=1.
// The last case: there are only 0 and 1 in a and the number of 0s is greater than ⌈n2⌉. We want to make ans=2 which is the minimum. Since the number of 1s is not greater than the number of 0s, we can rearrange a as [0,1,0,1,…,0,1,0,0,…] to make ans=2.

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
	var t, n, a, i, sum0 uint
	var f bool

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			sum0 = 0
			f = false

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a)

				if a == 0 {
					sum0++
				} else if a > 1 {
					f = true
				}
			}

			if sum0<<1 <= n+1 {
				fmt.Fprintln(writer, 0)
			} else if f || sum0 == n {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 2)
			}
		}
	}

	writer.Flush()
}
