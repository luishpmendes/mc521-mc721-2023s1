// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&page=show_problem&problem=3347

// UVA 12195 - Jingle Composing

// A. C. Marcos is taking his first steps in the direction of jingle composition. He is having some troubles,
// but at least he is achieving pleasant melodies and attractive rhythms.
// In music, a note has a pitch (its frequency, resulting in how high or low is the sound) and a duration
// (for how long the note should sound). In this problem we are interested only in the duration of the
// notes.
// A jingle is divided into a sequence of measures, and a measure is formed by a series of notes.
// The duration of a note is indicated by its shape. In this problem, we will use uppercase letters to
// indicate a note’s duration. The following table lists all the available notes:
// The duration of a measure is the sum of the durations of its notes. In Marcos’ jingles, each measure
// has the same duration. As Marcos is just a beginner, his famous teacher Johann Sebastian III taught
// him that the duration of a measure must always be 1.
// For example, Marcos wrote a composition containing five measures, of which the first four have the
// correct duration and the last one is wrong. In the example below, each measure is surrounded with
// slashes and each note is represented as in the table above.
// /HH/QQQQ/XXXTXTEQH/W/HW/
// Marcos likes computers as much as music. He wants you to write a program that determines, for
// each one of his compositions, how many measures have the right duration.

// Input

// The input contains several test cases. Each test case is described in a single line containing a string
// whose length is between 3 and 200 characters, inclusive, representing a composition. A composition
// begins and ends with a slash ‘/’. Measures in a composition are separated by a slash ‘/’. Each note in
// a measure is represented by the corresponding uppercase letter, as described above. You may assume
// that each composition contains at least one measure and that each measure contains at least one note.
// All characters in the input will be either slashes or one of the seven uppercase letters used to represent
// notes, as described above.
// The last test case is followed by a line containing a single asterisk.

// Output

// For each test case your program must output a single line, containing a single integer, the number of
// measures that have the right duration.

// Sample Input

// /HH/QQQQ/XXXTXTEQH/W/HW/
// /W/W/SQHES/
// /WE/TEX/THES/
// *

// Sample Output

// 4
// 3
// 0

// count the number of correct measures

#include <iostream>
#include <string>

int main () {
    std::string s;
    unsigned sum, ans;

    while (std::cin >> s && s != "*") {
        sum = 0;
        ans = 0;

        for (const char c : s) {
            if (c == '/') {
                if (sum == 64) {
                    ans++;
                }

                sum = 0;
            } else if (c == 'W') {
                sum += 64;
            } else if (c == 'H') {
                sum += 32;
            } else if (c == 'Q') {
                sum += 16;
            } else if (c == 'E') {
                sum += 8;
            } else if (c == 'S') {
                sum += 4;
            } else if (c == 'T') {
                sum += 2;
            } else if (c == 'X') {
                sum += 1;
            }
        }

        std::cout << ans << std::endl;
    }

    return 0;
}
