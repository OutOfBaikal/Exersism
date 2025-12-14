#include "eliuds_eggs.h"
unsigned int egg_count(int input) {
    unsigned int res = 0;
    for (int i = 0; i < 32; ++i) {
        if ((input >> i) & 1) {
            ++res;
        }
    } 
    return res;
}