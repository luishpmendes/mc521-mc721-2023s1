// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=602&page=show_problem&problem=4286

// UVA 12608 - Garbage Collection

// Garbage truck is collecting garbage on its route starting at the dump and collecting garbage from pick up
// points in order (closest first). Garbage truck returns
// to the dump and empties its contents if it either
// 1. is full or
// 2. cannnot load the garbage at the current point
// because it will put the load over the weight limit
// (the truck driver has no way of knowing the
// weight of the garbage at a particular pick up
// point until the truck reaches that point) or
// 3. there are no more pick up points left
// The pick up run ends when the truck returns to the dump after the case #3. In the case where
// both rule #1 and rule #3 apply, we consider only the latter one (the run is over). Also
// note that there are no partial pick ups (e.g. this is not allowed: get a portion of garbage
// at some point until truck is full and come back for the rest).
// Given the weight limit that truck can carry and the list of pick up points with garbage weights,
// what is the distance that the truck will cover following the rules above?

// Input

// Input starts with an integer T on the first line, the number of test cases. Each case starts with the line
// with two integers W and N, the weight limit and the number of pick up points, respectively.
// Then N lines follow, each containing integers xi and wi, the distance from the dump and the weight
// of the garbage at the pick up point respectively.
// Constraints are as given: 1 ≤ W ≤ 1000, 1 ≤ N ≤ 1000, 0 < xi ≤ 100, 000, 1 ≤ wi ≤ W. Distances
// are distinct and given in order, i.e. for each i < j, xi < xj .

// Output

// For each test case output an integer on a line — the distance covered by the garbage truck.

// Sample Input

// 3
// 2 2
// 1 1
// 2 2
// 6 3
// 1 1
// 2 2
// 3 3
// 3 3
// 1 2
// 2 2
// 3 1

// Sample Output

// 8
// 6
// 10

// simulation with several corner cases

#include <iostream>
#include <vector>

int main () {
    unsigned T, W, N, x, w, i, load, dist, last_pos;

    std::cin >> T;

    while (T--) {
        std::cin >> W >> N;

        load = dist = last_pos = 0;

        for (i = 0; i < N; i++) {
            std::cin >> x >> w;

            dist += x - last_pos;

            if (load + w <= W) {
                load += w;
            } else { // load + w > W
                dist += 2 * x;
                load = w;
            }

            if (load == W && i + 1 < N) {
                dist += 2 * x;
                load = 0;
            }

            last_pos = x;
        }

        dist += last_pos;

        std::cout << dist << std::endl;
    }

    return 0;
}
