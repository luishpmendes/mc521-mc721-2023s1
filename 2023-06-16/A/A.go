// https://codeforces.com/problemset/problem/520/B

// CodeForces 520B - Two Buttons

// Vasya has found a strange device. On the front panel of a device there are: a red button, a blue button and a display showing some positive integer. After clicking the red button, device multiplies the displayed number by two. After clicking the blue button, device subtracts one from the number on the display. If at some point the number stops being positive, the device breaks down. The display can show arbitrarily large numbers. Initially, the display shows number n.
// Bob wants to get number m on the display. What minimum number of clicks he has to make in order to achieve this result?

// Input
// The first and the only line of the input contains two distinct integers n and m (1 ≤ n, m ≤ 10^4), separated by a space .

// Output
// Print a single number — the minimum number of times one needs to push the button required to get the number m out of number n.

// Examples

// input
// 4 6
// output
// 2

// input
// 10 1
// output
// 9

// Note
// In the first example you need to push the blue button once, and then push the red button once.
// In the second example, doubling the number is unnecessary, so we need to push the blue button nine times.

// Tutorial
// The simplest solution is simply doing a breadth-first search. Construct a graph with numbers as vertices and edges leading from one number to another if an operation can be made to change one number to the other. We may note that it is never reasonable to make the number larger than 2m, so under provided limitations the graph will contain at most 2·104 vertices and 4·104 edges, and the BFS should work real fast.
// There is, however, an even faster solution. The problem can be reversed as follows: we should get the number n starting from m using the operations "add 1 to the number" and "divide the number by 2 if it is even".
// Suppose that at some point we perform two operations of type 1 and then one operation of type 2; but in this case one operation of type 2 and one operation of type 1 would lead to the same result, and the sequence would contain less operations then before. That reasoning implies that in an optimal answer more than one consecutive operation of type 1 is possible only if no operations of type 2 follow, that is, the only situation where it makes sense is when n is smaller than m and we just need to make it large enough. Under this constraint, there is the only correct sequence of moves: if n is smaller than m, we just add 1 until they become equal; else we divide n by 2 if it is even, or add 1 and then divide by 2 if it is odd. The length of this sequence can be found in .
// Challenge: suppose we have a generalized problem: we want to get n starting from m using two operations "subtract a" and "multiply by b". Generalize the solution to find the minimal number of moves to get from n to m in  time if a and b are coprime. Can you do it if a and b may have common divisors greater than 1?

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var n, m, ans, x, i uint

	for t, _ = fmt.Fscan(reader, &n, &m); t == 2; t, _ = fmt.Fscan(reader, &n, &m) {
		if n >= m {
			ans = n - m
		} else {
			x = m
			ans = 0

			for i = 1; ; i++ {
				if x&1 == 1 {
					ans++
					x++
				}

				x /= 2
				ans++

				if x < n {
					ans += n - x
					break
				}

				if x == n {
					break
				}
			}
		}

		fmt.Fprintln(writer, ans)
	}

	writer.Flush()
}
