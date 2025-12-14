#include "luhn.h"

bool luhn(const char *num) {
    size_t n = strlen(num);
    char* cleaned = malloc(sizeof(char) * (n + 1));
    size_t j = 0;
    for (size_t i = 0; i < n; ++i) {
        if (num[i] >= '0' && num[i] <= '9') {
            cleaned[j] = num[i];
            ++j;
        } else if (num[i] != ' ') {
            free(cleaned);
            return false;
        }
    }
    cleaned[j] = '\0';
    size_t len = j;
    if (len <= 1) {
        free(cleaned);
        return false;
    }
    int sum = 0;
    bool shouldDouble = false;
    
    for (int i = len - 1; i >= 0; --i) {
        int digit = cleaned[i] - '0';
        int value = digit;
        if (shouldDouble) {
            int doubled = digit * 2;
            if (doubled > 9) {
                doubled -= 9;
            }
            value = doubled;
        }
        sum += value;
        shouldDouble = !shouldDouble;
    }
    free(cleaned);
    return sum % 10 == 0;
}