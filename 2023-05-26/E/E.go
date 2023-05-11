// https://codeforces.com/problemset/problem/727/A

// CodeForces 727A - Transformation: from A to B

// Vasily has a number a, which he wants to turn into a number b. For this purpose, he can do two types of operations:
//  • multiply the current number by 2 (that is, replace the number x by 2·x);
//  • append the digit 1 to the right of current number (that is, replace the number x by 10·x + 1).
// You need to help Vasily to transform the number a into the number b using only the operations described above, or find that it is impossible.
// Note that in this task you are not required to minimize the number of operations. It suffices to find any way to transform a into b.

// Input
// The first line contains two positive integers a and b (1 ≤ a < b ≤ 109) — the number which Vasily has and the number he wants to have.

// Output
// If there is no way to get b from a, print "NO" (without quotes).
// Otherwise print three lines. On the first line print "YES" (without quotes). The second line should contain single integer k — the length of the transformation sequence. On the third line print the sequence of transformations x1, x2, ..., xk, where:
//  • x1 should be equal to a,
//  • xk should be equal to b,
//  • xi should be obtained from xi - 1 using any of two described operations (1 < i ≤ k).
// If there are multiple answers, print any of them.

// Examples

// input
// 2 162
// output
// YES
// 5
// 2 4 8 81 162

// input
// 4 42
// output
// NO

// input
// 100 40021
// output
// YES
// 5
// 100 200 2001 4002 40021

// Tutorial
// Let's solve this problem in reverse way — try to get the number A from B.
// Note, that if B ends with 1 the last operation was to append the digit 1 to the right of current number. Because of that let delete last digit of B and move to the new number.
// If the last digit is even the last operation was to multiply the current number by 2. Because of that let's divide B on 2 and move to the new number.
// In the other cases (if B ends with odd digit except 1) the answer is «NO».
// We need to repeat described algorithm after we got the new number. If on some step we got the number equals to A we find the answer, and if the new number is less than A the answer is «NO».

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
	var a, b uint
	var flag bool
	var path []uint

	for t, _ = fmt.Fscan(reader, &a, &b); t == 2; t, _ = fmt.Fscan(reader, &a, &b) {
		flag = true
		path = make([]uint, 0)

		for a < b && flag {
			path = append([]uint{b}, path...)

			if b&1 == 0 {
				b >>= 1
			} else {
				if b%10 == 1 {
					b /= 10
				} else {
					flag = false
				}
			}
		}

		if a != b {
			fmt.Fprintln(writer, "NO")
		} else {
			path = append([]uint{a}, path...)

			fmt.Fprintln(writer, "YES")
			fmt.Fprintln(writer, len(path))

			for _, v := range path {
				fmt.Fprint(writer, v, " ")
			}

			fmt.Fprintln(writer)
		}
	}

	writer.Flush()
}
