#include "rotational_cipher.h"
#include <string.h>
#include <stdlib.h>

char *rotate(const char *text, int shift_key) {
    size_t n = strlen(text);
    char *result = (char*)malloc(sizeof(char) * n);
    for (size_t i = 0; i < n; ++i) {
        if (text[i] >= 'A' && text[i] <= 'Z') {
        	result[i] = (text[i] - 'A' + shift_key) % 26 + 'A';
        } else if (text[i] >= 'a' && text[i] <= 'z') {
        	result[i] = (text[i] - 'a' + shift_key) % 26 + 'a';
        } else {
            result[i] = text[i];
        }
    }

    return result;
}