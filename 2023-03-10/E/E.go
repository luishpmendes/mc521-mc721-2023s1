// https://codeforces.com/problemset/problem/1374/B

// CodeForces 1374B - Multiply by 2, divide by 6

// You are given an integer n. In one move, you can either multiply n by two or divide n by 6 (if it is divisible by 6 without the remainder).
// Your task is to find the minimum number of moves needed to obtain 1 from n or determine if it's impossible to do that.
// You have to answer t independent test cases.

// Input
// The first line of the input contains one integer t (1≤t≤2⋅10^4) — the number of test cases. Then t test cases follow.
// The only line of the test case contains one integer n (1≤n≤10^9).

// Output
// For each test case, print the answer — the minimum number of moves needed to obtain 1 from n if it's possible to do that or -1 if it's impossible to obtain 1 from n.

// Example
// input
// 7
// 1
// 2
// 3
// 12
// 12345
// 15116544
// 387420489
// output
// 0
// -1
// 2
// -1
// -1
// 12
// 36

// Note
// Consider the sixth test case of the example. The answer can be obtained by the following sequence of moves from the given integer 15116544:
//  1. Divide by 6 and get 2519424;
//  2. divide by 6 and get 419904;
//  3. divide by 6 and get 69984;
//  4. divide by 6 and get 11664;
//  5. multiply by 2 and get 23328;
//  6. divide by 6 and get 3888;
//  7. divide by 6 and get 648;
//  8. divide by 6 and get 108;
//  9. multiply by 2 and get 216;
// 10. divide by 6 and get 36;
// 11. divide by 6 and get 6;
// 12. divide by 6 and get 1.

// Tutorial
// If the number consists of other primes than 2 and 3 then the answer is -1. Otherwise, let cnt2 be the number of twos in the factorization of n and cnt3 be the number of threes in the factorization of n. If cnt2>cnt3 then the answer is -1 because we can't get rid of all twos. Otherwise, the answer is (cnt3−cnt2)+cnt3.
// Time complexity: O(logn).

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
	var t, n, cnt2, cnt3 uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			cnt2, cnt3 = 0, 0

			for n%2 == 0 {
				n /= 2
				cnt2++
			}

			for n%3 == 0 {
				n /= 3
				cnt3++
			}

			if n == 1 && cnt2 <= cnt3 {
				fmt.Fprintln(writer, cnt3-cnt2+cnt3)
			} else {
				fmt.Fprintln(writer, -1)
			}
		}
	}

	writer.Flush()
}
