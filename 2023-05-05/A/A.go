// https://codeforces.com/problemset/problem/1598/A

// CodeForces 1598A - Computer Game

// Monocarp is playing a computer game. Now he wants to complete the first level of this game.
// A level is a rectangular grid of 2 rows and n columns. Monocarp controls a character, which starts in cell (1,1) — at the intersection of the 1-st row and the 1-st column.
// Monocarp's character can move from one cell to another in one step if the cells are adjacent by side and/or corner. Formally, it is possible to move from cell (x1,y1) to cell (x2,y2) in one step if |x1−x2|≤1 and |y1−y2|≤1. Obviously, it is prohibited to go outside the grid.
// There are traps in some cells. If Monocarp's character finds himself in such a cell, he dies, and the game ends.
// To complete a level, Monocarp's character should reach cell (2,n) — at the intersection of row 2 and column n.
// Help Monocarp determine if it is possible to complete the level.

// Input
// The first line contains a single integer t (1≤t≤100) — the number of test cases. Then the test cases follow. Each test case consists of three lines.
// The first line contains a single integer n (3≤n≤100) — the number of columns.
// The next two lines describe the level. The i-th of these lines describes the i-th line of the level — the line consists of the characters '0' and '1'. The character '0' corresponds to a safe cell, the character '1' corresponds to a trap cell.
// Additional constraint on the input: cells (1,1) and (2,n) are safe.

// Output
// For each test case, output YES if it is possible to complete the level, and NO otherwise.

// Example
// input
// 4
// 3
// 000
// 000
// 4
// 0011
// 1100
// 4
// 0111
// 1110
// 6
// 010101
// 101010
// output
// YES
// YES
// NO
// YES

// Note
// Consider the example from the statement.
// In the first test case, one of the possible paths is (1,1)→(2,2)→(2,3).
// In the second test case, one of the possible paths is (1,1)→(1,2)→(2,3)→(2,4).
// In the fourth test case, one of the possible paths is (1,1)→(2,2)→(1,3)→(2,4)→(1,5)→(2,6).

// Tutorial
// At first glance, it seems like a graph problem. And indeed, this problem can be solved by explicitly building a graph considering cells as the vertices and checking that there is a safe path from start to finish via DFS/BFS/DSU/any other graph algorithm or data structure you know. But there's a much simpler solution.
// Since there are only two rows in a matrix, it's possible to move from any cell in the column i to any cell in column i+1 (if they are both safe, of course). It means that as long as there is at least one safe cell in each column, it is possible to reach any column of the matrix (and the cell (2,n) as well).
// It's easy to see that if this condition is not met, there exists a column with two unsafe cells — and this also means that this column and columns to the right of it are unreachable. So, the problem is reduced to checking if there is a column without any unsafe cells. To implement this, you can read both rows of the matrix as strings (let these strings be s1 and s2) and check that there is a position i such that both s1[i] and s2[i] are equal to 1.

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
	var s1, s2 string
	var possible bool

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &s1, &s2)
			possible = true

			for i = 0; i < n && possible; i++ {
				if s1[i] == '1' && s2[i] == '1' {
					possible = false
				}
			}

			if possible {
				fmt.Fprintln(writer, "YES")
			} else {
				fmt.Fprintln(writer, "NO")
			}
		}
	}

	writer.Flush()
}
