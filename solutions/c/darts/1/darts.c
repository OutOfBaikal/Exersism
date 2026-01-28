#include "darts.h"
#include <math.h>

uint8_t score(coordinate_t landing_position) {
    float distance = hypotf(landing_position.x, landing_position.y);
    if (distance > 10.0f)
        return 0;
    else if (distance > 5.0f)
        return 1;
    else if (distance > 1.0f)
        return 5;
    else
        return 10;
}