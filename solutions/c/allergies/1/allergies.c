#include "allergies.h"

bool is_allergic_to(allergen_t al, unsigned int num) {
    return ((num >> al) & 1) == 1;
}
allergen_list_t get_allergens(unsigned int num) {
    allergen_list_t result;
    result.count = 0;
    for (int i = 0; i < ALLERGEN_COUNT; ++i) {
        bool allergic = ((num >> i) & 1) == 1;
        result.allergens[i] = allergic;
        if (allergic) {
            result.count++;
        }
    }
    return result;
}