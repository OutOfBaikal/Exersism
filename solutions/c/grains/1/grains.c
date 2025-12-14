#include "grains.h"

uint64_t square(uint8_t index) {
    if (index < 1 || index > 64)
        return 0; // или какое-то другое значение, обозначающее ошибку

    uint64_t res = 1;
    for (uint8_t i = 2; i <= index; ++i) {
        res <<= 1;
    }

    return res;
}

uint64_t total(void) {
    uint64_t sum = 0;
    uint64_t res = 1;

    for (int i = 1; i <= 64; ++i) {
        sum += res;
        res <<= 1;
    }

    return sum;
}