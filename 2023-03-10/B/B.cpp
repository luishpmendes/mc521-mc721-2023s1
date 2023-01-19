// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=78&page=show_problem&problem=2725

// UVA 11678 - Cards' Exchange

// Alice and Betty collect Pok´mon cards. The cards are printed for a game that imitates the e battle
// system of one of the most popular videogames in history, but Alice and Betty are too young to actually
// play the game, and are only interested in the actual cards. To make it easier, we’ll assume each card
// has an unique identifier, given as an integer number.
// Both girls have a set of cards, and, like most girls their age, like to trade the cards they have. They
// obviously don’t care about identical cards they both have, and they don’t want to receive repeated
// cards in the trade. Besides, the cards are traded in a single operation: Alice gives Betty N distinct
// cards and receives back other N distinct cards. The girls want to know what is the maximum number
// of cards they can trade. For instance, if Alice has cards {1, 1, 2, 3, 5, 7, 8, 8, 9, 15} and Betty has
// cards {2, 2, 2, 3, 4, 6, 10, 11, 11}, they can trade at most four cards.
// Write a program that given the sets of cards owned by Alice and Betty, determines the maximum
// number of cards they can trade.

// Input

// The input contains several test cases. The first line of a test case contains two integers A and B,
// separated by a blank space, indicating respectively the number of cards Alice and Betty have (1 ≤
// A ≤ 104 and 1 ≤ B ≤ 10^4). The second line contains A space-separated integers X_i, each indicating
// one of Alice’s cards (1 ≤ Xi ≤ 10^5). The third line contains B integers Yi separated by whitespaces,
// each number indicating one of Betty’s cards (1 ≤ Yi ≤ 10^5). Alice and Betty’s cards are listed in
// non-descending order.
// The end of input is indicated by a line containing only two zeros, separated by a blank space.

// Output

// For each test case, your program should print a single line, containing an integer number, indicating
// the maximum number of cards Alice and Betty can trade.

// Sample Input

// 1 1
// 1000
// 1000
// 3 4
// 1 3 5
// 2 4 6 8
// 10 9
// 1 1 2 3 5 7 8 8 9 15
// 2 2 2 3 4 6 10 11 11
// 0 0

// Sample Output

// 0
// 3
// 4

// actually just an array manipulation problem

#include <algorithm>
#include <iostream>
#include <set>

int main () {
    unsigned a, b, x, i;
    std::set <unsigned> alice, betty, diff;

    while (std::cin >> a >> b && (a || b)) {
        alice.clear();
        betty.clear();
        diff.clear();

        for (i = 0; i < a; i++) {
            std::cin >> x;
            alice.insert(x);
        }

        for (i = 0; i < b; i++) {
            std::cin >> x;
            betty.insert(x);
        }


        if (alice.size() < betty.size()) {
            std::set_difference(alice.begin(), alice.end(), betty.begin(), betty.end(), std::inserter(diff, diff.begin()));
        } else {
            std::set_difference(betty.begin(), betty.end(), alice.begin(), alice.end(), std::inserter(diff, diff.begin()));
        }

        std::cout << diff.size() << std::endl;
    }

    return 0;
}
