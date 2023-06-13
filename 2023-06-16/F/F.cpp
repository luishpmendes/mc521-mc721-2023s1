// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=26&page=show_problem&problem=2458

// UVA 11463 - Commandos
// A group of commandos were assigned a critical task. They are to destroy an enemy head quarter.
// The enemy head quarter consists of several buildings and the buildings are connected by roads. The
// commandos must visit each building and place a bomb at the base of each building. They start their
// mission at the base of a particular building and from there they disseminate to reach each building.
// The commandos must use the available roads to travel between buildings. Any of them can visit one
// building after another, but they must all gather at a common place when their task in done.
// In this problem, you will be given the description of different enemy headquarters. Your job is to
// determine the minimum time needed to complete the mission. Each commando takes exactly one unit
// of time to move between buildings.
// You may assume that the time required to place a bomb is negligible. Each commando can carry
// unlimited number of bombs and there is an unlimited supply of commando troops for the mission.

// Input
// The first line of input contains a number T < 50, where T denotes the number of test cases.
// Each case describes one head quarter scenario. The first line of each case starts with a positive
// integer N ≤ 100, where N denotes the number of buildings in the head quarter. The next line contains
// a positive integer R, where R is the number of roads connecting two buildings. Each of the next R
// lines contain two distinct numbers, 0 ≤ u, v < N, this means there is a road connecting building u
// to building v. The buildings are numbered from 0 to N − 1. The last line of each case contains two
// integers 0 ≤ s, d < N. Where s denotes the building from where the mission starts and d denotes the
// building where they must meet.
// You may assume that two buildings will be directly connected by at most one road. The input will
// be such that, it will be possible to go from any building to another by using one or more roads.

// Output
// For each case of input, there will be one line of output. It will contain the case number followed by the
// minimum time required to complete the mission. Look at the sample output for exact formatting.

// Sample Input
// 2
// 4
// 3
// 0 1
// 2 1
// 1 3
// 0 3
// 2
// 1
// 0 1
// 1 0

// Sample Output
// Case 1: 4
// Case 2: 1

// solution is easy with APSP information

#include <iostream>
#include <limits>
#include <vector>

int main() {
    unsigned t, c, n, r, u, v, s, d, w, ans;
    std::vector<std::vector<unsigned>> adj;

    while (std::cin >> t) {
        for (c = 1; c <= t; c++) {
            std::cin >> n >> r;

            adj.resize(n);
            adj.assign(n, std::vector<unsigned>(n, std::numeric_limits<unsigned>::max() >> 2));

            for (v = 0; v < n; v++) {
                adj[v][v] = 0;
            }

            while (r--) {
                std::cin >> u >> v;
                adj[u][v] = adj[v][u] = 1;
            }

            std::cin >> s >> d;

            for (w = 0; w < n; w++) {
                for (u = 0; u < n; u++) {
                    for (v = 0; v < n; v++) {
                        if (adj[u][v] > adj[u][w] + adj[w][v]) {
                            adj[u][v] = adj[u][w] + adj[w][v];
                        }
                    }
                }
            }

            ans = 0;

            for (v = 0; v < n; v++) {
                if (ans < adj[s][v] + adj[v][d]) {
                    ans = adj[s][v] + adj[v][d];
                }
            }

            std::cout << "Case " << c << ": " << ans << std::endl;
        }
    }
	
	return 0;
}
