// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=229&page=show_problem&problem=3053

// UVA 11902 - Dominator

// In graph theory, a node X dominates a node Y if every path
// from the predefined start node to Y must go through X. If Y
// is not reachable from the start node then node Y does not have
// any dominator. By definition, every node reachable from the
// start node dominates itself. In this problem, you will be given a
// directed graph and you have to find the dominators of every node
// where the 0-th node is the start node.
// As an example, for the graph shown right, 3 dominates 4 since
// all the paths from 0 to 4 must pass through 3. 1 doesn’t dominate
// 3 since there is a path 0-2-3 that doesn’t include 1.

// Input

// The first line of input will contain T (≤ 100) denoting the number
// of cases.
// Each case starts with an integer N (0 < N < 100) that represents the number of nodes in the graph. The next N lines contain N integers each. If the j-th (0
// based) integer of i-th (0 based) line is ‘1’, it means that there is an edge from node i to node j and
// similarly a ‘0’ means there is no edge.

// Output

// For each case, output the case number first. Then output 2N + 1 lines that summarizes the dominator
// relationship between every pair of nodes. If node A dominates node B, output ‘Y’ in cell (A, B),
// otherwise output ‘N’. Cell (A, B) means cell at A-th row and B-th column. Surround the output with
// ‘|’, ‘+’ and ‘-’ to make it more legible. Look at the samples for exact format.

// Sample Input

// 2
// 5
// 0 1 1 0 0
// 0 0 0 1 0
// 0 0 0 1 0
// 0 0 0 0 1
// 0 0 0 0 0
// 1
// 1

// Sample Output

// Case 1:
// +---------+
// |Y|Y|Y|Y|Y|
// +---------+
// |N|Y|N|N|N|
// +---------+
// |N|N|Y|N|N|
// +---------+
// |N|N|N|Y|Y|
// +---------+
// |N|N|N|N|Y|
// +---------+
// Case 2:
// +-+
// |Y|
// +-+

// disable vertex one by one, check if the reachability from vertex 0 changes

#include <iostream>
#include <vector>

void dfs(const std::vector<std::vector<unsigned>> & adj,
         unsigned u,
         unsigned excluded,
         std::vector<bool> & visited) {
    unsigned v;

    if (u != excluded) {
        visited[u] = true;

        for (v = 0; v < adj[u].size(); v++) {
            if (adj[u][v] && !visited[v]) {
                dfs(adj, v, excluded, visited);
            }
        }
    }
}

int main() {
    unsigned t, c, n, u, v;
    std::vector<std::vector<unsigned>> adj;
    std::vector<bool> visited_from_start, visited;
    std::vector<std::vector<bool>> ans;
    std::string separator;

    while (std::cin >> t) {
        for (c = 1; c <= t; c++) {
            std::cin >> n;
            adj.resize(n);
            adj.assign(n, std::vector<unsigned>(n, 0));
            visited_from_start.resize(n);
            visited_from_start.assign(n, false);
            visited.resize(n);
            ans.resize(n);
            ans.assign(n, std::vector<bool>(n, false));

            for (u = 0; u < n; u++) {
                for (v = 0; v < n; v++) {
                    std::cin >> adj[u][v];
                }
            }

            dfs(adj, 0, n, visited_from_start);

            for (u = 0; u < n; u++) {
                if (visited_from_start[u]) {
                    visited.assign(n, false);

                    dfs(adj, 0, u, visited);

                    for (v = 0; v < n; v++) {
                        if (visited_from_start[v] && !visited[v]) {
                            ans[u][v] = true;
                        }
                    }
                }
            }

            std::cout << "Case " << c << ":" << std::endl;
            separator = "+" + std::string(2 * n - 1, '-') + "+";
            std::cout << separator << std::endl;

            for (u = 0; u < n; u++) {
                for (v = 0; v < n; v++) {
                    std::cout << "|";

                    if (ans[u][v]) {
                        std::cout << "Y";
                    } else {
                        std::cout << "N";
                    }
                }

                std::cout << "|" << std::endl;
                std::cout << separator << std::endl;
            }
        }
    }

    return 0;
}
