// https://codeforces.com/problemset/problem/1472/B

// CodeForces 1472B - Fair Division

// Alice and Bob received n candies from their parents. Each candy weighs either 1 gram or 2 grams. Now they want to divide all candies among themselves fairly so that the total weight of Alice's candies is equal to the total weight of Bob's candies.
// Check if they can do that.
// Note that candies are not allowed to be cut in half.

// Input
// The first line contains one integer t (1≤t≤10^4) — the number of test cases. Then t test cases follow.
// The first line of each test case contains an integer n (1≤n≤100) — the number of candies that Alice and Bob received.
// The next line contains n integers a1,a2,…,an — the weights of the candies. The weight of each candy is either 1 or 2.
// It is guaranteed that the sum of n over all test cases does not exceed 10^5.

// Output
// For each test case, output on a separate line:
//  • "YES", if all candies can be divided into two sets with the same weight;
//  • "NO" otherwise.
// You can output "YES" and "NO" in any case (for example, the strings yEs, yes, Yes and YES will be recognized as positive).

// Example
// input
// 5
// 2
// 1 1
// 2
// 1 2
// 4
// 1 2 1 2
// 3
// 2 2 2
// 3
// 2 1 2
// output
// YES
// NO
// YES
// NO
// NO

// Note
// In the first test case, Alice and Bob can each take one candy, then both will have a total weight of 1.
// In the second test case, any division will be unfair.
// In the third test case, both Alice and Bob can take two candies, one of weight 1 and one of weight 2.
// In the fourth test case, it is impossible to divide three identical candies between two people.
// In the fifth test case, any division will also be unfair.

// Tutorial
// If the sum of all the weights is not divisible by two, then it is impossible to divide the candies between two people. If the sum is divisible, then let's count the number of candies with a weight of 1 and 2. Now, if we can find a way to collect half the sum with some candies, then these candies can be given to Alice, and all the rest can be given to Bob.
// Simple solution — let's iterate through how many candies of weight 2 we will give to Alice, then the remaining weight should be filled by candies of weight 1. If there are enough of them, then we have found a way of division.
// In fact, if the sum is even and there are at least two candies with weight 1 (there can't be one candy), then the answer is always "YES" (we can collect the weight as close to half as possible with weight 2 and then add weight 1). If there are no candies with weight 1, then you need to check whether n is even (since all the candies have the same weight, you just need to divide them in half).

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
	var t, n, a, i, sum uint
	var has1 bool

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			sum = 0
			has1 = false

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a)

				sum += a

				if a == 1 {
					has1 = true
				}
			}

			if sum&1 == 1 || !has1 && (sum>>1)&1 == 1 {
				fmt.Fprintln(writer, "NO")
			} else {
				fmt.Fprintln(writer, "YES")
			}
		}
	}

	writer.Flush()
}
