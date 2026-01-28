#include "prime_factors.h"
#include <math.h>
#include <stdbool.h>
#include <stdlib.h>
#include <stdio.h>
uint64_t* sieve(uint64_t n, uint64_t *count);

size_t find_factors(uint64_t n, uint64_t factors[static MAXFACTORS]) {
    if (n < 2) {
        return 0;
    }
    uint64_t count;
    uint64_t *primes;
    if (n > 1000000000) {
        primes = sieve((uint64_t)sqrt(n), &count);
    } else {
        primes = sieve(n, &count);
    }
    if (!primes) {
        return 0; // No primes found or memory allocation failed
    }
    size_t t = 0;
    for (uint64_t i = 0; i < count; ++i) {
        while (n % primes[i] == 0) {
            if (t >= MAXFACTORS) {
                break; // Prevent overflow
            }
            factors[t++] = primes[i];
            n /= primes[i];
        }
        if (n == 1)
            break;
    }
    if (n > 1 && t < MAXFACTORS) {
        factors[t++] = n;
    }
    free(primes);
    return t;
}

uint64_t* sieve(uint64_t n, uint64_t *count){
    if (n < 2) {
        *count = 0;
        return NULL; // Return 0 for the 0th prime
    }
    // Estimate the upper limit for the nth prime
    uint64_t limit = n * (log(n) + log(log(n)));
    if (limit < 2) {
        limit = 2; // Ensure a minimum limit
    }
    // Allocate memory for the sieve
    bool *flags = malloc(sizeof(bool) * (limit + 1));
    if (!flags) {
        *count = 0;
        printf("%llu\n", (unsigned long long)count);
        return NULL; // Memory allocation failed
    }
    // Initialize the sieve
    for (uint64_t i = 2; i <= limit; ++i) {
        flags[i] = false; // false means prime
    }
    // Sieve of Eratosthenes
    for (uint64_t i = 2; i * i <= limit; ++i) {
        if (!flags[i]) { // If i is prime
            for (uint64_t j = i * i; j <= limit; j += i) {
                flags[j] = true; // Mark multiples as non-prime
            }
        }
    }
    uint64_t *res = malloc(sizeof(uint64_t) * limit); // Allocate sufficient space
    if (!res) {
        free(flags);
        *count = 0;
        printf("%llu\n", (unsigned long long)count);
        return NULL; // Memory allocation failed
    }
    *count = 0;
    for (uint64_t i = 2; i <= limit; ++i) {
        if (!flags[i]) { // If i is prime
            res[(*count)++] = i; // Store the prime and increment count
        }
    }
    printf("%llu\n", (unsigned long long)count);
    free(flags);
    return res;
}