#include "pascals_triangle.h"
#include <stdlib.h>
#include <stdio.h>

void free_triangle(uint8_t **triangle, size_t rows) {
    if (triangle == NULL) return;
    for (size_t i = 0; i < rows; ++i) {
        free(triangle[i]);
    }
    free(triangle);
}
uint8_t **create_triangle(size_t rows) {
    if (rows == 0) {
        uint8_t **result = malloc(sizeof(uint8_t *) * 1);
        result[0] = calloc(1, sizeof(uint8_t));
        return result;
    }
    uint8_t **result = malloc(sizeof(uint8_t *) * rows);
    if (result == NULL) return NULL;
    for (size_t i = 0; i < rows; ++i) {
        size_t cols = i + 1;
        result[i] = calloc(rows, sizeof(uint8_t));
        if (result[i] == NULL) {
            free_triangle(result, i);
            return NULL;
        }
        
        result[i][0] = 1;
        if (cols > 1) {
            result[i][cols - 1] = 1;
        }
        for (size_t j = 1; j < cols - 1; ++j) {
            result[i][j] = result[i - 1][j - 1] + result[i - 1][j];
        }
        for (size_t j = 0; j < cols; ++j) {
            printf("%d ", result[i][j]);
        }
        printf("\n");
    }
    
    return result;
}