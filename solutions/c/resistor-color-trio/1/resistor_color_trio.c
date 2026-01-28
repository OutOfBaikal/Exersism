#include "resistor_color_trio.h"
#include <stdlib.h>
#include <stdio.h>
#include <math.h>
resistor_value_t color_code(resistor_band_t colors[]) {
    //int n = sizeof(colors) / sizeof(resistor_band_t);
    long res = 0;
    for (int i = 0; i < 2; ++i) {
        res *= 10;
        res += colors[i];
    }
    res *= pow(10.0, colors[2]);
    long int buf = res, a = 0;
    while (buf >= 1000) {
        buf /= 1000;
        a++;
    }
    resistor_value_t result;
    if (a == 0) {
        result.unit = OHMS;
    } else if (a == 1) {
        result.unit = KILOOHMS;
    } else if (a == 2) {
        result.unit = MEGAOHMS;
    } else {
        result.unit = GIGAOHMS;
    }
    printf("%ld %ld", res, buf);
    result.value = buf;
    return result;
}
