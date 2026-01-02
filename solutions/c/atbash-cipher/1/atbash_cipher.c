#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include "atbash_cipher.h"

static char transform_char(char c) {
    if (c >= 'a' && c <= 'z') {
        return 'z' - (c - 'a'); // Transform lowercase letters
    } else if (c >= 'A' && c <= 'Z') {
        return 'z' - (c - 'A'); // Transform uppercase letters
    }
    return c; // Return non-alphabetic characters unchanged
}

char *atbash_encode(const char *input) {
    if (!input) return NULL;

    size_t len = strlen(input);
    // Allocate buffer for transformed characters (worst case: all alphanumeric)
    char *transformed = calloc(len + 1, sizeof(char));
    if (!transformed) return NULL;

    size_t t_idx = 0;
    for (size_t i = 0; i < len; i++) {
        char c = input[i];
        if (isalpha(c)) {
            transformed[t_idx++] = transform_char(c);
        } else if (isdigit(c)) {
            transformed[t_idx++] = c;  // Keep digits
        }
        // Skip spaces, punctuation, etc.
    }

    // Now format into groups of 5
    size_t total_chars = t_idx;
    size_t group_count = (total_chars + 4) / 5;  // Ceiling division
    size_t output_len = total_chars + group_count - 1;  // Spaces between groups
    char *result = calloc(output_len + 1, sizeof(char));  // +1 for null
    if (!result) {
        free(transformed);
        return NULL;
    }

    size_t r_idx = 0;
    for (size_t i = 0; i < total_chars; i++) {
        result[r_idx++] = transformed[i];
        // Add space after every 5 chars, but not at the end
        if ((i + 1) % 5 == 0 && (i + 1) < total_chars) {
            result[r_idx++] = ' ';
        }
    }
    result[r_idx] = '\0';

    free(transformed);
    return result;
}

char *atbash_decode(const char *input) {
    if (!input) return NULL;

    size_t len = strlen(input);
    // Allocate buffer for transformed characters (worst case: all alphanumeric)
    char *transformed = calloc(len + 1, sizeof(char));
    if (!transformed) return NULL;

    size_t t_idx = 0;
    for (size_t i = 0; i < len; i++) {
        char c = input[i];
        if (isalpha(c)) {
            transformed[t_idx++] = transform_char(c);
        } else if (isdigit(c)) {
            transformed[t_idx++] = c;  // Keep digits
        }
        // Skip spaces, punctuation, etc.
    }

    // Now format into groups of 5
    size_t total_chars = t_idx;
    size_t group_count = (total_chars + 4) / 5;  // Ceiling division
    size_t output_len = total_chars + group_count - 1;  // Spaces between groups
    char *result = calloc(output_len + 1, sizeof(char));  // +1 for null
    if (!result) {
        free(transformed);
        return NULL;
    }

    size_t r_idx = 0;
    for (size_t i = 0; i < total_chars; i++) {
        result[r_idx++] = transformed[i];
    }
    result[r_idx] = '\0';

    free(transformed);
    return result;
}