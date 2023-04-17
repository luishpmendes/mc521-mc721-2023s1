// https://codeforces.com/problemset/problem/1791/E

// CodeForces 1791E - Negatives and Positives

// Given an array a consisting of n elements, find the maximum possible sum the array can have after performing the following operation any number of times:
//  • Choose 2 adjacent elements and flip both of their signs. In other words choose an index i such that 1≤i≤n−1 and assign ai=−ai and ai+1=−ai+1.

// Input
// The input consists of multiple test cases. The first line contains an integer t (1≤t≤1000) — the number of test cases. The descriptions of the test cases follow.
// The first line of each test case contains an integer n (2≤n≤2⋅10^5) — the length of the array.
// The following line contains n space-separated integers a1,a2,…,an (−10^9≤ai≤10^9).
// It is guaranteed that the sum of n over all test cases does not exceed 2⋅10^5.

// Output
// For each test case, output the maximum possible sum the array can have after performing the described operation any number of times.

// Example
// input
// 5
// 3
// -1 -1 -1
// 5
// 1 5 -5 0 2
// 3
// 1 2 3
// 6
// -1 10 9 8 7 6
// 2
// -1 -1
// output
// 1
// 13
// 6
// 39
// 2

// Note
// For the first test case, by performing the operation on the first two elements, we can change the array from [−1,−1,−1] to [1,1,−1], and it can be proven this array obtains the maximum possible sum which is 1+1+(−1)=1.
// For the second test case, by performing the operation on −5 and 0, we change the array from [1,5,−5,0,2] to [1,5,−(−5),−0,2]=[1,5,5,0,2], which has the maximum sum since all elements are non-negative. So, the answer is 1+5+5+0+2=13.
// For the third test case, the array already contains only positive numbers, so performing operations is unnecessary. The answer is just the sum of the whole array, which is 1+2+3=6.

// Tutorial
// We can notice that by performing any number of operations, the parity of the count of negative numbers won't ever change. Thus, if the number of negative numbers is initially even, we can make it equal to 0 by performing some operations. So, for an even count of negative numbers, the answer is the sum of the absolute values of all numbers (since we can make all of them positive). And if the count of negative numbers is odd, we must have one negative number at the end. We will choose the one smallest by absolute value and keep the rest positive (for simplicity, we consider −0 as a negative number).

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
	var t, n, i, sum, min uint
	var a []int
	var flag bool

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			a = make([]int, n)
			sum = 0
			flag = false
			min = 1e9

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a[i])

				if a[i] < 0 {
					flag = !flag
					a[i] = -a[i]
				}

				sum += uint(a[i])

				if min > uint(a[i]) {
					min = uint(a[i])
				}
			}

			if flag {
				sum -= min << 1
			}

			fmt.Fprintln(writer, sum)
		}
	}

	writer.Flush()
}
