#include "sublist.h"
#include <stdbool.h>

//typedef enum { EQUAL, UNEQUAL, SUBLIST, SUPERLIST } comparison_result_t;

bool is_match(int *list_to_compare, int *base_list, size_t compare_count, size_t begin_base);
bool is_equal(int *list_to_compare, int *base_list, size_t list_to_compare_element_count, size_t base_list_element_count, size_t begin_base);
bool is_superlist(int *list_to_compare, int *base_list, size_t list_to_compare_element_count, size_t base_list_element_count);

bool is_match(int *list_to_compare, int *base_list, size_t compare_count, size_t begin_base) {
    for (size_t i = 0; i < compare_count; ++i) {
        if (list_to_compare[i] != base_list[begin_base + i]) {
            return false;
        }
    }
    return true;
}

bool is_equal(int *list_to_compare, int *base_list, size_t list_to_compare_element_count, size_t base_list_element_count, size_t begin_base) {
    if (list_to_compare_element_count != base_list_element_count)
        return false;
    size_t j = begin_base;
    for (size_t i = 0; i < base_list_element_count; ++i) {
        if (list_to_compare[i] != base_list[j]) {
            return false;
        }
        ++j;
    }

    return true;
}

bool is_superlist(int *list_to_compare, int *base_list, size_t list_to_compare_element_count, size_t base_list_element_count) {
    if (base_list_element_count > list_to_compare_element_count) return false;
    size_t n = list_to_compare_element_count - base_list_element_count + 1;
    for (size_t i = 0; i < n; ++i) {
        if (is_match(base_list, list_to_compare, base_list_element_count, i)) {
            return true;
        }
    }
    return false;
}

comparison_result_t check_lists(int *list_to_compare, int *base_list,
                                size_t list_to_compare_element_count,
                                size_t base_list_element_count) {
    comparison_result_t result;
    if (is_equal(list_to_compare, base_list, list_to_compare_element_count, base_list_element_count, 0)) {
        result = EQUAL;
    } else if (is_superlist(list_to_compare, base_list, list_to_compare_element_count, base_list_element_count)) {
        result = SUPERLIST;
    } else if (is_superlist(base_list, list_to_compare, base_list_element_count, list_to_compare_element_count)) {
        result = SUBLIST; // Здесь было SUPERLIST, но должно быть SUBLIST
    } else {
        result = UNEQUAL;
    }

    return result;
}