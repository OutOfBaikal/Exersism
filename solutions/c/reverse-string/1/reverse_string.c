#include "reverse_string.h"
#include <string.h>
#include <stdlib.h>

char *reverse(const char *value) {
    int n = strlen(value);
    char *result = (char*)malloc(sizeof(char) * n);
    for (int i = n - 1; i >= 0; --i) {
        result[i] = value[n - i - 1];
    }
    return result;
}