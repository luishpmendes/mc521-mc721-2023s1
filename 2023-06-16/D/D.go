// https://codeforces.com/problemset/problem/1681/E

// CodeForces 1681E - Labyrinth Adventures

// You found a map of a weirdly shaped labyrinth. The map is a grid, consisting of n rows and n columns. The rows of the grid are numbered from 1 to n from bottom to top. The columns of the grid are numbered from 1 to n from left to right.
// The labyrinth has n layers. The first layer is the bottom left corner (cell (1,1)). The second layer consists of all cells that are in the grid and adjacent to the first layer by a side or a corner. The third layer consists of all cells that are in the grid and adjacent to the second layer by a side or a corner. And so on.
// The labyrinth with 5 layers, for example, is shaped as follows:
// The layers are separated from one another with walls. However, there are doors in these walls.
// Each layer (except for layer n) has exactly two doors to the next layer. One door is placed on the top wall of the layer and another door is placed on the right wall of the layer. For each layer from 1 to n−1 you are given positions of these two doors. The doors can be passed in both directions: either from layer i to layer i+1 or from layer i+1 to layer i.
// If you are standing in some cell, you can move to an adjacent by a side cell if a wall doesn't block your move (e.g. you can't move to a cell in another layer if there is no door between the cells).
// Now you have m queries of sort: what's the minimum number of moves one has to make to go from cell (x1,y1) to cell (x2,y2).

// Input
// The first line contains a single integer n (2≤n≤10^5) — the number of layers in the labyrinth.
// The i-th of the next n−1 lines contains four integers d1,x,d1,y,d2,x and d2,y (1≤d1,x,d1,y,d2,x,d2,y≤n) — the coordinates of the doors. Both cells are on the i-th layer. The first cell is adjacent to the top wall of the i-th layer by a side — that side is where the door is. The second cell is adjacent to the right wall of the i-th layer by a side — that side is where the door is.
// The next line contains a single integer m (1≤m≤2⋅10^5) — the number of queries.
// The j-th of the next m lines contains four integers x1,y1,x2 and y2 (1≤x1,y1,x2,y2≤n) — the coordinates of the cells in the j-th query.

// Output
// For each query, print a single integer — the minimum number of moves one has to make to go from cell (x1,y1) to cell (x2,y2).

// Examples

// input
// 2
// 1 1 1 1
// 10
// 1 1 1 1
// 1 1 1 2
// 1 1 2 1
// 1 1 2 2
// 1 2 1 2
// 1 2 2 1
// 1 2 2 2
// 2 1 2 1
// 2 1 2 2
// 2 2 2 2
// output
// 0
// 1
// 1
// 2
// 0
// 2
// 1
// 0
// 1
// 0

// input
// 4
// 1 1 1 1
// 2 1 2 2
// 3 2 1 3
// 5
// 2 4 4 3
// 4 4 3 3
// 1 2 3 3
// 2 2 4 4
// 1 4 2 3
// output
// 3
// 4
// 3
// 6
// 2

// Note
// Here is the map of the labyrinth from the second example. The doors are marked red.

// Tutorial
// WLOG, assume all queries ask to move from a lower layer to a higher layer. The first thing to notice in the problem is that it is always optimal to never go down a layer.
// You have an optimal path that is going down some layers, and then returning to the same layer. So it leaves a layer in some its cell and returns to it in some other cell (or the same one). The best distance it can achieve is the Manhattan distance between these two cells. However, we can also achieve the Manhattan distance by just going along this layer, and the answer will be at least as optimal.
// If the query asks about the cells of the same layer, just answer with the Manhattan distance. Otherwise, we can describe the path as follows: go from the first cell to some door on its layer, enter the door and go to another door on the next layer, so on until the layer of the second cell, where you go from a door to the second cell.
// Thus, we could potentially write dp[i][j] — the shortest distance from the start to the j-th door of the i-th layer. Initialize both doors of the first layer, take the best answer from the both doors of the last layer. That would be O(n) per query, which is too slow.
// Let's optimize it with some precalculations. In particular, we want to know the shortest distance between one door of some layer and one door of another layer.
// We can use the technique similar to binary lifting. Calculate the distance between a pair of doors on layers which are 2x apart for all x up to logn. Let d[pi][x][d1][d2] be the distance from door d1 of layer i to door d2 of layer i+2^x. dpi,0,d1,d2 can be initialized straightforwardly. Then, to calculate dp[i][x][d1][d2], we can use the values for x−1: dp[i][x−1][d1][t] and dp[i+^(2x−1)][x−1][t][d2] for some intermediate door t on layer i+2^(x−1).
// To obtain the answer, use O(logn) jumps to reach the layer one before the last one. Then iterate over the last door.
// Alternatively, you could pack this dynamic programming into a segment tree, use divide and conquer on queries or do square root decomposition.
// Overall complexity: O((n+m)logn).

package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t, n, m, lg, i, j, k, l, c, x1, x2, y1, y2, l1, l2, ans int
	var d [][]point
	var dp [][][][]int
	var ndp [2]int
	var tmp [2]int

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		d = make([][]point, n-1)

		for i = 0; i+1 < n; i++ {
			d[i] = make([]point, 2)

			for j = 0; j < 2; j++ {
				fmt.Fscan(reader, &d[i][j].x, &d[i][j].y)
				d[i][j].x--
				d[i][j].y--
			}
		}

		lg = 1

		for 1<<lg < n-1 {
			lg++
		}

		dp = make([][][][]int, n-2)

		for i = 0; i+2 < n; i++ {
			dp[i] = make([][][]int, lg)

			for j = 0; j < lg; j++ {
				dp[i][j] = make([][]int, 2)

				for k = 0; k < 2; k++ {
					dp[i][j][k] = make([]int, 2)

					for l = 0; l < 2; l++ {
						dp[i][j][k][l] = 1e18
					}
				}
			}
		}

		for i = 0; i+2 < n; i++ {
			for l = 0; l < 2; l++ {
				dp[i][0][0][l] = abs(d[i][0].x+1-d[i+1][l].x) + abs(d[i][0].y-d[i+1][l].y) + 1
				dp[i][0][1][l] = abs(d[i][1].x-d[i+1][l].x) + abs(d[i][1].y+1-d[i+1][l].y) + 1
			}
		}

		for j = 1; j < lg; j++ {
			for i = 0; i+2 < n; i++ {
				for k = 0; k < 2; k++ {
					for l = 0; l < 2; l++ {
						for c = 0; c < 2; c++ {
							if i+(1<<(j-1))+2 < n {
								dp[i][j][k][l] = min(dp[i][j][k][l], dp[i][j-1][k][c]+dp[i+(1<<(j-1))][j-1][c][l])
							}
						}
					}
				}
			}
		}

		fmt.Fscan(reader, &m)

		for ; m > 0; m-- {
			fmt.Fscan(reader, &x1, &y1, &x2, &y2)
			x1--
			y1--
			x2--
			y2--
			l1 = max(x1, y1)
			l2 = max(x2, y2)

			if l1 > l2 {
				l1, l2 = l2, l1
				x1, x2 = x2, x1
				y1, y2 = y2, y1
			}

			if l1 != l2 {
				ndp[0] = abs(x1-d[l1][0].x) + abs(y1-d[l1][0].y)
				ndp[1] = abs(x1-d[l1][1].x) + abs(y1-d[l1][1].y)

				for j = lg; j > 0; j-- {
					if l1+(1<<(j-1)) < l2 {
						tmp[0] = 1e18
						tmp[1] = 1e18

						for k = 0; k < 2; k++ {
							for l = 0; l < 2; l++ {
								tmp[l] = min(tmp[l], ndp[k]+dp[l1][j-1][k][l])
							}
						}

						ndp = tmp
						l1 += 1 << (j - 1)
					}
				}

				ans = 1e18
				ans = min(ans, ndp[0]+abs(d[l1][0].x+1-x2)+abs(d[l1][0].y-y2)+1)
				ans = min(ans, ndp[1]+abs(d[l1][1].x-x2)+abs(d[l1][1].y+1-y2)+1)
			} else {
				ans = abs(x1-x2) + abs(y1-y2)
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
