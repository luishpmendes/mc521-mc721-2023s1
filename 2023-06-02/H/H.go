// https://codeforces.com/problemset/problem/1559/C

// CodeForces 1559C - Mocha and Hiking
// The city where Mocha lives in is called Zhijiang. There are n+1 villages and 2n−1 directed roads in this city.
// There are two kinds of roads:
//  • n−1 roads are from village i to village i+1, for all 1≤i≤n−1.
//  • n roads can be described by a sequence a1,…,an. If ai=0, the i-th of these roads goes from village i to village n+1, otherwise it goes from village n+1 to village i, for all 1≤i≤n.
// Mocha plans to go hiking with Taki this weekend. To avoid the trip being boring, they plan to go through every village exactly once. They can start and finish at any villages. Can you help them to draw up a plan?

// Input
// Each test contains multiple test cases.
// The first line contains a single integer t (1≤t≤20) — the number of test cases. Each test case consists of two lines.
// The first line of each test case contains a single integer n (1≤n≤10^4) — indicates that the number of villages is n+1.
// The second line of each test case contains n integers a1,a2,…,an (0≤ai≤1). If ai=0, it means that there is a road from village i to village n+1. If ai=1, it means that there is a road from village n+1 to village i.
// It is guaranteed that the sum of n over all test cases does not exceed 104.

// Output
// For each test case, print a line with n+1 integers, where the i-th number is the i-th village they will go through. If the answer doesn't exist, print −1.
// If there are multiple correct answers, you can print any one of them.

// Example
// input
// 2
// 3
// 0 1 0
// 3
// 1 1 0
// output
// 1 4 2 3
// 4 1 2 3

// Note
// In the first test case, the city looks like the following graph:
// So all possible answers are (1→4→2→3), (1→2→3→4).
// In the second test case, the city looks like the following graph:
// So all possible answers are (4→1→2→3), (1→2→3→4), (3→4→1→2), (2→3→4→1).

// Tutorial
// If a1=1, then the path [(n+1)→1→2→⋯→n] is valid.
// If an=0, then the path [1→2→⋯→n→(n+1)] is valid.
// Otherwise, since a1=0∧an=1, there must exists an integer i (1≤i<n) where ai=0∧ai+1=1, then the path [1→2→⋯→i→(n+1)→(i+1)→(i+2)→⋯n] is valid.
// This is a step to prove that there always exists an Hamiltonian path in a tournament graph.

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
	var t, n, i, j uint
	var a []uint
	var flag bool

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			a = make([]uint, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a[i])
			}

			if a[0] == 1 {
				fmt.Fprint(writer, n+1, " ")

				for i = 1; i < n; i++ {
					fmt.Fprint(writer, i, " ")
				}

				fmt.Fprintln(writer, n)
			} else {
				flag = false

				for i = 0; i+1 < n && !flag; i++ {
					if a[i] == 0 && a[i+1] == 1 {
						flag = true

						for j = 1; j <= i+1; j++ {
							fmt.Fprint(writer, j, " ")
						}

						fmt.Fprint(writer, n+1, " ")

						for j = i + 2; j < n; j++ {
							fmt.Fprint(writer, j, " ")
						}

						fmt.Fprintln(writer, n)
					}
				}

				if !flag {
					for i = 1; i <= n; i++ {
						fmt.Fprint(writer, i, " ")
					}

					fmt.Fprintln(writer, n+1)
				}
			}
		}
	}

	writer.Flush()
}
