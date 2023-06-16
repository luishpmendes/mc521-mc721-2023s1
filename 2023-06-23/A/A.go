// https://codeforces.com/problemset/problem/1811/B

// CodeForces 1811B - Conveyor Belts
// Conveyor matrix mn is matrix of size n×n, where n is an even number. The matrix consists of concentric ribbons moving clockwise.
// In other words, the conveyor matrix for n=2 is simply a matrix 2×2, whose cells form a cycle of length 4 clockwise. For any natural k≥2, the matrix m2k is obtained by adding to the matrix m2k−2 an outer layer forming a clockwise cycle.
// You are standing in a cell with coordinates x1,y1 and you want to get into a cell with coordinates x2,y2. A cell has coordinates x,y if it is located at the intersection of the xth row and the yth column.
// Standing on some cell, every second you will move to the cell next in the direction of movement of the tape on which you are. You can also move to a neighboring cell by spending one unit of energy. Movements happen instantly and you can make an unlimited number of them at any time.
// Your task is to find the minimum amount of energy that will have to be spent to get from the cell with coordinates x1,y1 to the cell with coordinates x2,y2.
// For example, n=8 initially you are in a cell with coordinates 1,3 and you want to get into a cell with coordinates 6,4. You can immediately make 2 movements, once you are in a cell with coordinates 3,3, and then after 8 seconds you will be in the right cell.

// Input
// The first line contains an integer t (1≤t≤2⋅10^5) — the number of test cases.
// The descriptions of the test cases follow.
// The description of each test case consists of one string containing five integers n, x1, y1, x2 and y2 (1≤x1,y1,x2,y2≤n≤10^9) — matrix size and the coordinates of the start and end cells. It is guaranteed that the number n is even.

// Output
// For each test case, print one integer in a separate line — the minimum amount of energy that will have to be spent to get from the cell with coordinates x1,y1 to the cell with coordinates x2,y2.

// Example
// input
// 5
// 2 1 1 2 2
// 4 1 4 3 3
// 8 1 3 4 6
// 100 10 20 50 100
// 1000000000 123456789 987654321 998244353 500000004
// output
// 0
// 1
// 2
// 9
// 10590032

// Tutorial
// Note that the conveyor matrix n×n consists of n cycles, through each of which we can move without wasting energy. Now you need to find the distance between the cycles where the start and end cells are located. In one step from any cycle, you can go either to the cycle that is closer to the edge of the matrix, or to the cycle that is further from the edge of the matrix. It turns out that it is enough to find on which cycles there are cells on the edge and take their difference modulo.

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
	var t, n, x1, y1, x2, y2, layer1, layer2 uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &x1, &y1, &x2, &y2)

			layer1 = x1

			if layer1 > y1 {
				layer1 = y1
			}

			if layer1 > n-x1+1 {
				layer1 = n - x1 + 1
			}

			if layer1 > n-y1+1 {
				layer1 = n - y1 + 1
			}

			layer2 = x2

			if layer2 > y2 {
				layer2 = y2
			}

			if layer2 > n-x2+1 {
				layer2 = n - x2 + 1
			}

			if layer2 > n-y2+1 {
				layer2 = n - y2 + 1
			}

			if layer1 > layer2 {
				fmt.Fprintln(writer, layer1-layer2)
			} else {
				fmt.Fprintln(writer, layer2-layer1)
			}
		}
	}

	writer.Flush()
}
