#include "anagram.h"
#include <string.h>
#include <stdbool.h>

bool is_equal(char *str1, char *str2);

bool is_anagram(char *str1, char *str2);

void find_anagrams(const char *subject, struct candidates *candidates) {
    for (size_t i = 0; i < candidates->count; ++i) {
        if (is_equal((void *)subject, (void *)candidates->candidate[i].word)) {
            candidates->candidate[i].is_anagram = NOT_ANAGRAM;
            continue;
        }
        if (is_anagram((void *)subject, (void *)candidates->candidate[i].word)) {
            candidates->candidate[i].is_anagram = IS_ANAGRAM;
        } else {
            candidates->candidate[i].is_anagram = NOT_ANAGRAM;
        }
    }
}

bool is_equal(char *str1, char *str2) {
    size_t n1 = strlen(str1), n2 = strlen(str2);
    if (n1 != n2) {
        return false;
    }
    for (size_t i = 0; i < n1; ++i) {
        short unsigned int a1, a2;
        if (str1[i] >= 'A' && str1[i] <= 'Z') {
            a1 = str1[i] - 'A';
        } else if (str1[i] >= 'a' && str1[i] <= 'z') {
            a1 = str1[i] - 'a';
        }
        if (str2[i] >= 'A' && str2[i] <= 'Z') {
            a2 = str2[i] - 'A';
        } else if (str2[i] >= 'a' && str2[i] <= 'z') {
            a2 = str2[i] - 'a';
        }
        if (a1 != a2) {
            return false;
        }
    }
    return true;
}

bool is_anagram(char *str1, char *str2) {
    size_t n1 = strlen(str1), n2 = strlen(str2);
    if (n1 != n2) {
        return false;
    }
    char data1[26] = {0}, data2[26] = {0};
    for (size_t i = 0; i < n1; ++i) {
        short unsigned int a1, a2;
        if (str1[i] >= 'A' && str1[i] <= 'Z') {
            a1 = str1[i] - 'A';
        } else if (str1[i] >= 'a' && str1[i] <= 'z') {
            a1 = str1[i] - 'a';
        }
        if (str2[i] >= 'A' && str2[i] <= 'Z') {
            a2 = str2[i] - 'A';
        } else if (str2[i] >= 'a' && str2[i] <= 'z') {
            a2 = str2[i] - 'a';
        }
        ++data1[a1];
        ++data2[a2];
    }
    for (size_t i = 0; i < 26; ++i) {
        if (data1[i] != data2[i]) {
            return false;
        }
    }
    return true;
}