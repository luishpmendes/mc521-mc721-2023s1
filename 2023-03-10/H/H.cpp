// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=21&page=show_problem&problem=1844

// UVA 10903 - Rock-Paper-Scissors Tournament

// Rock-Paper-Scissors is game for two players, A and B, who each choose, independently of the other,
// one of rock, paper, or scissors. A player chosing paper wins over a player chosing rock; a player chosing
// scissors wins over a player chosing paper; a player chosing rock wins over a player chosing scissors. A
// player chosing the same thing as the other player neither wins nor loses.
// A tournament has been organized in which each of n players plays k rock-scissors-paper games with
// each of the other players — k ∗ n ∗ (n − 1)/2 games in total. Your job is to compute the win average for
// each player, defined as w/(w + l) where w is the number of games won, and l is the number of games
// lost, by the player.

// Input

// Input consists of several test cases. The first line of input for each case contains 1 ≤ n ≤ 100 1 ≤ k ≤ 100
// as defined above. For each game, a line follows containing p1, m1, p2, m2. 1 ≤ p1 ≤ n and 1 ≤ p2 ≤ n
// are distinct integers identifying two players; m1 and m2 are their respective moves (‘rock’, ‘scissors’,
// or ‘paper’). A line containing ‘0’ follows the last test case.

// Output

// Output one line each for player 1, player 2, and so on, through player n, giving the player’s win average
// rounded to three decimal places. If the win average is undefined, output ‘-’. Output an empty line
// between cases.

// Sample Input

// 2 4
// 1 rock 2 paper
// 1 scissors 2 paper
// 1 rock 2 rock
// 2 rock 1 scissors
// 2 1
// 1 rock 2 paper
// 0

// Sample Output

// 0.333
// 0.667
// 
// 0.000
// 1.000

// count win+losses, output win average

#include <iomanip>
#include <iostream>
#include <string>
#include <vector>

int main () {
    unsigned n, k, p1, p2, i;
    std::string m1, m2;
    std::vector<unsigned> wins, loses;
    bool flag = false;

    std::cout << std::fixed << std::setprecision(3);

    while (std:: cin >> n >> k) {
        wins.resize(n, 0);
        wins.assign(n, 0);
        loses.resize(n, 0);
        loses.assign(n, 0);

        for (i = 0; 2 * i + k * n < k * n * n; i++) {
            std::cin >> p1 >> m1 >> p2 >> m2;

            if (m1 == "rock" && m2 == "scissors") {
                wins[p1 - 1]++;
                loses[p2 - 1]++;
            } else if (m1 == "scissors" && m2 == "paper") {
                wins[p1 - 1]++;
                loses[p2 - 1]++;
            } else if (m1 == "paper" && m2 == "rock") {
                wins[p1 - 1]++;
                loses[p2 - 1]++;
            } else if (m1 == "scissors" && m2 == "rock") {
                wins[p2 - 1]++;
                loses[p1 - 1]++;
            } else if (m1 == "paper" && m2 == "scissors") {
                wins[p2 - 1]++;
                loses[p1 - 1]++;
            } else if (m1 == "rock" && m2 == "paper") {
                wins[p2 - 1]++;
                loses[p1 - 1]++;
            }
        }

        if (flag) {
            std::cout << std::endl;
        } else {
            flag = true;
        }

        for (i = 0; i < n; i++) {
            if (wins[i] + loses[i] == 0) {
                std::cout << "-" << std::endl;
            } else {
                std::cout << (double) wins[i] / (wins[i] + loses[i]) << std::endl;
            }
        }
    }

    return 0;
}
