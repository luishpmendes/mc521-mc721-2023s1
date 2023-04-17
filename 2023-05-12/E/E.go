// https://codeforces.com/problemset/problem/1113/A

// CodeForces 1113A - Sasha and His Trip

// Sasha is a very happy guy, that's why he is always on the move. There are n cities in the country where Sasha lives. They are all located on one straight line, and for convenience, they are numbered from 1 to n in increasing order. The distance between any two adjacent cities is equal to 1 kilometer. Since all roads in the country are directed, it's possible to reach the city y from the city x only if x<y.
// Once Sasha decided to go on a trip around the country and to visit all n cities. He will move with the help of his car, Cheetah-2677. The tank capacity of this model is v liters, and it spends exactly 1 liter of fuel for 1 kilometer of the way. At the beginning of the journey, the tank is empty. Sasha is located in the city with the number 1 and wants to get to the city with the number n. There is a gas station in each city. In the i-th city, the price of 1 liter of fuel is i dollars. It is obvious that at any moment of time, the tank can contain at most v liters of fuel.
// Sasha doesn't like to waste money, that's why he wants to know what is the minimum amount of money is needed to finish the trip if he can buy fuel in any city he wants. Help him to figure it out!

// Input
// The first line contains two integers n and v (2≤n≤100, 1≤v≤100)  — the number of cities in the country and the capacity of the tank.

// Output
// Print one integer — the minimum amount of money that is needed to finish the trip.

// Examples

// input
// 4 2
// output
// 4

// input
// 7 6
// output
// 6

// Note
// In the first example, Sasha can buy 2 liters for 2 dollars (1 dollar per liter) in the first city, drive to the second city, spend 1 liter of fuel on it, then buy 1 liter for 2 dollars in the second city and then drive to the 4-th city. Therefore, the answer is 1+1+2=4.
// In the second example, the capacity of the tank allows to fill the tank completely in the first city, and drive to the last city without stops in other cities.

// Tutorial
// When n−1<v the answer is n−1. Else you must notice, that it is optimal to fill the tank as soon as possible, because if you don't do that, you will have to spend more money in future. So to drive first v−1 kilometers just buy v−1 liters in the first city, then to drive between cities with numbers i and i+1, buy liter in the city with number i−v+1.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var n, v, result, i uint

	for t, _ = fmt.Fscan(reader, &n, &v); t == 2; t, _ = fmt.Fscan(reader, &n, &v) {
		if n < v+1 {
			result = n - 1
		} else {
			result = v - 1

			for i = 1; i+v <= n; i++ {
				result += i
			}
		}

		fmt.Fprintln(writer, result)
	}

	writer.Flush()
}
