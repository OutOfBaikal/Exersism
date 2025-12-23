#include "spiral_matrix.h"
#include <stdlib.h>

// typedef struct {
//    int size;
//    int **matrix;
// } spiral_matrix_t;

spiral_matrix_t* spiral_matrix_create(int size) {
    spiral_matrix_t *result = malloc(sizeof(spiral_matrix_t));
    if (size == 0) {
        return result;
    }
    result->size = size;
    result->matrix = malloc(sizeof(int*) * size);
    for (int i = 0; i < size; ++i) {
        result->matrix[i] = malloc(sizeof(int) * size);
    }
    int i_min = 0, i_max = size, j_min = 0, j_max = size;
    int num = 1;
    while (i_min < i_max && j_min < j_max) {
        for (int j = j_min; j < j_max; ++j) {
            result->matrix[i_min][j] = num;
            ++num;
        }
        ++i_min;
        for (int i = i_min; i < i_max; ++i) {
			result->matrix[i][j_max-1] = num;
			++num;
		}
		--j_max;

        if (i_min < i_max) {
			for (int j = j_max - 1; j >= j_min; --j) {
				result->matrix[i_max-1][j] = num;
				++num;
			}
			--i_max;
		}
		if (j_min < j_max) {
			for (int i = i_max - 1; i >= i_min; --i) {
				result->matrix[i][j_min] = num;
				++num;
			}
			++j_min;
		}
    }

    return result;
}

void spiral_matrix_destroy(spiral_matrix_t* matrix) {
    for (int i = 0; i < matrix->size; ++i) {
        free(matrix->matrix[i]);
    }
    free(matrix->matrix);
    free(matrix);
    return;
}