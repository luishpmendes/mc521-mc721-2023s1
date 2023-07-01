// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=4&page=show_problem&problem=195

// UVA 259 - Software Allocation
// A computing center has ten different computers (numbered 0 to 9) on which applications can run. The
// computers are not multi-tasking, so each machine can run only one application at any time. There are
// 26 applications, named A to Z. Whether an application can run on a particular computer can be found
// in a job description (see below).
// Every morning, the users bring in their applications for that day. It is possible that two users bring
// in the same application; in that case two different, independent computers will be allocated for that
// application.
// A clerk collects the applications, and for each different application he makes a list of computers on
// which the application could run. Then, he assigns each application to a computer. Remember: the
// computers are not multi-tasking, so each computer must handle at most one application in total. (An
// application takes a day to complete, so that sequencing i.e. one application after another on the same
// machine is not possible.)
// A job description consists of
// 1. one upper case letter A...Z, indicating the application.
// 2. one digit 1...9, indicating the number of users who brought in the application.
// 3. a blank (space character.)
// 4. one or more different digits 0...9, indicating the computers on which the application can run.
// 5. a terminating semicolon ‘;’.
// 6. an end-of-line.

// Input
// The input for your program is a textfile. For each day it contains one or more job descriptions, separated
// by a line containing only the end-of-line marker. The input file ends with the standard end-of-file marker.
// For each day your program determines whether an allocation of applications to computers can be done,
// and if so, generates a possible allocation.

// Output
// The output is also a textfile. For each day it consists of one of the following:
// • ten characters from the set {‘A’ ... ‘Z’, ‘_’}, indicating the applications allocated to computers
// 0 to 9 respectively if an allocation was possible. An underscore ‘_’ means that no application is
// allocated to the corresponding computer.
// • a single character ‘!’, if no allocation was possible.

// Sample Input
// A4 01234;
// Q1 5;
// P4 56789;
// 
// A4 01234;
// Q1 5;
// P5 56789;

// Sample Output
// AAAA_QPPPP
// !

// With the given Edmonds Karp’s code above, solving a (basic/standard) Network Flow problem, especially Max Flow, is now simpler. It is now a matter of:
// 1. Recognizing that the problem is indeed a Network Flow problem
// (this will get better after you solve more Network Flow problems).
// 2. Constructing the appropriate flow graph (i.e. if using our code shown earlier:
// Initiate the residual matrix res and set the appropriate values for ‘s’ and ‘t’).
// 3. Running the Edmonds Karp’s code on this flow graph.
// In this subsection, we show an example of modeling the flow (residual) graph of UVa 259 -
// Software Allocation17. The abridged version of this problem is as follows: You are given up
// to 26 applications/apps (labeled ‘A’ to ‘Z’), up to 10 computers (numbered from 0 to 9),
// the number of persons who brought in each application that day (one digit positive integer,
// or [1..9]), the list of computers that a particular application can run, and the fact that
// each computer can only run one application that day. Your task is to determine whether an
// allocation (that is, a matching) of applications to computers can be done, and if so, generates
// a possible allocation. If no, simply print an exclamation mark ‘!’.
// One (bipartite) flow graph
// formulation is shown in Figure 4.27. We index the vertices from [0..37] as there
// are 26 + 10 + 2 special vertices = 38 vertices. The
// source s is given index 0, the
// 26 possible apps are given indices from [1..26], the 10
// possible computers are given
// indices from [27..36], and
// finally the sink t is given index 37.
// Then, we link apps to computers as mentioned in the problem description. We link source
// s to all apps and link all computers to sink t. All edges in this flow graph are directed edges.
// The problem says that there can be more than one (say, X) users bringing in a particular
// app A in a given day. Thus, we set the edge weight (capacity) from source s to a particular
// app A to X. The problem also says that each computer can only be used once. Thus, we
// set the edge weight from each computer B to sink t to 1. The edge weight between apps to
// computers is set to ∞. With this arrangement, if there is a flow from an app A to a computer
// B and finally to sink t, that flow corresponds to one matching between that particular app
// A and computer B.
// Once we have this flow graph, we can pass it to our Edmonds Karp’s implementation
// shown earlier to obtain the Max Flow mf. If mf is equal to the number of applications
// brought in that day, then we have a solution, i.e. if we have X users bringing in app A, then
// X different paths (i.e. matchings) from A to sink t must be found by the Edmonds Karp’s
// algorithm (and similarly for the other apps).
// The actual app → computer assignments can be found by simply checking the backward
// edges from computers (vertices 27 - 36) to apps (vertices 1 - 26). A backward edge (computer
// → app) in the residual matrix res will contain a value +1 if the corresponding forward edge
// (app → computer) is selected in the paths that contribute to the Max Flow mf. This is also
// the reason why we start the flow graph with directed edges from apps to computers only

#include <iostream>
#include <limits>
#include <queue>
#include <string>
#include <utility>
#include <vector>

void augment(std::vector<std::vector<unsigned>> & res,
             const std::vector<unsigned> & p,
             const unsigned & s,
             unsigned & f,
             unsigned v,
             unsigned minEdge) {
    // traverse BFS spanning tree from s->t
    if (v == s) {
        // record minEdge in a global var f
        f = minEdge;
    } else if (p[v] != p.size()) {
        augment(res, p, s, f, p[v], std::min(minEdge, res[p[v]][v]));

        res[p[v]][v] -= f;
        res[v][p[v]] += f;
    }
}

unsigned edmondsKarp(const std::vector<std::vector<unsigned>> & adj,
                     std::vector<std::vector<unsigned>> & res,
                     const unsigned & s,
                     const unsigned & t) {
    // mf stands for max_flow
    unsigned mf = 0, f, u;
    std::vector<bool> visited (adj.size(), false);
    std::vector<unsigned> p (adj.size(), adj.size());
    std::queue<unsigned> q;

    // O(VE^2) Edmonds Karp’s algorithm
    while(true) {
        f = 0;
        visited.assign(adj.size(), false);
        p.assign(adj.size(), adj.size());
        visited[s] = true;
        q.push(s);

        while (!q.empty()) {
            u = q.front();
            q.pop();

            if (u == t) {
                // immediately stop BFS if we already reach sink t
                while (!q.empty()) {
                    q.pop();
                }
            } else {
                for (const unsigned & v : adj[u]) {
                    if (res[u][v] > 0 && !visited[v]) {
                        visited[v] = true;
                        q.push(v);
                        p[v] = u;
                    }
                }
            }
        }

        // find the min edge weight ‘f’ in this path, if any
        augment(res, p, s, f, t, std::numeric_limits<unsigned>::max());

        if (f == 0) {
            // we cannot send any more flow (‘f’ = 0), terminate
            break;
        }

        // we can still send a flow, increase the max flow!
        mf += f;
    }

    // this is the max flow value
    return mf;
}

int main() {
    std::string line;
    std::vector<std::vector<unsigned>> adj, res;
    unsigned source = 0, sink = 37, app, computer, n, total, i, mf;
    bool alloc;
    
    while (std::getline(std::cin, line)) {
        adj.assign(38, std::vector<unsigned>());
        res.assign(38, std::vector<unsigned>(38, 0));
        total = 0;

        for (computer = 27; computer < 37; computer++) {
            adj[computer].push_back(sink);
            adj[sink].push_back(computer);
            res[computer][sink] = 1;
        }

        do {
            app = line[0] - 'A' + 1;
            n = line[1] - '0';
            adj[source].push_back(app);
            adj[app].push_back(source);
            res[source][app] = n;
            total += n;
            
            for (i = 3; line[i] != ';'; i++) {
                computer = line[i] - '0' + 27;
                adj[app].push_back(computer);
                adj[computer].push_back(app);
                res[app][computer] = std::numeric_limits<unsigned>::max();
            }
        } while (std::getline(std::cin, line) && line != "");

        mf = edmondsKarp(adj, res, source, sink);

        if (mf != total) {
            std::cout << "!" << std::endl;
        } else {
            for (computer = 27; computer < 37; computer++) {
                alloc = false;

                for (app = 1; app < 27 && !alloc; app++) {
                    if (res[computer][app]) {
                        std::cout << (char)(app + 'A' - 1);
                        alloc = true;
                    }
                }

                if (!alloc) {
                    std::cout << "_";
                }
            }

            std::cout << std::endl;
        }
    }

	return 0;
}
