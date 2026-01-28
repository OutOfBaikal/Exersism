#include "perfect_numbers.h"
kind classify_number(int n) {
    if (n < 1) {
        return ERROR;
    }
    int sum = 0;
    for (int i = 1; i < n / 2 + 1; ++i) {
        if (n % i == 0) {
            sum += i;
        }
    }
    if (sum == n) {
        return PERFECT_NUMBER;
    } else if (sum > n) {
        return ABUNDANT_NUMBER;
    }
    return DEFICIENT_NUMBER;
}