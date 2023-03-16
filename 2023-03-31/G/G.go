// https://codeforces.com/problemset/problem/1343/A

// CodeForces 1343A - Candies

// Recently Vova found n candy wrappers. He remembers that he bought x candies during the first day, 2x candies during the second day, 4x candies during the third day, …, 2k−1x candies during the k-th day. But there is an issue: Vova remembers neither x nor k but he is sure that x and k are positive integers and k>1.
// Vova will be satisfied if you tell him any positive integer x so there is an integer k>1 that x+2x+4x+⋯+2k−1x=n. It is guaranteed that at least one solution exists. Note that k>1.
// You have to answer t independent test cases.

// Input
// The first line of the input contains one integer t (1≤t≤104) — the number of test cases. Then t test cases follow.
// The only line of the test case contains one integer n (3≤n≤109) — the number of candy wrappers Vova found. It is guaranteed that there is some positive integer x and integer k>1 that x+2x+4x+⋯+2k−1x=n.

// Output
// Print one integer — any positive integer value of x so there is an integer k>1 that x+2x+4x+⋯+2k−1x=n.

// Example
// input
// 7
// 3
// 6
// 7
// 21
// 28
// 999999999
// 999999984
// output
// 1
// 2
// 1
// 7
// 4
// 333333333
// 333333328

// Note
// In the first test case of the example, one of the possible answers is x=1,k=2. Then 1⋅1+2⋅1 equals n=3.
// In the second test case of the example, one of the possible answers is x=2,k=2. Then 1⋅2+2⋅2 equals n=6.
// In the third test case of the example, one of the possible answers is x=1,k=3. Then 1⋅1+2⋅1+4⋅1 equals n=7.
// In the fourth test case of the example, one of the possible answers is x=7,k=2. Then 1⋅7+2⋅7 equals n=21.
// In the fifth test case of the example, one of the possible answers is x=4,k=3. Then 1⋅4+2⋅4+4⋅4 equals n=28.

// Tutorial
// Notice that ∑i=0k−12i=2k−1. Thus we can replace the initial equation with the following: (2k−1)x=n. So we can iterate over all possible k in range [2;29] (because 230−1>109) and check if n is divisible by 2k−1. If it is then we can print x=n2k−1.
// P.S. I know that so many participants found the formula ∑i=0k−12i=2k−1 using geometric progression sum but there is the other way to understand this and it is a way more intuitive for me. Just take a look at the binary representation of numbers: we can notice that 20=1,21=10,22=100 and so on. Thus 20=1,20+21=11,20+21+22=111 and so on. And if we add one to this number consisting of k ones then we get 2k.

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
	var t, n, x, k, sum uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)

			for k = 2; k < 30; k++ {
				sum = (1 << k) - 1

				if n%sum == 0 {
					x = n / sum
					break
				}
			}

			fmt.Fprintln(writer, x)
		}
	}

	writer.Flush()
}
