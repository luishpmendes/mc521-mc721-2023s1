// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=4&page=show_problem&problem=183

// UVA 247 - Calling Circles

// If you’ve seen television commercials for long-distance phone companies lately, you’ve noticed that
// many companies have been spending a lot of money trying to convince people that they provide the
// best service at the lowest cost. One company has “calling circles.” You provide a list of people that
// you call most frequently. If you call someone in your calling circle (who is also a customer of the same
// company), you get bigger discounts than if you call outside your circle. Another company points out
// that you only get the big discounts for people in your calling circle, and if you change who you call
// most frequently, it’s up to you to add them to your calling circle.
// LibertyBell Phone Co. is a new company that thinks they have the calling plan that can put other
// companies out of business. LibertyBell has calling circles, but they figure out your calling circle for you.
// This is how it works. LibertyBell keeps track of all phone calls. In addition to yourself, your calling
// circle consists of all people whom you call and who call you, either directly or indirectly.
// For example, if Ben calls Alexander, Alexander calls Dolly, and Dolly calls Ben, they are all within
// the same circle. If Dolly also calls Benedict and Benedict calls Dolly, then Benedict is in the same
// calling circle as Dolly, Ben, and Alexander. Finally, if Alexander calls Aaron but Aaron doesn’t call
// Alexander, Ben, Dolly, or Benedict, then Aaron is not in the circle.
// You’ve been hired by LibertyBell to write the program to determine calling circles given a log of
// phone calls between people.

// Input

// The input file will contain one or more data sets. Each data set begins with a line containing two
// integers, n and m. The first integer, n, represents the number of different people who are in the data
// set. The maximum value for n is 25. The remainder of the data set consists of m lines, each representing
// a phone call. Each call is represented by two names, separated by a single space. Names are first names
// only (unique within a data set), are case sensitive, and consist of only alphabetic characters; no name
// is longer than 25 letters.
// For example, if Ben called Dolly, it would be represented in the data file as
// Ben Dolly
// Input is terminated by values of zero (0) for n and m.

// Output

// For each input set, print a header line with the data set number, followed by a line for each calling
// circle in that data set. Each calling circle line contains the names of all the people in any order within
// the circle, separated by comma-space (a comma followed by a space). Output sets are separated by
// blank lines.

// Sample Input

// 5 6
// Ben Alexander
// Alexander Dolly
// Dolly Ben
// Dolly Benedict
// Benedict Dolly
// Alexander Aaron
// 14 34
// John Aaron
// Aaron Benedict
// Betsy John
// Betsy Ringo
// Ringo Dolly
// Benedict Paul
// John Betsy
// John Aaron
// Benedict George
// Dolly Ringo
// Paul Martha
// George Ben
// Alexander George
// Betsy Ringo
// Alexander Stephen
// Martha Stephen
// Benedict Alexander
// Stephen Paul
// Betsy Ringo
// Quincy Martha
// Ben Patrick
// Betsy Ringo
// Patrick Stephen
// Paul Alexander
// Patrick Ben
// Stephen Quincy
// Ringo Betsy
// Betsy Benedict
// Betsy Benedict
// Betsy Benedict
// Betsy Benedict
// Betsy Benedict
// Betsy Benedict
// Quincy Martha
// 0 0

// Sample Output

// Calling circles for data set 1:
// Ben, Alexander, Dolly, Benedict
// Aaron
// 
// Calling circles for data set 2:
// John, Betsy, Ringo, Dolly
// Aaron
// Benedict
// Paul, George, Martha, Ben, Alexander, Stephen, Quincy, Patrick

// SCC + printing solution

#include <iostream>
#include <map>
#include <queue>
#include <set>
#include <stack>
#include <string>
#include <vector>

void tarjan_aux(
        const std::vector<std::vector<unsigned>> & adj_list,
        std::stack<unsigned> & S,
        std::vector <bool> & in_S,
        unsigned & i,
        std::vector <int> & index,
        std::vector <int> & low_link,
        unsigned u,
        std::set<std::set<unsigned>> & scc) {
    index[u] = i;
    low_link[u] = i;
    i++;
    S.push(u);
    in_S[u] = true;

    for (const unsigned & v : adj_list[u]) {
        if (index[v] < 0) {
            tarjan_aux(adj_list, S, in_S, i, index, low_link, v, scc);

            if (low_link[u] > low_link[v]) {
                low_link[u] = low_link[v];
            }
        } else if (in_S[v]) {
            if (low_link[u] > low_link[v]) {
                low_link[u] = low_link[v];
            }
        }
    }

    if (low_link[u] == index[u]) {
        unsigned v = adj_list.size() + 1;
        std::set <unsigned> component;
        do {
            v = S.top();
            S.pop();
            in_S[v] = false;
            component.insert(v);
        } while (u != v);

        scc.insert(component);
    }
}

std::set<std::set<unsigned>> tarjan(
        const std::vector<std::vector<unsigned>> & adj_list) {
    std::set<std::set<unsigned>> result;
    std::stack<unsigned> S;
    std::vector<bool> in_S(adj_list.size(), false);
    unsigned i = 0;
    std::vector<int> index (adj_list.size(), -1),
                     low_link (adj_list.size(), -1);
    
    for (unsigned u = 0; u < adj_list.size(); u++) {
        if (index[u] < 0) {
            tarjan_aux(adj_list, S, in_S, i, index, low_link, u, result);
        }
    }
    return result;
}

int main() {
    unsigned n, m, t;
    std::vector<std::vector<unsigned>> adj_list;
    std::map<std::string, unsigned> name_to_id;
    std::vector<std::string> id_to_name;
    std::string s1, s2;
    std::set<std::set<unsigned>> scc;

    t = 0;

    while (std::cin >> n >> m && (m || n)) {
        adj_list.resize(n);
        adj_list.assign(n, std::vector<unsigned>());
        name_to_id.clear();
        id_to_name.clear();

        while (m--) {
            std::cin >> s1 >> s2;

            if (name_to_id.find(s1) == name_to_id.end()) {
                name_to_id[s1] = name_to_id.size();
                id_to_name.push_back(s1);
            }

            if (name_to_id.find(s2) == name_to_id.end()) {
                name_to_id[s2] = name_to_id.size();
                id_to_name.push_back(s2);
            }

            adj_list[name_to_id[s1]].push_back(name_to_id[s2]);
        }

        scc = tarjan(adj_list);

        if (t) {
            std::cout << std::endl;
        }

        std::cout << "Calling circles for data set " << ++t << ":" << std::endl;

        for (const std::set<unsigned> & component : scc) {
            for (const unsigned & u : component) {
                std::cout << id_to_name[u];

                if (u != *component.rbegin()) {
                    std::cout << ", ";
                }
            }
            std::cout << std::endl;
        }
    }

    return 0;
}
