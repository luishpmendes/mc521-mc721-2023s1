// https://codeforces.com/problemset/problem/1808/A

// CodeForces 1808A - Lucky Numbers
// Olympus City recently launched the production of personal starships. Now everyone on Mars can buy one and fly to other planets inexpensively.
// Each starship has a number —some positive integer x. Let's define the luckiness of a number x as the difference between the largest and smallest digits of that number. For example, 142857 has 8 as its largest digit and 1 as its smallest digit, so its luckiness is 8−1=7. And the number 111 has all digits equal to 1, so its luckiness is zero.
// Hateehc is a famous Martian blogger who often flies to different corners of the solar system. To release interesting videos even faster, he decided to buy himself a starship. When he came to the store, he saw starships with numbers from l to r inclusively. While in the store, Hateehc wanted to find a starship with the luckiest number.
// Since there are a lot of starships in the store, and Hateehc can't program, you have to help the blogger and write a program that answers his question.

// Input
// The first line contains an integer t (1≤t≤10000) —the number of test cases.
// Each of the following t lines contains a description of the test case. The description consists of two integers l and r (1≤l≤r≤10^6) — the largest and smallest numbers of the starships in the store.

// Output
// Print t lines, one line for each test case, containing the luckiest starship number in the store.
// If there are several ways to choose the luckiest number, output any of them.

// Example
// input
// 5
// 59 63
// 42 49
// 15 15
// 53 57
// 1 100
// output
// 60
// 49
// 15
// 57
// 90

// Note
// Let's look at two test examples:
//  • the luckiness of the number 59 is 9−5=4;
//  • the luckiness of 60 equals 6−0=6;
//  • the luckiness of 61 equals 6−1=5;
//  • the luckiness of 62 equals 6−2=4;
//  • the luckiness of 63 is 6−3=3.
// Thus, the luckiest number is 60.
// In the fifth test example, the luckiest number is 90.

// Tutorial
// Let r−l≥100. Then it is easy to see that there exists a number k such that l≤k≤r, and k≡90mod100. Then the number k is the answer.
// If r−l<100, you can find the answer by trying all the numbers.
// Despite the impressive constant, from the theoretical point of view we obtain a solution with asymptotics O(1) for the answer to the query.

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
	var t, l, r, ans, i, max_luckiness, max_luck_number, luckiness, number, min_digit, max_digit, digit uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &l, &r)
			ans = 0

			if r >= l+100 {
				for i = r; i >= l && ans == 0; i-- {
					if i%100 == 90 {
						ans = i
					}
				}
			} else {
				max_luckiness = 0
				max_luck_number = 0

				for i = r; i >= l && max_luckiness < 9; i-- {
					number = i
					min_digit = 9
					max_digit = 0

					for number > 0 {
						digit = number % 10
						number /= 10

						if min_digit > digit {
							min_digit = digit
						}

						if max_digit < digit {
							max_digit = digit
						}
					}

					luckiness = max_digit - min_digit

					if max_luckiness < luckiness || max_luck_number == 0 {
						max_luckiness = luckiness
						max_luck_number = i
					}
				}

				ans = max_luck_number
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
