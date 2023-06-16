// https://codeforces.com/problemset/problem/1814/A

// CodeForces 1814A - Coins
// In Berland, there are two types of coins, having denominations of 2 and k burles.
// Your task is to determine whether it is possible to represent n burles in coins, i. e. whether there exist non-negative integers x and y such that 2⋅x+k⋅y=n.

// Input
// The first line contains a single integer t (1≤t≤10^4) — the number of test cases.
// The only line of each test case contains two integers n and k (1≤k≤n≤10^18; k≠2).

// Output
// For each test case, print YES if it is possible to represent n burles in coins; otherwise, print NO. You may print each letter in any case (YES, yes, Yes will all be recognized as positive answer, NO, no and nO will all be recognized as negative answer).

// Example
// input
// 4
// 5 3
// 6 1
// 7 4
// 8 8
// output
// YES
// YES
// NO
// YES

// Note
// In the first test case, you can take one coin with denomination 2 and one coin with denomination k=3.
// In the second test case, you can take three coins with denomination 2. Alternatively, you can take six coins with denomination k=1.
// In the third test case, there is no way to represent 7 burles.
// In the fourth test case, you can take one coin with denomination k=8.

// Tutorial
// Note that 2 coins with denomination k can be replaced with k coins with denomination 2. So, if the answer exists, then there is also such a set of coins, where there is no more than one coin with denomination k. Therefore, it is enough to iterate through the number of coins with denomination k (from 0 to 1) and check that the remaining number is non-negative and even (i. e. it can be represented as some number of coins with denomination 2).

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
	var t, n, k uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &k)

			if n%2 == 0 || (n >= k && (n-k)&1 == 0) {
				fmt.Fprintln(writer, "YES")
			} else {
				fmt.Fprintln(writer, "NO")
			}
		}
	}

	writer.Flush()
}
