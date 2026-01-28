#include "sum_of_multiples.h"
#include <stdlib.h>


unsigned int sum(const unsigned int *factors, const size_t number_of_factors,
                 const unsigned int limit) {
     if (limit <= 0) return 0;

    char *used = calloc(limit, sizeof(char));
    if (!used) {
        return 0;
    }

    for (size_t d = 0; d < number_of_factors; d++) {
        int divisor = factors[d];
        if (divisor == 0) continue;

        for (unsigned int i = divisor; i < limit; i += divisor) {
            used[i] = 1;
        }
    }

    int res = 0;
    for (unsigned int i = 0; i < limit; i++) {
        if (used[i]) {
            res += i;
        }
    }

    free(used);
    return res;
                 }