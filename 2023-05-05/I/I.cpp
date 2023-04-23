// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=26&page=show_problem&problem=2445

// 11450 - Wedding shopping

// One of our best friends is getting married and we all
// are nervous because he is the first of us who is doing
// something similar. In fact, we have never assisted
// to a wedding, so we have no clothes or accessories,
// and to solve the problem we are going to a famous
// department store of our city to buy all we need: a
// shirt, a belt, some shoes, a tie, etcetera.
// We are offered different models for each class of
// garment (for example, three shirts, two belts, four
// shoes, ...). We have to buy one model of each class of
// garment, and just one.
// As our budget is limited, we cannot spend more
// money than it, but we want to spend the maximum possible. It’s possible that we cannot buy one
// model of each class of garment due to the short amount of money we have.

// Input

// The first line of the input contains an integer,N, indicating the number of test cases. For each test case,
// some lines appear, the first one contains two integers, M and C, separated by blanks (1 ≤ M ≤ 200,
// and 1 ≤ C ≤ 20), where M is the available amount of money and C is the number of garments you
// have to buy. Following this line, there are C lines, each one with some integers separated by blanks; in
// each of these lines the first integer, K (1 ≤ K ≤ 20), indicates the number of different models for each
// garment and it is followed by K integers indicating the price of each model of that garment.

// Output

// For each test case, the output should consist of one integer indicating the maximum amount of money
// necessary to buy one element of each garment without exceeding the initial amount of money. If there
// is no solution, you must print ‘no solution’.

// Sample Input

// 3
// 100 4
// 3 8 6 4
// 2 5 10
// 4 1 3 3 7
// 4 50 14 23 8
// 20 3
// 3 4 6 8
// 2 5 10
// 4 1 3 5 5
// 5 3
// 3 6 4 8
// 2 10 6
// 4 7 3 1 7

// Sample Output

// 75
// 19
// no solution

// discussed

#include <algorithm>
#include <cstring>
#include <iostream>
#include <limits>
#include <vector>

int main() {
    unsigned g, money, k, TC, M, C;
    // price[g (<= 20)][model (<= 20)]
    std::vector<std::vector<unsigned>> price;
    // reachable table[g (<= 20)][money (<= 200)]
    std::vector<std::vector<bool>> reachable;
    
    while (std::cin >> TC) {
        while (TC--) {
            std::cin >> M >> C;

            price.resize(C, std::vector<unsigned>(21));
            reachable.resize(C, std::vector<bool>(201, false));
            reachable.assign(C, std::vector<bool>(201, false));

            for (g = 0; g < C; g++) {
                // we store K in price[g][0]
                std::cin >> price[g][0];

                for (money = 1; money <= price[g][0]; money++) {
                    std::cin >> price[g][money];
                }
            }

            // initial values (base cases)
            for (g = 1; g <= price[0][0]; g++) {
                // to prevent array index out of bound
                if (M >= price[0][g]) {
                    // using first garment g = 0
                    reachable[0][M - price[0][g]] = true;
                }
            }

            // for each remaining garment
            for (g = 1; g < C; g++) {
                for (money = 0; money < M; money++) {
                    if (reachable[g - 1][money]) {
                        for (k = 1; k <= price[g][0]; k++) {
                            if (money >= price[g][k]) {
                                // also reachable now
                                reachable[g][money - price[g][k]] = true;
                            }
                        }
                    }
                }
            }

            for (money = 0; money <= M && !reachable[C - 1][money]; money++);

            if (money == M + 1) {
                // last row has no on bit
                std::cout << "no solution" << std::endl;
            } else {
                std::cout << M - money << std::endl;
            }
        }
    }

    return 0;
}
