#include "nucleotide_count.h"
#include <stddef.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>

char *count(const char *dna_strand) {
    int h[256] = {0}; 

    size_t n = strlen(dna_strand);
    for (size_t i = 0; i < n; ++i) {
        unsigned char c = (unsigned char)dna_strand[i];
        if (c != 'A' && c != 'C' && c != 'G' && c != 'T') {
            char *result = malloc(1); 
            if (result) {
                result[0] = '\0'; 
            }
            return result;
        }
        h[c]++;
    }
    char *result = malloc(128);
    if (!result) return NULL;
    snprintf(result, 128, "A:%d C:%d G:%d T:%d", h['A'], h['C'], h['G'], h['T']);
    return result;
}