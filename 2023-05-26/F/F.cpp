// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=441&page=show_problem&problem=4027

// UVA 12582 - Wedding of Sultan

// As usual Sultan Mahmud is very busy. He works days and nights at office. If you ask him, “Sultan,
// which day of the week is this?” He will look at you for a while and say, “I think I have 3 more days
// till deadline!” But one day the scenario changed after receiving a call. He usually ignores phone calls
// from everyone (even from his fiance) but this time he couldn’t ignore because of the importance of the
// person! This person was his to-be mother-in-law. So he received the call and heard, “Son, only 30
// days left of your wedding ceremony, so I am sending a tailor for the measurement for your suit.” Sultan
// now remembered, he is about to get married but looking at thyself, he got surprised! When did he get
// so fat! “Umm.. Mom can it be arranged 10 days later?” He wants to buy some time so that he can
// exercise and lose extra weight. So he immediately went out with his bicycle to the large garden beside
// his house.
// There are several trails in the garden. A trail starts from one water sprinkler to another and the
// sprinkler are marked by distinct letters from ‘A’ to ‘Z’. The trails are designed in such a way that from
// the sprinkler at entrance you can go to any other sprinkler using exactly one path if you do not traverse
// a trail more than once.
// While traversing the trails with his cycle, Sultan notes the names of the sprinklers in his notepad.
// He will write down the name of a sprinkler if he enters the sprinkler for the first time or leaves this
// sprinkler for the last time. And not surprisingly, geek Sultan follows a peculiar method to ensure that
// he visits all the trails of the garden. When he comes to a sprinkler he looks for a trail which he has not
// traversed yet. If he finds such trail, he follows that one. Otherwise, he uses the trail that he used to
// come here for the first time except if it’s the entrance he stops exercising. He always starts from the
// entrance and guess what, his peculiar strategy always guarantees to finish him at the entrance and all
// the trails are also visited.
// For example, in the above garden the main entrance is at A. So Sultan will start from A. When
// Sultan is at A, he can choose either of the trails. Say he chooses the trail leading to E. Then he can
// choose the trail to G or trail to F. Say he chooses F. Now he does not have any unvisited trail from F,
// so he will go back to E. Now he must choose trail to G and then similarly will come back to E and back
// to A. Then he will go towards B. Now he again has two choices. He can go to C or D, say he goes to
// C, then he will be back to B, then will go to D, and hence back to B and also back to A thus finishing
// his exercise. So after his exercise you will find: AEFFGGEBDDCCBA in his notepad. Can you find
// the number of trails attached to the sprinklers just looking at the sequence written in the notepad?

// Input

// First line of the test file contains a positive integer T (T ≤ 100) denoting the number of test cases.
// Hence follows T lines, each containing a valid sequence of sprinkler names. A sprinkler name will
// always be capital Latin letter (‘A’, ‘B’, . . . , ‘Z’). You may assume that there will be at least two sprinklers
// in garden, otherwise there would have been no meaning of exercise right?

// Output

// For each case output the case number in the first line, followed by the number of trails for each sprinkler.
// First print the sprinkler name followed by the count of trails. These lines should be in lexicographical
// order of sprinkler name. Note that, you should not print about a sprinkler which is not present in the
// garden. Look at the sample input output for more specific format of input output.

// Sample Input

// 2
// AEFFGGEBDDCCBA
// ZAABBZ

// Sample Output

// Case 1
// A = 2
// B = 3
// C = 1
// D = 1
// E = 3
// F = 1
// G = 1
// Case 2
// A = 1
// B = 1
// Z = 2

// given graph DFS traversal, count the degree of each vertex

#include <iostream>
#include <map>
#include <stack>
#include <string>

int main() {
    unsigned t, c, i;
    std::string str;
    std::stack<char> s;
    std::map<char, unsigned> degree;
    
    while (std::cin >> t) {
        for (c = 1; c <= t; c++) {
            std::cin >> str;
            s = std::stack<char>();
            degree.clear();

            s.push(str[0]);

            for (i = 1; i < str.size(); i++) {
                if (s.top() == str[i]) {
                    s.pop();
                } else {
                    degree[s.top()]++;
                    degree[str[i]]++;
                    s.push(str[i]);
                }
            }

            std::cout << "Case " << c << "\n";

            for (std::map<char, unsigned>::iterator it = degree.begin();
                 it != degree.end();
                 it++) {
                std::cout << it->first << " = " << it->second << "\n";
            }
        }
    }

    return 0;
}
