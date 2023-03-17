// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=21&page=show_problem&problem=1917

// UVA 10976 - Fractions Again?!

// It is easy to see that for every fraction in the form 1
// k
// (k > 0), we can always find two positive integers
// x and y, x ≥ y, such that:
// 1/k = 1/x − 1/y
// Now our question is: can you write a program that counts how many such pairs of x and y there
// are for any given k?

// Input

// Input contains no more than 100 lines, each giving a value of k (0 < k ≤ 10000).

// Output

// For each k, output the number of corresponding (x, y) pairs, followed by a sorted list of the values of x
// and y, as shown in the sample output.

// Sample Input

// 2
// 12

// Sample Output

// 2
// 1/2 = 1/6 + 1/3
// 1/2 = 1/4 + 1/4
// 8
// 1/12 = 1/156 + 1/13
// 1/12 = 1/84 + 1/14
// 1/12 = 1/60 + 1/15
// 1/12 = 1/48 + 1/16
// 1/12 = 1/36 + 1/18
// 1/12 = 1/30 + 1/20
// 1/12 = 1/28 + 1/21
// 1/12 = 1/24 + 1/24

// total solutions is asked upfront; therefore do brute force twice

#include <iostream>
#include <utility>
#include <vector>

int main () {
    unsigned k, y;
    std::vector<std::pair<unsigned, unsigned>> solutions;

    while (std::cin >> k) {
        solutions.clear();

        for (y = k + 1; y <= 2 * k; y++) {
            if ((k * y) % (y - k) == 0) {
                solutions.push_back(std::make_pair((k * y) / (y - k), y));
            }
        }

        std::cout << solutions.size() << std::endl;

        for (const std::pair<unsigned, unsigned> & xy : solutions) {
            std::cout << "1/" << k << " = 1/" << xy.first << " + 1/" << xy.second << std::endl;
        }
    }

    return 0;
}
