// https://codeforces.com/problemset/problem/1353/B

// CodeForces 1353B - Two Arrays And Swaps

// You are given two arrays a and b both consisting of n positive (greater than zero) integers. You are also given an integer k.
// In one move, you can choose two indices i and j (1≤i,j≤n) and swap ai and bj (i.e. ai becomes bj and vice versa). Note that i and j can be equal or different (in particular, swap a2 with b2 or swap a3 and b9 both are acceptable moves).
// Your task is to find the maximum possible sum you can obtain in the array a if you can do no more than (i.e. at most) k such moves (swaps).
// You have to answer t independent test cases.

// Input
// The first line of the input contains one integer t (1≤t≤200) — the number of test cases. Then t test cases follow.
// The first line of the test case contains two integers n and k (1≤n≤30;0≤k≤n) — the number of elements in a and b and the maximum number of moves you can do. The second line of the test case contains n integers a1,a2,…,an (1≤ai≤30), where ai is the i-th element of a. The third line of the test case contains n integers b1,b2,…,bn (1≤bi≤30), where bi is the i-th element of b.

// Output
// For each test case, print the answer — the maximum possible sum you can obtain in the array a if you can do no more than (i.e. at most) k swaps.

// Example
// input
// 5
// 2 1
// 1 2
// 3 4
// 5 5
// 5 5 6 6 5
// 1 2 5 4 3
// 5 3
// 1 2 3 4 5
// 10 9 10 10 9
// 4 0
// 2 2 4 3
// 2 4 2 3
// 4 4
// 1 2 2 1
// 4 4 5 4
// output
// 6
// 27
// 39
// 11
// 17

// Note
// In the first test case of the example, you can swap a1=1 and b2=4, so a=[4,2] and b=[3,1].
// In the second test case of the example, you don't need to swap anything.
// In the third test case of the example, you can swap a1=1 and b1=10, a3=3 and b3=10 and a2=2 and b4=10, so a=[10,10,10,4,5] and b=[1,9,3,2,9].
// In the fourth test case of the example, you cannot swap anything.
// In the fifth test case of the example, you can swap arrays a and b, so a=[4,4,5,4] and b=[1,2,2,1].

// Tutorial
// Each move we can choose the minimum element in a, the maximum element in b and swap them (if the minimum in a is less than maximum in b). If we repeat this operation k times, we get the answer. This can be done in O(n3), O(n2) but authors solution is O(nlogn).

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
	var c int
	var t, n, k, i, ans uint
	var a, b []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &k)
			a = make([]uint, n)
			b = make([]uint, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a[i])
			}

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &b[i])
			}

			sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
			sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
			ans = 0

			for i = 0; i < n; i++ {
				if i < k {
					if a[i] < b[i] {
						ans += b[i]
					} else {
						ans += a[i]
					}
				} else {
					ans += a[i]
				}
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
