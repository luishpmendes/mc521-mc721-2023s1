// https://codeforces.com/problemset/problem/1360/D

// CodeForces 1360D - Buying Shovels

// Polycarp wants to buy exactly n shovels. The shop sells packages with shovels. The store has k types of packages: the package of the i-th type consists of exactly i shovels (1≤i≤k). The store has an infinite number of packages of each type.
// Polycarp wants to choose one type of packages and then buy several (one or more) packages of this type. What is the smallest number of packages Polycarp will have to buy to get exactly n shovels?
// For example, if n=8 and k=7, then Polycarp will buy 2 packages of 4 shovels.
// Help Polycarp find the minimum number of packages that he needs to buy, given that he:
//  • will buy exactly n shovels in total;
//  • the sizes of all packages he will buy are all the same and the number of shovels in each package is an integer from 1 to k, inclusive.

// Input
// The first line contains an integer t (1≤t≤100) — the number of test cases in the input. Then, t test cases follow, one per line.
// Each test case consists of two positive integers n (1≤n≤10^9) and k (1≤k≤10^9) — the number of shovels and the number of types of packages.

// Output
// Print t answers to the test cases. Each answer is a positive integer — the minimum number of packages.

// Example
// input
// 5
// 8 7
// 8 1
// 6 10
// 999999733 999999732
// 999999733 999999733
// output
// 2
// 8
// 1
// 999999733
// 1

// Note
// The answer to the first test case was explained in the statement.
// In the second test case, there is only one way to buy 8 shovels — 8 packages of one shovel.
// In the third test case, you need to buy a 1 package of 6 shovels.

// Tutorial
// If Polycarp buys a packages of b shovels and gets exactly n shovels in total, then a⋅b=n, that is, a and b are divisors of n. Then the problem reduces to the following, you need to find the maximum divisor of the number n not greater than k. To do this, iterate over all the numbers x from 1 to √n inclusive and check whether n is divisible by x. If so, then x and nx — are both divisors of n and you can use them to try to improve the answer.

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
	var t, n, k, ans, i uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &k)

			ans = n

			for i = 1; i*i <= n; i++ {
				if n%i == 0 {
					if i <= k {
						if ans*i > n {
							ans = n / i
						}
					}

					if n <= k*i {
						if ans > i {
							ans = i
						}
					}
				}
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
