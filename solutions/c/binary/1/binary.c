#include "binary.h"
#include <string.h>
int convert(const char *input) {
    int n = strlen(input);
    int res = 0;
    for (int i = n - 1; i >= 0; --i) {
        if (input[i] != '0' && input[i] != '1') {
            return -1;
        }
        res += (input[i] & 1) << (n - 1 - i);
    }
    return res;
}