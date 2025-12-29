#include "protein_translation.h"
#include <string.h>

// typedef enum {
//    Methionine,
//    Phenylalanine,
//    Leucine,
//    Serine,
//    Tyrosine,
//    Cysteine,
//    Tryptophan,
// } amino_acid_t;

// typedef struct {
//    bool valid;
//    size_t count;
//    amino_acid_t amino_acids[MAX_AMINO_ACIDS];
// } protein_t;

protein_t protein(const char *const rna) {
    protein_t result;
    size_t len = strlen(rna);
    
    // Проверка длины строки
    if (len == 0) {
        result.valid = true;
        result.count = 0;
        return result;
    }
    
    for (size_t i = 0; i < len; i += 3) {
        // Выделяем текущий кодон (3 символа)
        char codon[4] = { rna[i], rna[i + 1], rna[i + 2], '\0' };
        if (strcmp(codon, "UAA") == 0 || strcmp(codon, "UAG") == 0 || strcmp(codon, "UGA") == 0) {
            len = i;
        }
    }
     if (len % 3 != 0) {
         result.valid = false;
         result.count = 0;
         return result;
    }

    result.valid = true;  // Изначально считаем строку валидной
    result.count = 0;

    for (size_t i = 0; i < len; i += 3) {
        // Выделяем текущий кодон (3 символа)
        char codon[4] = { rna[i], rna[i + 1], rna[i + 2], '\0' };

        // Сравниваем кодон через if-else (switch со строками невозможен)
        if (strcmp(codon, "AUG") == 0) {
            result.amino_acids[result.count] = Methionine;
        } else if (strcmp(codon, "UUU") == 0 || strcmp(codon, "UUC") == 0) {
            result.amino_acids[result.count] = Phenylalanine;
        } else if (strcmp(codon, "UUA") == 0 || strcmp(codon, "UUG") == 0) {
            result.amino_acids[result.count] = Leucine;
        } else if (strcmp(codon, "UCU") == 0 || strcmp(codon, "UCC") == 0 ||
                   strcmp(codon, "UCA") == 0 || strcmp(codon, "UCG") == 0) {
            result.amino_acids[result.count] = Serine;
        } else if (strcmp(codon, "UAU") == 0 || strcmp(codon, "UAC") == 0) {
            result.amino_acids[result.count] = Tyrosine;
        } else if (strcmp(codon, "UGU") == 0 || strcmp(codon, "UGC") == 0) {
            result.amino_acids[result.count] = Cysteine;
        } else if (strcmp(codon, "UGG") == 0) {
            result.amino_acids[result.count] = Tryptophan;
        } else if (strcmp(codon, "UAA") == 0 || strcmp(codon, "UAG") == 0 ||
                   strcmp(codon, "UGA") == 0) {
            // Стоп-кодон: завершаем трансляцию
            break;
        } else {
            // Неизвестный кодон: метка как невалидный
            result.valid = false;
            result.count = 0;
            break;
        }

        result.count++;
    }

    return result;
}