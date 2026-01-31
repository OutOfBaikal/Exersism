#include "scrabble_score.h"
#include <stddef.h>
#include <string.h>
#include <stdlib.h>

char* to_lower(const char *word);

unsigned int score(const char *word) {
    unsigned int res = 0;
    size_t n = strlen(word);
    char *lower = to_lower(word);
    int data[26];
    data[0] = 1;
    data[1] = 3;
    data[2] = 3;
    data[3] = 2;
    data[4] = 1;
    data[5] = 4;
    data[6] = 2;
    data[7] = 4;
    data[8] = 1;
    data[9] = 8;
    data[10] = 5;
    data[11] = 3;
    data[12] = 1;
    data[13] = 1;
    data[14] = 1;
    data[15] = 3;
    data[16] = 10;
    data[17] = 1;
    data[18] = 1;
    data[19] = 1;
    data[20] = 1;
    data[21] = 4;
    data[22] = 4;
    data[23] = 8;
    data[24] = 4;
    data[25] = 10;

    for (size_t i = 0; i < n; ++i) {
        res += data[lower[i] - 'a'];
    }
    free(lower);
    return res;
}

char* to_lower(const char *word) {
    size_t n = strlen(word);
    char *result = (char*)malloc(sizeof(char) * n);
    
    for (size_t i = 0; i < n; ++i) {
        if (word[i] >= 'A' && word[i] <= 'Z') {
            result[i] = word[i] + ('a' - 'A');
        } else {
            result[i] = word[i];
        }
    }

    return result;
}