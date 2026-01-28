#include "square_root.h"
#include <stdint.h>
unsigned int square_root(unsigned int num) {
    float number = (float)num;
    const float x2 = number * 0.5F;
	const float threehalfs = 1.5F;

	union {
		float f;
		uint32_t i;
	} conv = {number};
	conv.i = 0x5f3759df - ( conv.i >> 1 );
	conv.f *= threehalfs - x2 * conv.f * conv.f;
	return (unsigned int)(1 / conv.f);
}