// https://codeforces.com/problemset/problem/1759/C

// CodeForces 1759C - Thermostat
// Vlad came home and found out that someone had reconfigured the old thermostat to the temperature of a.
// The thermostat can only be set to a temperature from l to r inclusive, the temperature cannot change by less than x. Formally, in one operation you can reconfigure the thermostat from temperature a to temperature b if |a−b|≥x and l≤b≤r.
// You are given l, r, x, a and b. Find the minimum number of operations required to get temperature b from temperature a, or say that it is impossible.

// Input
// The first line of input data contains the single integer t (1≤t≤10^4) — the number of test cases in the test.
// The descriptions of the test cases follow.
// The first line of each case contains three integers l, r and x (−10^9≤l≤r≤10^9, 1≤x≤10^9) — range of temperature and minimum temperature change.
// The second line of each case contains two integers a and b (l≤a,b≤r) — the initial and final temperatures.

// Output
// Output t numbers, each of which is the answer to the corresponding test case. If it is impossible to achieve the temperature b, output -1, otherwise output the minimum number of operations.

// Example
// input
// 10
// 3 5 6
// 3 3
// 0 15 5
// 4 5
// 0 10 5
// 3 7
// 3 5 6
// 3 4
// -10 10 11
// -5 6
// -3 3 4
// 1 0
// -5 10 8
// 9 2
// 1 5 1
// 2 5
// -1 4 3
// 0 2
// -6 3 6
// -1 -4
// output
// 0
// 2
// 3
// -1
// 1
// -1
// 3
// 1
// 3
// -1

// Note
// In the first example, the thermostat is already set up correctly.
// In the second example, you can achieve the desired temperature as follows: 4→10→5.
// In the third example, you can achieve the desired temperature as follows: 3→8→2→7.
// In the fourth test, it is impossible to make any operation.

// Tutorial
// First let's consider the cases when the answer exists:
//  • If a=b, then the thermostat is already set up and the answer is 0.
//  • else if |a−b|≥x, then it is enough to reconfigure the thermostat in 1 operation.
//  • else if exist such temperature c, that |a−c|≥x and |b−c|≥x, then you can configure the thermostat in 2 operations. If such c exists between l and r, we can chose one of bounds: a→l→b or a→r→b.
//  • we need to make 3 operations if times if we cannot reconfigure through one of the boundaries as above, but we can through both: a→l→r→b or a→r→l→b
// If we can't get the temperature b in one of these ways, the answer is −1.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, t, l, r, a, b, x int

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &l, &r, &x, &a, &b)

			if a == b {
				fmt.Fprintln(writer, 0)
			} else if a >= x+b || b >= x+a {
				fmt.Fprintln(writer, 1)
			} else if (a <= r-x && b <= r-x) || (a >= x+l && b >= x+l) {
				fmt.Fprintln(writer, 2)
			} else if (r >= x+b && a >= x+l) || (r >= x+a && b >= x+l) {
				fmt.Fprintln(writer, 3)
			} else {
				fmt.Fprintln(writer, -1)
			}
		}
	}

	writer.Flush()
}
