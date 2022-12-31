// https://codeforces.com/problemset/problem/1338/A

// CodeForces 1338A - Powered Addition

// You have an array a of length n. For every positive integer x you are going to perform the following operation during the x-th second:
//  • Select some distinct indices i1,i2,…,ik which are between 1 and n inclusive, and add 2^(x−1) to each corresponding position of a. Formally, aij:=aij+2^(x−1) for j=1,2,…,k. Note that you are allowed to not select any indices at all.
// You have to make a nondecreasing as fast as possible. Find the smallest number T such that you can make the array nondecreasing after at most T seconds.
// Array a is nondecreasing if and only if a1≤a2≤…≤an.
// You have to answer t independent test cases.

// Input
// The first line contains a single integer t (1≤t≤10^4) — the number of test cases.
// The first line of each test case contains single integer n (1≤n≤10^5) — the length of array a. It is guaranteed that the sum of values of n over all test cases in the input does not exceed 105.
// The second line of each test case contains n integers a1,a2,…,an (−10^9≤ai≤10^9).

// Output
// For each test case, print the minimum number of seconds in which you can make a nondecreasing.

// Example
// input
// 3
// 4
// 1 7 6 5
// 5
// 1 2 3 4 5
// 2
// 0 -4
// output
// 2
// 0
// 3

// Note
// In the first test case, if you select indices 3,4 at the 1-st second and 4 at the 2-nd second, then a will become [1,7,7,8]. There are some other possible ways to make a nondecreasing in 2 seconds, but you can't do it faster.
// In the second test case, a is already nondecreasing, so answer is 0.
// In the third test case, if you do nothing at first 2 seconds and select index 2 at the 3-rd second, a will become [0,0].

// Tutorial
// First, let's define b as ideal destination of a, when we used operations.
// Observation 1. Whatever you select any b, there is only one way to make it, because there is no more than single way to make specific amount of addition. That means we just have to select optimal destination of a.
// For example, if you want to make a1 from 10 to 21, then you must do 10→11→13→21. There is no other way to make 10 to 21 using given operations.
// So now we have to minimize max(b1−a1,b2−a2,…,bn−an), as smaller differences leads to use shorter time to make a nondecreasing.
// Observation 2. b is optimal when bi is the maximum value among b1,b2,…,bi−1 and ai.
// Because for each position i, we have to make bi−ai as smallest possible. Since bi should be not smaller than previous b values and also ai, we derived such formula.
// So from position 1,2,…,n, greedily find a bi, and check how many seconds needed to convert ai to bi. The answer is maximum needed seconds among all positions.
// Time complexity is O(n), but you can do O(nlogn) with "std::set" or whatever.
// Behind story of D2C/D1A:
//  • Originally, there was a different problem for this position. But it used XOR. As I made new D2E/D1C problem, I threw old D2C/D1A away and put this.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, n, i uint
	var a, max, diff int

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			max = -1e9
			diff = 0

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a)

				if max < a {
					max = a
				}

				if diff < max-a {
					diff = max - a
				}
			}

			if diff == 0 {
				fmt.Fprintln(writer, 0)
			} else {
				fmt.Fprintln(writer, uint(math.Log2(float64(diff)))+1)
			}
		}
	}

	writer.Flush()
}
