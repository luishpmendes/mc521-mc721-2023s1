// https://codeforces.com/problemset/problem/1408/F

// CodeForces 1408F - Two Different

// You are given an integer n.
// You should find a list of pairs (x1,y1), (x2,y2), ..., (xq,yq) (1≤xi,yi≤n) satisfying the following condition.
// Let's consider some function f:N×N→N (we define N as the set of positive integers). In other words, f is a function that returns a positive integer for a pair of positive integers.
// Let's make an array a1,a2,…,an, where ai=i initially.
// You will perform q operations, in i-th of them you will:
//  1. assign t=f(axi,ayi) (t is a temporary variable, it is used only for the next two assignments);
//  2. assign axi=t;
//  3. assign ayi=t.
// In other words, you need to simultaneously change axi and ayi to f(axi,ayi). Note that during this process f(p,q) is always the same for a fixed pair of p and q.
// In the end, there should be at most two different numbers in the array a.
// It should be true for any function f.
// Find any possible list of pairs. The number of pairs should not exceed 5⋅10^5.

// Input
// The single line contains a single integer n (1≤n≤15000).

// Output
// In the first line print q (0≤q≤5⋅10^5) — the number of pairs.
// In each of the next q lines print two integers. In the i-th line print xi, yi (1≤xi,yi≤n).
// The condition described in the statement should be satisfied.
// If there exists multiple answers you can print any of them.

// Examples

// input
// 3
// output
// 1
// 1 2

// input
// 4
// output
// 2
// 1 2
// 3 4

// Note
// In the first example, after performing the only operation the array a will be [f(a1,a2),f(a1,a2),a3]. It will always have at most two different numbers.
// In the second example, after performing two operations the array a will be [f(a1,a2),f(a1,a2),f(a3,a4),f(a3,a4)]. It will always have at most two different numbers.

// Tutorial
// Claim: for each k, we can perform operations on 2^k elements to make all numbers the same.
// You can prove this fact with induction by k. For k=0 this fact is obvious, For k>0, at first change first 2^(k−1) and last 2^(k−1) points to the same value (assume that those values are x and y, respectively). And then, perform operations on points i and i+2^(k−1), to simultaneously change them to f(x,y). After that, all values will be equal ■.
// Using this observation, it is easy to leave only two different values in the array. Find the maximum 2^k≤n, and then change first 2^k numbers to the same number, and then change last 2^k elements to the same number (note that 2^k+2^k>n, so in the end there will be only two different elements).

package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	a, b uint
}

func solve(ans *[]pair, l, r uint) {
	var mid, i uint

	if l+1 < r {
		mid = (l + r) / 2

		for i = 0; i+l < mid; i++ {
			*ans = append(*ans, pair{l + i, mid + i})
		}

		solve(ans, l, mid)
		solve(ans, mid, r)

	}
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var n, k uint
	var ans []pair

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		ans = make([]pair, 0, 500000)

		k = 1

		for (k << 1) < n {
			k <<= 1
		}

		solve(&ans, 0, k)
		solve(&ans, n-k, n)

		fmt.Fprintln(writer, len(ans))

		for _, p := range ans {
			fmt.Fprintln(writer, p.a+1, p.b+1)
		}
	}

	writer.Flush()
}
