#include "nth_prime.h"
#include "math.h"
#include "stdlib.h"
#include <stdbool.h>
#include <stdio.h>
uint32_t nth(uint32_t n) {
    if (n == 0) {
        return 0; // Return 0 for the 0th prime
    }

    // Estimate the upper limit for the nth prime
    uint32_t limit = n * (log(n) + log(log(n)));
    if (limit <= 2) {
        limit = 5; // Ensure a minimum limit
    }

    // Allocate memory for the sieve
    bool *flags = (bool*)malloc(sizeof(bool) * (limit + 1));
    if (!flags) {
        return 0; // Memory allocation failed
    }

    // Initialize the sieve
    for (uint32_t i = 2; i <= limit; ++i) {
        flags[i] = false; // false means prime
    }

    // Sieve of Eratosthenes
    for (uint32_t i = 2; i * i <= limit; ++i) {
        if (!flags[i]) { // If i is prime
            for (uint32_t j = i * i; j <= limit; j += i) {
                flags[j] = true; // Mark multiples as non-prime
            }
        }
    }

    // Collect primes
    uint32_t count = 0;
    uint32_t res = 0;
    for (uint32_t i = 2; i <= limit; ++i) {
        if (!flags[i]) { // If i is prime
            count++;
            if (count == n) {
                res = i; // Found the nth prime
                break;
            }
        }
    }

    free(flags);
    return res;
}