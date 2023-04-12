// https://codeforces.com/problemset/problem/1526/B

// CodeForces 1526B - I Hate 1111

// You are given an integer x. Can you make x by summing up some number of 11,111,1111,11111,…? (You can use any number among them any number of times).
// For instance,
//  • 33=11+11+11
//  • 144=111+11+11+11

// Input
// The first line of input contains a single integer t (1≤t≤10000) — the number of testcases.
// The first and only line of each testcase contains a single integer x (1≤x≤10^9) — the number you have to make.

// Output
// For each testcase, you should output a single string. If you can make x, output "YES" (without quotes). Otherwise, output "NO".
// You can print each letter of "YES" and "NO" in any case (upper or lower).

// Example
// input
// 3
// 33
// 144
// 69
// output
// YES
// YES
// NO

// Note
// Ways to make 33 and 144 were presented in the statement. It can be proved that we can't present 69 this way.

// Tutorial

// Hint 0
// Read the name of the problem ;)

// Hint 1
// 1111=11⋅101

// Hint 2
// All numbers other than 11 and 111 are useless.

// Solution

// Method 1

// Notice that 1111=11⋅100+11 and similarly 11111=111⋅100+11. This implies that we can construct 1111 and all bigger numbers using only 11 and 111. So it suffices to check whether we can construct X from 11 and 111 only.
// Suppose X=A⋅11+B⋅111, where A,B≥0. Suppose B=C⋅11+D, where D<11. Then X=(A+C⋅111)⋅11+D⋅111. So we can just brute force all 11 value of D to check whether X can be made.

// Method 2

// Since gcd(11,111)=1, by the Chicken McNugget Theorem, all numbers greater than 1099 can be written as a sum of 11 and 111. We can use brute force to find the answer to all values less than or equal to 1099 and answer yes for all other numbers.

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
	var t, x, i uint
	var ans bool

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &x)
			ans = false

			for i = 0; i < 11 && !ans; i++ {
				if x >= 111*i && (x-111*i)%11 == 0 {
					ans = true
				}
			}

			if ans {
				fmt.Fprintln(writer, "YES")
			} else {
				fmt.Fprintln(writer, "NO")
			}
		}
	}

	writer.Flush()
}
