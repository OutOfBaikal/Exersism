#include "gigasecond.h"
void gigasecond(time_t input, char *output, size_t size) {
    input += 1000000000;
    struct tm *tm_info = localtime(&input);
    strftime(output, size, "%Y-%m-%d %H:%M:%S", tm_info);
}