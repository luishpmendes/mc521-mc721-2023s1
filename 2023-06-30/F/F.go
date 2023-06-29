// https://codeforces.com/problemset/problem/1797/A

// CodeForces 1797A - Li Hua and Maze
// There is a rectangular maze of size n×m. Denote (r,c) as the cell on the r-th row from the top and the c-th column from the left. Two cells are adjacent if they share an edge. A path is a sequence of adjacent empty cells.
// Each cell is initially empty. Li Hua can choose some cells (except (x1,y1) and (x2,y2)) and place an obstacle in each of them. He wants to know the minimum number of obstacles needed to be placed so that there isn't a path from (x1,y1) to (x2,y2).
// Suppose you were Li Hua, please solve this problem.

// Input
// The first line contains the single integer t (1≤t≤500) — the number of test cases.
// The first line of each test case contains two integers n,m (4≤n,m≤10^9) — the size of the maze.
// The second line of each test case contains four integers x1,y1,x2,y2 (1≤x1,x2≤n,1≤y1,y2≤m) — the coordinates of the start and the end.
// It is guaranteed that |x1−x2|+|y1−y2|≥2.

// Output
// For each test case print the minimum number of obstacles you need to put on the field so that there is no path from (x1,y1) to (x2,y2).

// Example
// input
// 3
// 4 4
// 2 2 3 3
// 6 7
// 1 1 2 3
// 9 9
// 5 1 3 6
// output
// 4
// 2
// 3

// Note
// In test case 1, you can put obstacles on (1,3),(2,3),(3,2),(4,2). Then the path from (2,2) to (3,3) will not exist.

// Tutorial
// We can put obstacles around (x1,y1) or (x2,y2) and the better one is the answer. More formally, let's define a function f:
// f(x,y)=⎧⎩⎨2,3,4,(x,y) is on the corner(x,y) is on the border(x,y) is in the middle
// Then the answer is min{f(x1,y1),f(x2,y2)}.
// Without loss of generality, assume that f(x1,y1)≤f(x2,y2). As the method is already given, the answer is at most f(x1,y1). Let's prove that the answer is at least f(x1,y1).
// If (x1,y1) is on the corner, we can always find two paths from (x1,y1) to (x2,y2) without passing the same cell (except (x1,y1) and (x2,y2)).
// Similarly, we can always find three or four paths respectively if (x1,y1) is on the border or in the middle.
// As the paths have no common cell, we need to put an obstacle on each path, so the answer is at least f(x1,y1).
// In conclusion, the answer is exactly f(x1,y1). As we assumed that f(x1,y1)≤f(x2,y2), the answer to the original problem is min{f(x1,y1),f(x2,y2)}.
// Time complexity: O(1).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func f(n, m, x, y uint) uint {
	if (x == 1 || x == n) && (y == 1 || y == m) {
		return 2
	} else if x == 1 || x == n || y == 1 || y == m {
		return 3
	} else {
		return 4
	}
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, n, m, x1, y1, x2, y2, ans1, ans2 uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &m, &x1, &y1, &x2, &y2)
			ans1 = f(n, m, x1, y1)
			ans2 = f(n, m, x2, y2)

			if ans1 <= ans2 {
				fmt.Fprintln(writer, ans1)
			} else {
				fmt.Fprintln(writer, ans2)
			}
		}
	}

	writer.Flush()
}
