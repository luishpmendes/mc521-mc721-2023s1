// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=279&page=show_problem&problem=3886

// UVA 12455 - Bars

// Some things grow if you put them together.
// We have some metallic bars, theirs length known, and, if necessary, we want to solder some of them
// in order to obtain another one being exactly a given length long. No bar can be cut up. Is it possible?

// Input

// The first line of the input contains an integer,t, 0 ≤ t ≤ 50, indicating the number of test cases. For
// each test case, three lines appear, the first one contains a number n, 0 ≤ n ≤ 1000, representing the
// length of the bar we want to obtain. The second line contains a number p, 1 ≤ p ≤ 20, representing the
// number of bars we have. The third line of each test case contains p numbers, representing the length
// of the p bars.

// Output

// For each test case the output should contain a single line, consists of the string ‘YES’ or the string ‘NO’,
// depending on whether solution is possible or not.

// Sample Input

// 4
// 25
// 4
// 10 12 5 7
// 925
// 10
// 45 15 120 500 235 58 6 12 175 70
// 120
// 5
// 25 25 25 25 25
// 0
// 2
// 13 567

// Sample Output

// NO
// YES
// NO
// YES

// discussed

#include <iostream>
#include <vector>

int main () {
    unsigned t, n, p, sum, i, j;
    std::vector<unsigned> l;
    bool found;

    while (std::cin >> t) {
        while (t--) {
            std::cin >> n >> p;

            l.resize(p);

            for (unsigned & e : l) {
                std::cin >> e;
            }

            found = false;

            for (i = 0; i < (1u << p) && !found; i++) {
                sum = 0;

                for (j = 0; j < p; j++) {
                    if (i & (1 << j)) {
                        sum += l[j];
                    }
                }

                if (sum == n) {
                    found = true;
                }
            }

            if (found) {
                std::cout << "YES" << std::endl;
            } else {
                std::cout << "NO" << std::endl;
            }
        }
    }

    return 0;
}
