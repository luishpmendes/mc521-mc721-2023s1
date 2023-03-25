// https://codeforces.com/problemset/problem/1365/B

// CodeForces 1365B - Trouble Sort

// Ashish has n elements arranged in a line.
// These elements are represented by two integers ai — the value of the element and bi — the type of the element (there are only two possible types: 0 and 1). He wants to sort the elements in non-decreasing values of ai.
// He can perform the following operation any number of times:
//  • Select any two elements i and j such that bi≠bj and swap them. That is, he can only swap two elements of different types in one move.
// Tell him if he can sort the elements in non-decreasing values of ai after performing any number of operations.

// Input
// The first line contains one integer t (1≤t≤100) — the number of test cases. The description of the test cases follows.
// The first line of each test case contains one integer n (1≤n≤500) — the size of the arrays.
// The second line contains n integers ai (1≤ai≤10^5)  — the value of the i-th element.
// The third line containts n integers bi (bi∈{0,1})  — the type of the i-th element.

// Output
// For each test case, print "Yes" or "No" (without quotes) depending on whether it is possible to sort elements in non-decreasing order of their value.
// You may print each letter in any case (upper or lower).

// Example
// input
// 5
// 4
// 10 20 20 30
// 0 1 0 1
// 3
// 3 1 2
// 0 1 1
// 4
// 2 2 4 8
// 1 1 1 1
// 3
// 5 15 4
// 0 0 0
// 4
// 20 10 100 50
// 1 0 0 1
// output
// Yes
// Yes
// Yes
// No
// Yes

// Note
// For the first case: The elements are already in sorted order.
// For the second case: Ashish may first swap elements at positions 1 and 2, then swap elements at positions 2 and 3.
// For the third case: The elements are already in sorted order.
// For the fourth case: No swap operations may be performed as there is no pair of elements i and j such that bi≠bj. The elements cannot be sorted.
// For the fifth case: Ashish may swap elements at positions 3 and 4, then elements at positions 1 and 2.

// Tutorial
// Key Idea:
// If there is at least one element of type 0 and at least one element of type 1, we can always sort the array.
// Solution:
// If all the elements are of the same type, we cannot swap any two elements. So, in this case, we just need to check if given elements are already in sorted order.
// Otherwise, there is at least one element of type 0 and at least one element of type 1. In this case, it is possible to swap any two elements! We can swap elements of different types using only one operation. Suppose we want to swap two elements a and b of the same type. We can do it in 3 operations. Let c be an element of the type different from a and b. We can first swap a and c, then swap b and c and then swap a and c again. In doing so, c remains at its initial position and a, b are swapped. This is exactly how we swap two integers using a temporary variable. Since we can swap any two elements, it is always possible to sort the array in this case.
// Time complexity: O(n)

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
	var t, n, i, curr_a, prev_a uint
	var b, sorted, have0, have1 bool

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)

			sorted = true
			have0 = false
			have1 = false

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &curr_a)

				if i > 0 && curr_a < prev_a {
					sorted = false
				}

				prev_a = curr_a
			}

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &b)

				if !b {
					have0 = true
				} else {
					have1 = true
				}
			}

			if (have0 && have1) || sorted {
				fmt.Fprintln(writer, "Yes")
			} else {
				fmt.Fprintln(writer, "No")
			}
		}
	}

	writer.Flush()
}
