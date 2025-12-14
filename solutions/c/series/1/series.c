#include "series.h"
#include <string.h>
#include <stdlib.h>
#include <stddef.h>

// typedef struct slices {
//    unsigned int substring_count;
//    char **substring;   // array of pointers of dimension substring_count
// } slices_t;

slices_t slices(char *input_text, unsigned int substring_length) {
    slices_t result;
    unsigned int len = strlen(input_text);
    if (substring_length > len || substring_length == 0) {
        result.substring_count = 0;
        result.substring = NULL;
        return result;
    }
    result.substring_count = len - substring_length + 1;
    result.substring = malloc(sizeof(char*) * result.substring_count);
    if (!result.substring) {
        result.substring_count = 0;
        return result;
    }
    
    for (size_t i = 0; i < result.substring_count; ++i) {
        result.substring[i] = malloc(sizeof(char) * (substring_length + 1));
        if (!result.substring[i]) {
            for (size_t k = 0; k < i; ++k) {
                free(result.substring[k]);
            }
            free(result.substring);
            result.substring_count = 0;
            return result;
        }
        for (size_t j = 0; j < substring_length; ++j) {
            result.substring[i][j] = input_text[i + j];
        }
        result.substring[i][substring_length] = '\0';
    }

    return result;
}