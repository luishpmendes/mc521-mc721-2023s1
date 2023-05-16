// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=5&page=show_problem&problem=251

// UVA 315 - Network

// A Telephone Line Company (TLC) is establishing a new telephone cable network. They are connecting
// several places numbered by integers from 1 to N. No two places have the same number. The lines
// are bidirectional and always connect together two places and in each place the lines end in a telephone
// exchange. There is one telephone exchange in each place. From each place it is possible to reach through
// lines every other place, however it need not be a direct connection, it can go through several exchanges.
// From time to time the power supply fails at a place and then the exchange does not operate. The
// officials from TLC realized that in such a case it can happen that besides the fact that the place with
// the failure is unreachable, this can also cause that some other places cannot connect to each other. In
// such a case we will say the place (where the failure occured) is critical. Now the officials are trying to
// write a program for finding the number of all such critical places. Help them.

// Input

// The input file consists of several blocks of lines. Each block describes one network. In the first line
// of each block there is the number of places N < 100. Each of the next at most N lines contains the
// number of a place followed by the numbers of some places to which there is a direct line from this place.
// These at most N lines completely describe the network, i.e., each direct connection of two places in the
// network is contained at least in one row. All numbers in one line are separated by one space. Each
// block ends with a line containing just ‘0’. The last block has only one line with N = 0.

// Output

// The output contains for each block except the last in the input file one line containing the number of
// critical places.

// Sample Input

// 5
// 5 1 2 3 4
// 0
// 6
// 2 1 3
// 5 4 6 2
// 0
// 0

// Sample Output

// 1
// 2

// finding articulation points

#include <iostream>
#include <set>
#include <sstream>
#include <string>
#include <vector>

void articulation_point(const std::vector<std::vector<int>> & adj_list,
                        std::vector<int> & dfs_num,
                        std::vector<int> & dfs_low,
                        std::vector<int> & dfs_parent,
                        int & dfs_number_counter,
                        const int & dfs_root,
                        int & root_children,
                        int u,
                        std::set<int> & articulation_vertices) {
    dfs_low[u] = dfs_num[u] = dfs_number_counter++; // dfs_low[u] <= dfs_num[u]

    for (const int & v : adj_list[u]) {
        if (dfs_num[v] == -1) { // a tree edge
            dfs_parent[v] = u;

            if (u == dfs_root) {
                // special case if u is a root
                root_children++;
            }

            articulation_point(adj_list,
                               dfs_num,
                               dfs_low,
                               dfs_parent,
                               dfs_number_counter,
                               dfs_root,
                               root_children,
                               v,
                               articulation_vertices);

            if (dfs_low[v] >= dfs_num[u]) { // for articulation point
                articulation_vertices.insert(u); // store this information first
            }

            if (dfs_low[u] > dfs_low[v]) {
                // update dfs_low[u]
                dfs_low[u] = dfs_low[v];
            }
        } else if (v != dfs_parent[u]) { // a back edge and not direct cycle
            if (dfs_low[u] > dfs_num[v]) {
                // update dfs_low[u]
                dfs_low[u] = dfs_num[v];
            }
        }
    }
}

int main() {
    int n, u, v;
    std::vector<std::vector<int>> adj_list;
    std::string s;
    std::stringstream ss;
    std::vector<int> dfs_num, dfs_low, dfs_parent;
    int dfs_number_counter, dfs_root, root_children;
    std::set<int> articulation_vertices;

    while (std::cin >> n && n) {
        adj_list.resize(n);
        adj_list.assign(n, std::vector<int>());
        dfs_num.resize(n);
        dfs_num.assign(n, -1);
        dfs_low.resize(n);
        dfs_low.assign(n, 0);
        dfs_parent.resize(n);
        dfs_parent.assign(n, 0);
        articulation_vertices.clear();
        dfs_number_counter = 0;
        std::cin.ignore();

        while (std::getline(std::cin, s) && s != "0") {
            ss = std::stringstream(s);
            ss >> u;

            while (ss >> v) {
                adj_list[u - 1].push_back(v - 1);
                adj_list[v - 1].push_back(u - 1);
            }
        }

        for (u = 0; u < n; u++) {
            if (dfs_num[u] == -1) {
                dfs_root = u;
                root_children = 0;
                articulation_point(adj_list,
                                   dfs_num,
                                   dfs_low,
                                   dfs_parent,
                                   dfs_number_counter,
                                   dfs_root,
                                   root_children,
                                   u,
                                   articulation_vertices);

                // special case
                if (root_children > 1) {
                    articulation_vertices.insert(dfs_root);
                } else {
                    articulation_vertices.erase(dfs_root);
                }
            }
        }

        std::cout << articulation_vertices.size() << std::endl;
    }

    return 0;
}
