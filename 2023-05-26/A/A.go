// https://codeforces.com/problemset/problem/1367/B

// CodeForces 1367B - Even Array

// You are given an array a[0…n−1] of length n which consists of non-negative integers. Note that array indices start from zero.
// An array is called good if the parity of each index matches the parity of the element at that index. More formally, an array is good if for all i (0≤i≤n−1) the equality i mod 2 = a[i] mod 2 holds, where x mod 2 is the remainder of dividing x by 2.
// For example, the arrays [0,5,2,1] and [0,17,0,3] are good, and the array [2,4,6,7] is bad, because for i=1, the parities of i and a[i] are different: i mod 2 = 1 mod 2 = 1, but a[i] mod 2 = 4 mod 2 = 0.
// In one move, you can take any two elements of the array and swap them (these elements are not necessarily adjacent).
// Find the minimum number of moves in which you can make the array a good, or say that this is not possible.

// Input
// The first line contains a single integer t (1≤t≤1000) — the number of test cases in the test. Then t test cases follow.
// Each test case starts with a line containing an integer n (1≤n≤40) — the length of the array a.
// The next line contains n integers a0,a1,…,an−1 (0≤ai≤1000) — the initial array.

// Output
// For each test case, output a single integer — the minimum number of moves to make the given array a good, or -1 if this is not possible.

// Example
// input
// 4
// 4
// 3 2 7 6
// 3
// 3 2 6
// 1
// 7
// 7
// 4 9 2 1 18 3 0
// output
// 2
// 1
// -1
// 0

// Note
// In the first test case, in the first move, you can swap the elements with indices 0 and 1, and in the second move, you can swap the elements with indices 2 and 3.
// In the second test case, in the first move, you need to swap the elements with indices 0 and 1.
// In the third test case, you cannot make the array good.

// Tutorial
// We split all the positions in which the parity of the index does not match with the parity of the element into two arrays. If there is an odd number in the even index, add this index to the e array. Otherwise, if there is an even number in the odd index, add this index to the o array. Note that if the sizes of the o and e arrays are not equal, then there is no answer. Otherwise, the array a can be made good by doing exactly |o| operations by simply swapping all the elements in the o and e arrays.

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
	var t, n, a, i, even_mismatchs, odd_mismatchs uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)

			even_mismatchs = 0
			odd_mismatchs = 0

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a)

				if i&1 != a&1 {
					if i&1 == 0 {
						even_mismatchs++
					} else {
						odd_mismatchs++
					}
				}
			}

			if even_mismatchs != odd_mismatchs {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, even_mismatchs)
			}
		}
	}

	writer.Flush()
}
