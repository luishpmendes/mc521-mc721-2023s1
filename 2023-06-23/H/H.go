// https://codeforces.com/problemset/problem/1816/A

// CodeForces 1816A - Ian Visits Mary
// Ian and Mary are frogs living on lattice points of the Cartesian coordinate plane, with Ian living on (0,0) and Mary living on (a,b).
// Ian would like to visit Mary by jumping around the Cartesian coordinate plane. Every second, he jumps from his current position (xp,yp) to another lattice point (xq,yq), such that no lattice point other than (xp,yp) and (xq,yq) lies on the segment between point (xp,yp) and point (xq,yq).
// As Ian wants to meet Mary as soon as possible, he wants to jump towards point (a,b) using at most 2 jumps. Unfortunately, Ian is not good at maths. Can you help him?
// A lattice point is defined as a point with both the x-coordinate and y-coordinate being integers.

// Input
// The first line contains a single integer t (1≤t≤500) — the number of test cases. The description of test cases follows.
// The first and only line of each test case contains two integers a and b (1≤a,b≤10^9) — the coordinates of the lattice point where Mary lives.

// Output
// For each test case, print an integer n (1≤n≤2) on the first line, denoting the number of jumps Ian uses in order to meet Mary. Note that you do not need to minimize the number of jumps.
// On the i-th line of the next n lines, print two integers 0≤xi,yi≤109 separated by a space, denoting Ian's location (xi,yi) after the i-th jump. xn=a, yn=b must hold.
// Ian's initial location and his locations after each of the n jumps need not be distinct.
// If there are multiple solutions, output any.

// Example
// input
// 8
// 3 4
// 4 4
// 3 6
// 2 2
// 1 1
// 7 3
// 2022 2023
// 1000000000 1000000000
// output
// 1
// 3 4
// 2
// 3 2
// 4 4
// 2
// 5 3
// 3 6
// 2
// 1 0
// 2 2
// 1
// 1 1
// 1
// 7 3
// 1
// 2022 2023
// 2
// 69420420 469696969
// 1000000000 1000000000

// Note
// In the first test case:
// (0,0)→(3,4)
// In the second test case:
// (0,0)→(3,2)→(4,4)
// In the third test case:
// (0,0)→(5,3)→(3,6)

// Tutorial
// Let us show that Ian can get to Mary within two moves. Indeed, consider (x1,y1) and (x2,y2) where x2−x1=1, since the value of the x-coordinates of a point lying on the segment joining (x1,y1) and (x2,y2), and is not at the endpoints, is strictly between x1 and x2, but there are no integers strictly between x1 and x2, the point must not be a lattice point. Similarly, no lattice points lie on the segment joining (x1,y1) and (x2,y2) with y2−y1=1 except for the endpoints. Therefore, Ian can jump from (0,0) to (x−1,1) to (x,y).

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
	var t, a, b uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &a, &b)
			fmt.Fprintln(writer, 2)
			fmt.Fprintln(writer, a-1, 1)
			fmt.Fprintln(writer, a, b)
		}
	}

	writer.Flush()
}
