// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=11&page=show_problem&problem=866

// UVA 925 - No more prerequisites, please!
// Recently, I had to prepare a small students guide concerning our Computer Science graduation. I asked
// all my colleagues to send me some specific pieces of information about the courses they were responsible
// for. One of these pieces of information was the set of prerequisites that is defined for each course c,
// that is, the set of courses that the student is supposed to have completed before he can attend course
// c.
// Some of my colleagues were so zealous of their task that they sent me more information than I
// needed . . . I’ll explain:
// Let ca, cb, cc, cd and ce be courses of our Computer Science graduation.
// Suppose ca is a prerequisite for cb; cb and cc are prerequisites for cd; cd is a prerequisite
// for ce.
// John, who is responsible for course cd, told me that the prerequisites for cd were cc, cb and
// ca;
// Elizabeth, who is responsible for course ce, told me that the prerequisites for ce were cd,
// cc, cb and ca.
// But I only needed prerequisites cc and cb for course cd because cb already has ca as its prerequisite!
// Likewise, I only needed prerequisite cd for course ce because cd already has cc, cb, and ca as its
// prerequisites!
// Unfortunately, not all my colleagues asked me whether I wanted the whole set of prerequisites or
// the minimum one; they asked John instead!! He told them that “the most information, the better”!
// I had no courage to tell them that, by sending me more information than the one I needed, they
// were turning my task of gathering all their info more difficult. How I wished for some robot that I could
// feed with the information I had, and that gave me, for each course, the minimum set of prerequisites!
// Never too late. . .
// Your task consists of writing a program that, given the name and prerequisites of a set of courses,
// computes the minimum set of prerequisites for all courses.

// Input
// The first line contains the number C of test cases that follow (0 < C < 1000).
// Each test case starts with a line containing the number k (0 < k ≤ 120) of courses that are to be
// processed. The next k lines contain the names of the k courses, one per line. The next line contains the
// number j of courses that have prerequisites (0 ≤ j ≤ 120). The next j lines contain, for each course
// that has some prerequisite, the course name, the number of prerequisites it has, and the names of the
// courses that are its prerequisites, separated by a single space.
// Course names have maximum length 20, do not contain any spaces, and are composed of characters
// in the set {‘a’..‘z’}.
// There are no cycles on prerequisites, that is, it is never the case that some course c has a prerequisite
// course cp that has c as a prerequisite (directly or indirectly).

// Output
// The first line contains the number C of test cases that follow (0 < C < 1000).
// Each test case starts with a line containing the number k (0 < k ≤ 120) of courses that are to be
// processed. The next k lines contain the names of the k courses, one per line. The next line contains the
// number j of courses that have prerequisites (0 ≤ j ≤ 120). The next j lines contain, for each course
// that has some prerequisite, the course name, the number of prerequisites it has, and the names of the
// courses that are its prerequisites, separated by a single space.
// Course names have maximum length 20, do not contain any spaces, and are composed of characters
// in the set {‘a’..‘z’}.
// There are no cycles on prerequisites, that is, it is never the case that some course c has a prerequisite
// course cp that has c as a prerequisite (directly or indirectly).

// Sample Input
// 2
// 5
// cc
// ca
// ce
// cb
// cd
// 3
// ce 4 cd cb cc ca
// cd 3 cb cc ca
// cb 1 ca
// 9
// dg
// di
// dc
// df
// de
// dd
// da
// dh
// db
// 7
// dg 3 da de df
// dd 2 da db
// df 1 de
// dc 1 db
// dh 5 da de dg df dc
// de 2 da dd
// di 2 df dd

// Sample Output
// cb 1 ca
// cd 2 cb cc
// ce 1 cd
// dc 1 db
// dd 2 da db
// de 1 dd
// df 1 de
// dg 1 df
// dh 2 dc dg
// di 1 df

// transitive closure++

#include <algorithm>
#include <iostream>
#include <map>
#include <string>
#include <vector>

int main() {
    unsigned t, n, m, c, i, j, k;
    std::vector<std::string> id_to_name;
    std::map<std::string, unsigned> name_to_id;
    std::vector<std::vector<unsigned>> adj;
    std::string u, v;
    std::vector<unsigned> result;
    
    while (std::cin >> t) {
        while (t--) {
            std::cin >> n;

            id_to_name.resize(n);
            name_to_id.clear();
            adj.resize(n);
            adj.assign(n, std::vector<unsigned>(n, 0));

            for (i = 0; i < n; i++) {
                std::cin >> id_to_name[i];
            }

            std::sort(id_to_name.begin(), id_to_name.end());

            for (i = 0; i < n; i++) {
                name_to_id[id_to_name[i]] = i;
            }

            std::cin >> m;

            while (m--) {
                std::cin >> u >> c;

                while (c--) {
                    std::cin >> v;
                    adj[name_to_id[u]][name_to_id[v]] = 1;
                }
            }

            for (k = 0; k < n; k++) {
                for (i = 0; i < n; i++) {
                    for (j = 0; j < n; j++) {
                        if (adj[i][k] != 0 && adj[k][j] != 0) {
                            if (adj[i][j] < adj[i][k] + adj[k][j]) {
                                adj[i][j] = adj[i][k] + adj[k][j];
                            }
                        }
                    }
                }
            }

            for (i = 0; i < n; i++) {
                result.clear();

                for (j = 0; j < n; j++) {
                    if (adj[i][j] == 1) {
                        result.push_back(j);
                    }
                }

                if (result.size() > 0) {
                    std::cout << id_to_name[i] << " " << result.size();

                    for (j = 0; j < result.size(); j++) {
                        std::cout << " " << id_to_name[result[j]];
                    }

                    std::cout << std::endl;
                }
            }
        }
    }
	
	return 0;
}
