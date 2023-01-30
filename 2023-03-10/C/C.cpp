// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=21&page=show_problem&problem=1886

// UVA 10945 - Mother bear

// Unforunately for our lazy “heroes”, the nuts were
// planted by an evil bear known as.. Dave, and they’ve
// fallen right into his trap. Dave is not just any bear,
// he’s a talking bear, but he can only understand sentences that are palindromes. While Larry was dazed
// and confused, Ryan figured this out, but need a way
// to make sure his sentences are palindromic. So he
// pulled out his trustly iPod, which thankfully have
// this program you wrote just for this purpose... or did
// you?

// Input

// You’ll be given many sentences. You have to determine if they are palindromes or not, ignoring case and punctuations. Every sentence will only contain
// the letters A-Z, a-z, ‘.’, ‘,’, ‘!’, ‘?’. The end of input will be a line containing the word ‘DONE’, which
// should not be processed.

// Output

// On each input, output ‘You won't be eaten!’ if it is a palindrome, and ‘Uh oh..’ if it is not a
// palindrome.

// Sample Input

// Madam, Im adam!
// Roma tibi subito motibus ibit amor.
// Me so hungry!
// Si nummi immunis
// DONE

// Sample Output

// You won't be eaten!
// You won't be eaten!
// Uh oh..
// You won't be eaten!

// palindrome

#include <algorithm>
#include <iostream>
#include <string>

bool solve(const std::string & s) {
    std::string t;
    unsigned i;

    for (const char & c : s) {
        if (std::isalpha(c)) {
            t += std::tolower(c);
        }
    }

    std::transform(t.begin(), t.end(), t.begin(), ::tolower);
    
    for (i = 0; i < t.size() / 2; i++) {
        if (t[i] != t[t.size() - i - 1]) {
            return false;
        }
    }

    return true;
}

int main () {
    std::string s;

    while (std::getline(std::cin, s) && s != "DONE") {
        if (solve(s)) {
            std::cout << "You won't be eaten!" << std::endl;
        } else {
            std::cout << "Uh oh.." << std::endl;
        }
    }

    return 0;
}
