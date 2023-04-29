// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=20&page=show_problem&problem=1768

// UVA 10827 - Maximum sum on a torus

// A grid that wraps both horizontally and vertically is called a torus.
// Given a torus where each cell contains an integer, determine the subrectangle with the largest sum. The sum of a sub-rectangle is the sum
// of all the elements in that rectangle. The grid below shows a torus
// where the maximum sub-rectangle has been shaded.

// Input

// The first line in the input contains the number of test cases (at most
// 18). Each case starts with an integer N (1 ≤ N ≤ 75) specifying the
// size of the torus (always square). Then follows N lines describing the
// torus, each line containing N integers between -100 and 100, inclusive.

// Output

// For each test case, output a line containing a single integer: the maximum sum of a sub-rectangle within
// the torus.

// Sample Input

// 2
// 5
// 1 -1 0 0 -4
// 2 3 -2 -3 2
// 4 1 -1 5 0
// 3 -2 1 -3 2
// -3 2 4 1 -4
// 3
// 1 2 3
// 4 5 6
// 7 8 9

// Sample Output

// 15
// 45

// copy n × n matrix into n × 2n matrix; then this problem becomes a standard problem again

#include <iostream>
#include <vector>

int main() {
    unsigned t, n, i, j, k, l;
    std::vector<std::vector<int>> m;
    int max_sub_rect, sub_rect;

    while (std::cin >> t) {
        while (t--) {
            std::cin >> n;
            m.resize(2 * n);
            m.assign(2 * n, std::vector<int>(2 * n));

            for (i = 0; i < n; i++) {
                for (j = 0; j < n; j++) {
                    std::cin >> m[i][j];
                    m[i + n][j] = m[i][j + n] = m[i + n][j + n] = m[i][j];
                }
            }

            for (i = 0; i < 2 * n; i++) {
                for (j = 0; j < 2 * n; j++) {
                    if (i > 0) {
                        // if possible, add from top
                        m[i][j] += m[i - 1][j];
                    }

                    if (j > 0) {
                        // if possible, add from left
                        m[i][j] += m[i][j - 1];
                    }

                    if (i > 0 && j > 0) {
                        // avoid double count
                        m[i][j] -= m[i - 1][j - 1];
                    }
                    //  inclusion-exclusion principle
                }
            }

            max_sub_rect = -100 * 75 * 75; // the lowest possible value for this problem

            for (i = 0; i < n; i++) { // start coordinate
                for (j = 0; j < n; j++) {  // start coordinate
                    for (k = i; k < i + n; k++) { // end coordinate
                        for (l = j; l < j + n; l++) { // end coordinate
                            sub_rect = m[k][l]; // sum of all items from (0, 0) to (k, l): O(1)

                            if (i > 0) {
                                // if possible, subtract from top
                                sub_rect -= m[i - 1][l];
                            }

                            if (j > 0) {
                                // if possible, subtract from left
                                sub_rect -= m[k][j - 1];
                            }

                            if (i > 0 && j > 0) {
                                // avoid double subtraction
                                sub_rect += m[i - 1][j - 1];
                            }

                            if (max_sub_rect < sub_rect) {
                                max_sub_rect = sub_rect;
                            }
                        }
                    }
                }
            }

            std::cout << max_sub_rect << std::endl;
        }
    }

    return 0;
}
