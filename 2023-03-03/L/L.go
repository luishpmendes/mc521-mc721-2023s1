// https://codeforces.com/problemset/problem/1438/C

// CodeForces 1438C - Engineer Artem

// Artem is building a new robot. He has a matrix a consisting of n rows and m columns. The cell located on the i-th row from the top and the j-th column from the left has a value ai,j written in it.
// If two adjacent cells contain the same value, the robot will break. A matrix is called good if no two adjacent cells contain the same value, where two cells are called adjacent if they share a side.
// Artem wants to increment the values in some cells by one to make a good.
// More formally, find a good matrix b that satisfies the following condition —
//  • For all valid (i,j), either bi,j=ai,j or bi,j=ai,j+1.
// For the constraints of this problem, it can be shown that such a matrix b always exists. If there are several such tables, you can output any of them. Please note that you do not have to minimize the number of increments.

// Input
// Each test contains multiple test cases. The first line contains the number of test cases t (1≤t≤10). Description of the test cases follows.
// The first line of each test case contains two integers n,m (1≤n≤100, 1≤m≤100)  — the number of rows and columns, respectively.
// The following n lines each contain m integers. The j-th integer in the i-th line is ai,j (1≤ai,j≤10^9).

// Output
// For each case, output n lines each containing m integers. The j-th integer in the i-th line is bi,j.

// Example
// input
// 3
// 3 2
// 1 2
// 4 5
// 7 8
// 2 2
// 1 1
// 3 3
// 2 2
// 1 3
// 2 2
// output
// 1 2
// 5 6
// 7 8
// 2 1
// 4 3
// 2 4
// 3 2

// Note
// In all the cases, you can verify that no two adjacent cells have the same value and that b is the same as a with some values incremented by one.

// Tutorial
// The increment by one operation essentially allows us to change the parity of any position. Let's color the matrix like a chessboard. Since every pair of adjacent cells consist of cells with different colors, we can make values at all black cells even and values at all white cells odd.

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
	var t, n, m, a, i, j uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &m)

			for i = 0; i < n; i++ {
				for j = 0; j < m; j++ {
					fmt.Fscan(reader, &a)

					if (i+j)&1 != a&1 {
						a++
					}

					fmt.Fprint(writer, a)

					if j+1 < m {
						fmt.Fprint(writer, " ")
					}
				}

				fmt.Fprintln(writer)
			}
		}
	}

	writer.Flush()
}
