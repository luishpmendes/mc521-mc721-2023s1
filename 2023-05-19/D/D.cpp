// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=24&page=show_problem&problem=2267

// 11292 - Dragon of Loowater

// Once upon a time, in the Kingdom of Loowater, a minor nuisance
// turned into a major problem.
// The shores of Rellau Creek in
// central Loowater had always been
// a prime breeding ground for geese.
// Due to the lack of predators, the
// geese population was out of control.
// The people of Loowater mostly
// kept clear of the geese. Occasionally, a goose would attack one of
// the people, and perhaps bite off a
// finger or two, but in general, the
// people tolerated the geese as a minor nuisance.
// One day, a freak mutation
// occurred, and one of the geese
// spawned a multi-headed firebreathing dragon. When the
// dragon grew up, he threatened to
// burn the Kingdom of Loowater to
// a crisp. Loowater had a major problem. The king was alarmed, and called on his knights to slay the
// dragon and save the kingdom.
// The knights explained: “To slay the dragon, we must chop off all its heads. Each knight can chop
// off one of the dragon’s heads. The heads of the dragon are of different sizes. In order to chop off a
// head, a knight must be at least as tall as the diameter of the head. The knights’ union demands that
// for chopping off a head, a knight must be paid a wage equal to one gold coin for each centimetre of the
// knight’s height.”
// Would there be enough knights to defeat the dragon? The king called on his advisors to help him
// decide how many and which knights to hire. After having lost a lot of money building Mir Park, the
// king wanted to minimize the expense of slaying the dragon. As one of the advisors, your job was to
// help the king. You took it very seriously: if you failed, you and the whole kingdom would be burnt to
// a crisp!

// Input

// The input contains several test cases. The first line of each test case contains two integers between 1 and
// 20000 inclusive, indicating the number n of heads that the dragon has, and the number m of knights in
// the kingdom. The next n lines each contain an integer, and give the diameters of the dragon’s heads,
// in centimetres. The following m lines each contain an integer, and specify the heights of the knights of
// Loowater, also in centimetres.
// The last test case is followed by a line containing ‘0 0’.

// Output

// For each test case, output a line containing the minimum number of gold coins that the king needs to
// pay to slay the dragon. If it is not possible for the knights of Loowater to slay the dragon, output the
// line ‘Loowater is doomed!’.

// Sample Input

// 2 3
// 5
// 4
// 7
// 8
// 4
// 2 1
// 5
// 5
// 10
// 0 0

// Sample Output

// 11
// Loowater is doomed!

// Sort the Input First

#include <algorithm>
#include <iostream>
#include <vector>

int main() {
    unsigned n, m, gold, d, k;
    std::vector<unsigned> dragon, knight;

    while (std::cin >> n >> m && (n || m)) {
        dragon.resize(n);
        knight.resize(m);

        for (unsigned & x : dragon) {
            std::cin >> x;
        }

        for (unsigned & x : knight) {
            std::cin >> x;
        }

        // array dragon is sorted in non decreasing order
        std::sort(dragon.begin(), dragon.end());
        // array knight is sorted in non decreasing order
        std::sort(knight.begin(), knight.end());

        gold = d = k = 0;

        while (d < n && k < m) { // still have dragon heads or knights
            while (dragon[d] > knight[k] && k < m) {
                // find the required knight
                k++;
            }

            if (k == m) {
                // no knight can kill this dragon head, doomed :S
                break;
            }

            // the king pay this amount of gold
            gold += knight[k];

            // next dragon head please
            d++;
            // next knight please
            k++;
        }

        if (d == n) {
            // all dragon heads are chopped
            std::cout << gold << std::endl;
        } else {
            std::cout << "Loowater is doomed!" << std::endl;
        }
    }

    return 0;
}
