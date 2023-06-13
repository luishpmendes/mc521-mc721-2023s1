// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=20&page=show_problem&problem=1742

// UVA 10801 - Lift Hopping
// Ted the bellhop: “I’m coming up and if there isn’t
// a dead body by the time I get there, I’ll make one
// myself. You!”
// Robert Rodriguez, "Four Rooms."
// A skyscraper has no more than 100 floors, numbered from 0 to 99. It has n (1 ≤ n ≤ 5) elevators
// which travel up and down at (possibly) different speeds. For each i in {1, 2, . . . n}, elevator number
// i takes Ti (1 ≤ Ti ≤ 100) seconds to travel between any two adjacent floors (going up or down).
// Elevators do not necessarily stop at every floor. What’s worse, not every floor is necessarily accessible
// by an elevator.
// You are on floor 0 and would like to get to floor k as quickly as possible. Assume that you do not
// need to wait to board the first elevator you step into and (for simplicity) the operation of switching an
// elevator on some floor always takes exactly a minute. Of course, both elevators have to stop at that
// floor. You are forbiden from using the staircase. No one else is in the elevator with you, so you don’t
// have to stop if you don’t want to. Calculate the minimum number of seconds required to get from floor
// 0 to floor k (passing floor k while inside an elevator that does not stop there does not count as “getting
// to floor k”).

// Input
// The input will consist of a number of test cases. Each test case will begin with two numbers, n and k,
// on a line. The next line will contain the numbers T1, T2, . . . Tn.
// Finally, the next n lines will contain sorted lists of integers – the first line will list the floors visited
// by elevator number 1, the next one will list the floors visited by elevator number 2, etc.

// Output
// For each test case, output one number on a line by itself - the minimum number of seconds required to
// get to floor k from floor 0. If it is impossible to do, print ‘IMPOSSIBLE’ instead.

// Explanation of examples
// In the first example, take elevator 1 to floor 13 (130 seconds), wait 60 seconds to switch to elevator
// 2 and ride it to floor 30 (85 seconds) for a total of 275 seconds.
// In the second example, take elevator 1 to floor 10, switch to elevator 2 and ride it until floor 25.
// There, switch back to elevator 1 and get off at the 30’th floor. The total time is 10 ∗ 10 + 60 + 15 ∗ 1 +
// 60 + 5 ∗ 10 = 285 seconds.
// In example 3, take elevator 1 to floor 30, then elevator 2 to floor 20 and then elevator 3 to floor 50.
// In the last example, the one elevator does not stop at floor 1.

// Sample Input
// 2 30
// 10 5
// 0 1 3 5 7 9 11 13 15 20 99
// 4 13 15 19 20 25 30
// 2 30
// 10 1
// 0 5 10 12 14 20 25 30
// 2 4 6 8 10 12 14 22 25 28 29
// 3 50
// 10 50 100
// 0 10 30 40
// 0 20 30
// 0 20 50
// 1 1
// 2
// 0 2 4 6 8 10

// Sample Output
// 275
// 285
// 3920
// IMPOSSIBLE

// model the graph carefully!

#include <iostream>
#include <limits>
#include <queue>
#include <string>
#include <sstream>
#include <utility>
#include <vector>

std::vector<int> dijkstra(const std::vector<std::vector<std::pair<unsigned, int>>> & adj, unsigned s) {
    std::vector<int> dist(adj.size(), std::numeric_limits<int>::max() >> 2);
    std::priority_queue<std::pair<int, unsigned>, std::vector<std::pair<int, unsigned>>, std::greater<std::pair<int, unsigned>>> pq;
    int d, w;
    unsigned u, v;
    dist[s] = 0;
    pq.push(std::make_pair(0, s));

    while (!pq.empty()) {
        d = pq.top().first;
        u = pq.top().second;
        pq.pop();

        if (d <= dist[u]) {

            for (const std::pair<unsigned, int> & edge : adj[u]) {
                v = edge.first;
                w = edge.second;

                if (dist[v] > d + w) {
                    dist[v] = d + w;
                    pq.push(std::make_pair(dist[v], v));
                }
            }
        }
    }

    return dist;
}

int main() {
    unsigned n, k, i, floor, u, v;
    std::vector<unsigned> t;
    std::vector<int> stops, dist;
    std::string s;
    std::stringstream ss;
    std::vector<std::vector<std::pair<unsigned, int>>> adj(100);
    int ans;

    while (std::cin >> n >> k) {
        t.resize(n);
        adj.assign(100, std::vector<std::pair<unsigned, int>>());

        for (unsigned & x : t) {
            std::cin >> x;
        }

        std::cin.ignore();

        for (i = 0; i < n; i++) {
            stops.clear();
            stops.reserve(100);
            std::getline(std::cin, s);
            ss.clear();
            ss.str(s);

            while (ss >> floor) {
                stops.push_back(floor);
            }

            for (u = 0; u < stops.size(); u++) {
                for (v = u + 1; v < stops.size(); v++) {
                    adj[stops[u]].push_back(std::make_pair(stops[v], (stops[v] - stops[u]) * t[i] + 60));
                    adj[stops[v]].push_back(std::make_pair(stops[u], (stops[v] - stops[u]) * t[i] + 60));
                }
            }
        }

        dist = dijkstra(adj, 0);

        if (dist[k] >= std::numeric_limits<int>::max() >> 2) {
            std::cout << "IMPOSSIBLE" << std::endl;
        } else {
            ans = dist[k];

            if (k != 0) {
                ans -= 60;
            }

            std::cout << ans << std::endl;
        }
    }
	
	return 0;
}
