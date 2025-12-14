#include "resistor_color_duo.h"
#include <stdlib.h>

uint16_t color_code(resistor_band_t colors[]) {
    //int n = sizeof(colors) / sizeof(resistor_band_t);
    int res = 0;
    for (int i = 0; i < 2; ++i) {
        res *= 10;
        res += colors[i];
    }
    if (res > 99)
        res /= 10;

    return res;
}
