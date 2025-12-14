#include "binary_search.h"
#include <stdlib.h>

int compare_ints(const void *a, const void *b); 

int compare_ints(const void *a, const void *b) {
    int int_a = *(const int *) a;
    int int_b = *(const int *) b;
    if (int_a < int_b) return -1;
    else if (int_a > int_b) return 1;
    else return 0;
}

const int *binary_search(int value, const int *arr, size_t length) {
    //qsort((void *)arr, length, sizeof(int *), compare_ints);
    int left = 0, right = length - 1;
    while (left <= right) {
        const int mid = (left + right) >> 1;
        if (arr[mid] == value) {
            return &arr[mid];
        }
        if (arr[mid] < value) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return NULL;
}