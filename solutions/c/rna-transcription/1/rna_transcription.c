#include "rna_transcription.h"
#include <string.h>
#include <stdlib.h>

char *to_rna(const char *dna) {
    size_t n = strlen(dna);
    char *res = (char*)malloc(sizeof(char) * n);
    for (size_t i = 0; i < n; ++i) {
        if (dna[i] == 'G')
            res[i] = 'C';
        else if (dna[i] == 'C')
            res[i] = 'G';
        else if (dna[i] == 'T')
            res[i] = 'A';
        else if (dna[i] == 'A')
            res[i] = 'U';
    }
    return res;
}