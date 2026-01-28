#include "pangram.h"
#include <stdbool.h>
#include <string.h>
bool is_pangram(const char *sentence) {
    if (sentence == NULL) {
        return false;
    }
    char data[26] = {0};
    int n = strlen(sentence);
    if (n < 26) {
        return false;
    }
    for (int i = 0; i < n; ++i) {
        if (sentence[i] >= 'A' && sentence[i] <= 'Z') {
            ++data[sentence[i] - 'A'];
        }
        else if (sentence[i] >= 'a' && sentence[i] <= 'z') {
            ++data[sentence[i] - 'a'];
        }
    }
    for (int i = 0; i < 26; ++i) {
        if (data[i] == 0)
            return false;
    }
    return true;
}