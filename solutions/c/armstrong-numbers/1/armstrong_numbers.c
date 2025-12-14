#include "armstrong_numbers.h"
#include <math.h>

bool is_armstrong_number(int candidate) {
    int nums[12];
    int buf = candidate;
    int sum = 0;
    int j = 0;
    while (buf > 0) {
        nums[j] = buf % 10;
        buf /= 10;
        ++j;
    }
    for (int i = 0; i < j; ++i)
        sum += pow(nums[i], j);
    return sum == candidate;
}