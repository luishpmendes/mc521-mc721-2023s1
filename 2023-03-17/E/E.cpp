// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=3&page=show_problem&problem=92

// UVA 156 - Ananagrams

// Most crossword puzzle fans are used to anagrams — groups of words with the same letters in different
// orders — for example OPTS, SPOT, STOP, POTS and POST. Some words however do not have this
// attribute, no matter how you rearrange their letters, you cannot form another word. Such words are
// called ananagrams, an example is QUIZ.
// Obviously such definitions depend on the domain within which we are working; you might think
// that ATHENE is an ananagram, whereas any chemist would quickly produce ETHANE. One possible
// domain would be the entire English language, but this could lead to some problems. One could restrict
// the domain to, say, Music, in which case SCALE becomes a relative ananagram (LACES is not in the
// same domain) but NOTE is not since it can produce TONE.
// Write a program that will read in the dictionary of a restricted domain and determine the relative
// ananagrams. Note that single letter words are, ipso facto, relative ananagrams since they cannot be
// “rearranged” at all. The dictionary will contain no more than 1000 words.

// Input

// Input will consist of a series of lines. No line will be more than 80 characters long, but may contain any
// number of words. Words consist of up to 20 upper and/or lower case letters, and will not be broken
// across lines. Spaces may appear freely around words, and at least one space separates multiple words
// on the same line. Note that words that contain the same letters but of differing case are considered to
// be anagrams of each other, thus ‘tIeD’ and ‘EdiT’ are anagrams. The file will be terminated by a line
// consisting of a single ‘#’

// Output

// Output will consist of a series of lines. Each line will consist of a single word that is a relative ananagram
// in the input dictionary. Words must be output in lexicographic (case-sensitive) order. There will always
// be at least one relative ananagram.

// Sample Input

// ladder came tape soon leader acme RIDE lone Dreis peat
// ScAlE orb eye Rides dealer NotE derail LaCeS drIed
// noel dire Disk mace Rob dries
// #

// Sample Output

// Disk
// NotE
// derail
// drIed
// eye
// ladder
// soon

// easier with algorithm::sort

#include <algorithm>
#include <cctype>
#include <iostream>
#include <map>
#include <sstream>
#include <string>
#include <vector>

int main () {
    std::string line, word;
    std::vector<std::string> words, ananagrams;
    std::stringstream ss;
    std::map<std::string, unsigned> anagrams_counter;
    unsigned i;

    while (std::getline(std::cin, line)) {
        if (line == "#") {
            break;
        }

        ss = std::stringstream(line);

        while (ss >> word) {
            words.push_back(word);
        }
    }

    for (i = 0; i < words.size(); i++) {
        word = std::string(words[i]);
        std::transform(word.begin(), word.end(), word.begin(), tolower);
        std::sort(word.begin(), word.end());

        if (anagrams_counter.find(word) == anagrams_counter.end()) {
            anagrams_counter[word] = 1;
        } else {
            anagrams_counter[word]++;
        }
    }

    for (i = 0; i < words.size(); i++) {
        word = std::string(words[i]);
        std::transform(word.begin(), word.end(), word.begin(), tolower);
        std::sort(word.begin(), word.end());

        if (anagrams_counter[word] == 1) {
            ananagrams.push_back(words[i]);
        }
    }

    std::sort(ananagrams.begin(), ananagrams.end());

    for (i = 0; i < ananagrams.size(); i++) {
        std::cout << ananagrams[i] << std::endl;
    }

    return 0;
}
