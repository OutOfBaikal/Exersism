#include "hamming.h"
#include <stdlib.h>
#include <string.h>
int compute(const char *lhs, const char *rhs) {
    int a = strlen(lhs);
    if (strlen(lhs) != strlen(rhs)) {
        return -1;
    }
    int res = 0;
    for (int i = 0; i < a; ++i) {
        if (lhs[i] != rhs[i])
            ++res;
    }
    return res;
}