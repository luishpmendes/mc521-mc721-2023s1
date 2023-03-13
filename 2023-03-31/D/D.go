// https://codeforces.com/problemset/problem/1555/A

// CodeForces 1555A - PizzaForces

// PizzaForces is Petya's favorite pizzeria. PizzaForces makes and sells pizzas of three sizes: small pizzas consist of 6 slices, medium ones consist of 8 slices, and large pizzas consist of 10 slices each. Baking them takes 15, 20 and 25 minutes, respectively.
// Petya's birthday is today, and n of his friends will come, so he decided to make an order from his favorite pizzeria. Petya wants to order so much pizza that each of his friends gets at least one slice of pizza. The cooking time of the order is the total baking time of all the pizzas in the order.
// Your task is to determine the minimum number of minutes that is needed to make pizzas containing at least n slices in total. For example:
//  • if 12 friends come to Petya's birthday, he has to order pizzas containing at least 12 slices in total. He can order two small pizzas, containing exactly 12 slices, and the time to bake them is 30 minutes;
//  • if 15 friends come to Petya's birthday, he has to order pizzas containing at least 15 slices in total. He can order a small pizza and a large pizza, containing 16 slices, and the time to bake them is 40 minutes;
//  • if 300 friends come to Petya's birthday, he has to order pizzas containing at least 300 slices in total. He can order 15 small pizzas, 10 medium pizzas and 13 large pizzas, in total they contain 15⋅6+10⋅8+13⋅10=300 slices, and the total time to bake them is 15⋅15+10⋅20+13⋅25=750 minutes;
//  • if only one friend comes to Petya's birthday, he can order a small pizza, and the time to bake it is 15 minutes.

// Input
// The first line contains a single integer t (1≤t≤10^4) — the number of testcases.
// Each testcase consists of a single line that contains a single integer n (1≤n≤10^16) — the number of Petya's friends.

// Output
// For each testcase, print one integer — the minimum number of minutes that is needed to bake pizzas containing at least n slices in total.

// Example
// input
// 6
// 12
// 15
// 300
// 1
// 9999999999999999
// 3
// output
// 30
// 40
// 750
// 15
// 25000000000000000
// 15

// Tutorial
// Note that the "speed" of cooking 1 slice of pizza is the same for all sizes — 1 slice of pizza for 2.5 minutes.
// If n is odd, then we will increase it by 1 (since the pizza is cooked only with an even number of pieces). Now the value of n is always even. If n<6, then for such n the answer is equal to the answer for n=6, so we can say that n=max(n,6). While n≥12 we can order a small pizza. Eventually the value of n will be equal to 6, 8 or 10. This means that for any n there will be a set of pizzas with exactly n slices. Then the answer is n∗2.5 (in the solution, it is better to use the formula n/2∗5).

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
	var t, n uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)

			if n+1 < 6 {
				fmt.Fprintln(writer, 15)
			} else {
				fmt.Fprintln(writer, (n+1)/2*5)
			}
		}
	}

	writer.Flush()
}
