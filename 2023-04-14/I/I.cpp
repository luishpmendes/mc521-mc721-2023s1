// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=243&page=show_problem&problem=3342

// 12190 - Electric Bill

// It’s year 2100. Electricity has become very expensive. Recently, your electricity company raised the
// power rates once more. The table below shows the new rates (consumption is always a positive integer):
// Range Price
// (Crazy-Watt-hour) (Americus)
// 1 ∼ 100 2
// 101 ∼ 10000 3
// 10001 ∼ 1000000 5
// > 1000000 7
// This means that, when calculating the amount to pay, the first 100 CWh have a price of 2 Americus
// each; the next 9900 CWh (between 101 and 10000) have a price of 3 Americus each and so on.
// For instance, if you consume 10123 CWh you will have to pay 2 × 100 + 3 × 9900 + 5 × 123 = 30515
// Americus.
// The evil mathematicians from the company have found a way to gain even more money. Instead of
// telling you how much energy you have consumed and how much you have to pay, they will show you
// two numbers related to yourself and to a random neighbor:
// A: the total amount to pay if your consumptions were billed together; and
// B: the absolute value of the difference between the amounts of your bills.
// If you can’t figure out how much you have to pay, you must pay another 100 Americus for such
// a “service”. You are very economical, and therefore you are sure you cannot possibly consume more
// than any of your neighbors. So, being smart, you know you can compute how much you have to pay.
// For example, suppose the company informed you the following two numbers: A = 1100 and B = 300.
// Then you and your neighbor’s consumptions had to be 150 CWh and 250 CWh respectively. The total
// consumption is 400 CWh and then A is 2×100+ 3×300 = 1100. You have to pay 2×100+ 3×50 = 350
// Americus, while your neighbor must pay 2 × 100 + 3 × 150 = 650 Americus, so B is |350 − 650| = 300.
// Not willing to pay the additional fee, you decided to write a computer program to find out how
// much you have to pay.

// Input

// The input contains several test cases. Each test case is composed of a single line, containing two integers
// A and B, separated by a single space, representing the numbers shown to you (1 ≤ A, B ≤ 10^9). You
// may assume there is always a unique solution, that is, there exists exactly one pair of consumptions
// that produces those numbers.
// The last test case is followed by a line containing two zeros separated by a single space.

// Output

// For each test case in the input, your program must print a single line containing one integer, representing
// the amount you have to pay.

// Sample Input

// 1100 300
// 35515 27615
// 0 0

// Sample Output

// 350
// 2900

// binary search the answer + algebra

#include <algorithm>
#include <iostream>
#include <string>
#include <vector>

unsigned consumption (unsigned value) {
    unsigned result = 0;

    if (value > 0) {
        if (value < 2 * 100) {
            result += value / 2;
        } else {
            result += 100;
        }
    }

    if (value > 2 * 100) {
        if (value < 3 * 9900 + 2 * 100) {
            result += (value - 2 * 100) / 3;
        } else {
            result += 9900;
        }
    }

    if (value > 3 * 9900 + 2 * 100) {
        if (value < 5 * 990000 + 3 * 9900 + 2 * 100) {
            result += (value - 3 * 9900 - 2 * 100) / 5;
        } else {
            result += 990000;
        }
    }

    if (value > 5 * 990000 + 3 * 9900 + 2 * 100) {
        result += (value - 5 * 990000 - 3 * 9900 - 2 * 100) / 7;
    }

    return result;
}

unsigned value (unsigned consumption) {
    unsigned result = 0;

    if (consumption > 0) {
        if (consumption < 100) {
            result += consumption * 2;
        } else {
            result += 2 * 100;
        }
    }

    if (consumption > 100) {
        if (consumption < 10000) {
            result += 3 * (consumption - 100);
        } else {
            result += 3 * 9900;
        }
    }

    if (consumption > 10000) {
        if (consumption < 1000000) {
            result += 5 * (consumption - 10000);
        } else {
            result += 5 * 990000;
        }
    }

    if (consumption > 1000000) {
        result += 7 * (consumption - 1000000);
    }
    
    return result;
}

int main () {
    unsigned a, b, total, lo, hi, mid, ans, diff;
    bool found;

    while (std::cin >> a >> b && (a || b)) {
        found = false;
        total = consumption(a);
        ans = 0;
        lo = 0;
        hi = total;

        while (lo < hi && !found) {
            mid = (lo + hi) / 2;
            diff = value(total - mid) - value(mid);

            if (diff == b) {
                ans = mid;
                found = true;
            } else if (diff > b) {
                lo = mid + 1;
            } else {
                hi = mid - 1;
            }
        }

        if (!found) {
            ans = lo;
        }

        std::cout << value(ans) << std::endl;
    }

    return 0;
}
