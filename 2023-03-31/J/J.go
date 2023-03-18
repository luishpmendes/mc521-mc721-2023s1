// https://codeforces.com/problemset/problem/466/B

// CodeForces 466B - Wonder Room

// The start of the new academic year brought about the problem of accommodation students into dormitories. One of such dormitories has a a × b square meter wonder room. The caretaker wants to accommodate exactly n students there. But the law says that there must be at least 6 square meters per student in a room (that is, the room for n students must have the area of at least 6n square meters). The caretaker can enlarge any (possibly both) side of the room by an arbitrary positive integer of meters. Help him change the room so as all n students could live in it and the total area of the room was as small as possible.

// Input

// The first line contains three space-separated integers n, a and b (1 ≤ n, a, b ≤ 109) — the number of students and the sizes of the room.

// Output

// Print three integers s, a1 and b1 (a ≤ a1; b ≤ b1) — the final area of the room and its sizes. If there are multiple optimal solutions, print any of them.

// Examples

// input
// 3 3 5
// output
// 18
// 3 6

// input
// 2 4 4
// output
// 16
// 4 4

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
	var n, a, b, sq, a1, b1, tmpb, i uint64
	var f bool

	for t, _ = fmt.Fscan(reader, &n, &a, &b); t == 3; t, _ = fmt.Fscan(reader, &n, &a, &b) {
		if 6*n <= a*b {
			fmt.Fprintln(writer, a*b)
			fmt.Fprintln(writer, a, b)
		} else {
			f = false

			if a > b {
				a, b = b, a
				f = true
			}

			sq = 1e18

			for i = a; i*i <= 6*n; i++ {
				tmpb = 6 * n / i

				if tmpb*i < 6*n {
					tmpb++
				}

				if tmpb > b && tmpb*i < sq {
					sq = tmpb * i
					a1 = i
					b1 = tmpb
				}
			}

			if f {
				a1, b1 = b1, a1
			}

			fmt.Fprintln(writer, sq)
			fmt.Fprintln(writer, a1, b1)
		}
	}

	writer.Flush()
}
