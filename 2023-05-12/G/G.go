// https://codeforces.com/problemset/problem/1566/C

// CodeForces 1566C - MAX-MEX Cut

// A binary string is a string that consists of characters 0 and 1. A bi-table is a table that has exactly two rows of equal length, each being a binary string.
// Let MEX of a bi-table be the smallest digit among 0, 1, or 2 that does not occur in the bi-table. For example, MEX for [0011,1010] is 2 , because 0 and 1 occur in the bi-table at least once. MEX for [111,111] is 0, because 0 and 2 do not occur in the bi-table, and 0<2.
// You are given a bi-table with ncolumns. You should cut it into any number of bi-tables (each consisting of consecutive columns) so that each column is in exactly one bi-table. It is possible to cut the bi-table into a single bi-table — the whole bi-table.
// What is the maximal sum of MEXof all resulting bi-tables can be?

// Input
// The input consists of multiple test cases. The first line contains a single integer t (1≤t≤10^4) — the number of test cases. Description of the test cases follows.
// The first line of the description of each test case contains a single integer n (1≤n≤10^5) — the number of columns in the bi-table.
// Each of the next two lines contains a binary string of length n — the rows of the bi-table.
// It's guaranteed that the sum of n over all test cases does not exceed 10^5.

// Output
// For each test case print a single integer — the maximal sum of MEX of all bi-tables that it is possible to get by cutting the given bi-table optimally.

// Example
// input
// 4
// 7
// 0101000
// 1101100
// 5
// 01100
// 10101
// 2
// 01
// 01
// 6
// 000000
// 111111
// output
// 8
// 8
// 2
// 12

// Note
// In the first test case you can cut the bi-table as follows:
//  • [0,1], its MEX is 2.
//  • [10,10], its MEX is 2.
//  • [1,1], its MEX is 0.
//  • [0,1], its MEX is 2.
//  • [0,0], its MEX is 1.
//  • [0,0], its MEX is 1.
// The sum of MEX is 8.

// Tutorial
// Another possible solution with dp is based on the fact that we should not take any segments with length more than x, where x is some small number. We can just take some random big enough x and not prove anything. There exists a solution which does not consider segments with length bigger than 5.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mex(s string) uint {
	var fl, result uint
	fl = 0
	result = 0

	for _, nx := range s {
		if nx == '0' {
			fl |= 1
		} else if nx == '1' {
			fl |= 2
		}
	}

	if fl == 3 {
		result = 2
	} else if fl == 1 {
		result = 1
	}

	return result
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, n, i, j, aux uint
	var s1, s2 string
	var s strings.Builder
	var dp []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &s1, &s2)
			dp = make([]uint, n+1)

			for i = 0; i < n; i++ {
				s = strings.Builder{}

				for j = 0; j < 5 && i+j < n; j++ {
					s.WriteByte(s1[i+j])
					s.WriteByte(s2[i+j])
					aux = dp[i] + mex(s.String())

					if dp[i+j+1] < aux {
						dp[i+j+1] = aux
					}
				}
			}

			fmt.Fprintln(writer, dp[n])
		}
	}

	writer.Flush()
}
