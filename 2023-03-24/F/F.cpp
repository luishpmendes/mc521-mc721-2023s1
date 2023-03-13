// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&page=show_problem&problem=2628

// UVA 11581 - Grid Successors

// Consider a 3 × 3 grid of numbers g where each cell contains either a ‘0’
// or a ‘1’. We define a function f that transforms such a grid: each cell
// of the grid f(g) is the sum (modulo 2) of its adjacent cells in g (two
// cells are considered adjacent if and only if they share a common side).
// Furthermore, we define f(i)(g) recursively f(0)(g) = g and
// f(i+1)(g) = f(f(i)(g)) (where i ≤ 0). Finally, for any grid h, let kg(h) be
// the number of indices i such that h = f(i)(g) (we may have kg(h) = ∞).
// Given a grid g, your task is to compute the greatest index i such that
// kg(f(i)(g)) is finite.

// Input

// Input begins with the number of test cases on its own line. Each case consists of a blank line followed
// by three lines of three characters, each either ‘1’ or ‘0’. The j’th character of the i’th row of the test
// case is the value in the j’th cell of the i’th row of the grid g.

// Output

// For each test case, output the greatest index i such that kg(f(i)(g)) is finite.
// If there is no such index, output ‘-1’.

// Sample Input

// 3
// 
// 111
// 100
// 001
// 
// 101
// 000
// 101
// 
// 000
// 000
// 000

// Sample Output

// 3
// 0
// -1

// simulate the process

#include <algorithm>
#include <iostream>
#include <string>
#include <vector>

std::vector<std::vector<bool>> f(const std::vector<std::vector<bool>> & g) {
    std::vector<std::vector<bool>> h(3, std::vector<bool>(3));
    unsigned i, j;

    for (i = 0; i < 3; i++) {
        for (j = 0; j < 3; j++) {
            h[i][j] = false;

            if (i > 0) {
                h[i][j] = h[i][j] ^ g[i - 1][j];
            }

            if (i < 2) {
                h[i][j] = h[i][j] ^ g[i + 1][j];
            }

            if (j > 0) {
                h[i][j] = h[i][j] ^ g[i][j - 1];
            }

            if (j < 2) {
                h[i][j] = h[i][j] ^ g[i][j + 1];
            }
        }
    }

    return h;
}

int main () {
    unsigned t, i, j;
    int ans;
    std::string s;
    std::vector<std::vector<bool>> g(3, std::vector<bool>(3));

    while (std::cin >> t) {
        while (t--) {
            for (i = 0; i < 3; i++) {
                std::cin >> s;

                for (j = 0; j < 3; j++) {
                    g[i][j] = (s[j] == '1');
                }
            }

            for (ans = 0; std::any_of(g.begin(), g.end(), [](const std::vector<bool> & v) {
                                                            return std::any_of(v.begin(), v.end(), [](bool b) {
                                                                                                        return b;
                                                                                                    });
                                                        }); ans++) {
                g = f(g);
            }

            std::cout << ans - 1 << std::endl;
        }
    }

    return 0;
}
