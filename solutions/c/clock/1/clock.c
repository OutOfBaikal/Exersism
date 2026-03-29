#include "clock.h"
#include <string.h>
#include <stdio.h>
#include <stdbool.h>

// typedef struct {
//    char text[MAX_STR_LEN];
// } clock_t;

clock_t clock_create(int hour, int minute) {
    clock_t result;
    while (minute < 0) {
        minute += 60;
        hour--;
    }
    int h = (hour + minute / 60) % 24;
    int m = minute % 60;
    while (h < 0) {
        h += 24;
    }
    
    snprintf(result.text, MAX_STR_LEN, "%02d:%02d", h, m);
    return result;
}
clock_t clock_add(clock_t clock, int minute_add) {
    int h, m;
    sscanf(clock.text, "%d:%d", &h, &m);
    int total_m = m + minute_add;
    h = (h + total_m / 60) % 24;
    while (h < 0) {
        h += 24;
    }
    m = total_m % 60;
    while (m < 0) {
        m += 60;
        h -= 1;
    }
    while (h < 0) {
        h += 24;
    }
    clock_t result;
    snprintf(result.text, MAX_STR_LEN, "%02d:%02d", h, m);
    return result;
}
clock_t clock_subtract(clock_t clock, int minute_subtract) {
    int h, m;
    sscanf(clock.text, "%d:%d", &h, &m);
    int total_m = m - minute_subtract;
    h = (h + total_m / 60) % 24;
    
    while (h < 0) {
        h += 24;
    }
    m = total_m % 60;
    
    if (m < 0) {
        m += 60;
        h -= 1;
    }
    
    while (h < 0) {
        h += 24;
    }

    clock_t result;
    snprintf(result.text, MAX_STR_LEN, "%02d:%02d", h, m);
    return result;
}
bool clock_is_equal(clock_t a, clock_t b) {
    return strcmp(a.text, b.text) == 0;
}