#include <time.h>
#include <stddef.h>
#include <math.h>
#include <stdlib.h>  
#include "dnd_character.h"

int ability(void) {
    time(NULL);
    int data[4];
    int res = 0;
    for (int i = 0; i < 4; ++i) {
        data[i] = 1 + rand() % 6;
        res += data[i];
    }
    int min = 7;
    for (int i = 0; i < 4; ++i) {
        if (min > data[i])
            min = data[i];
    }
    return res - min;
}

int modifier(int score) {
    return floor((score - 10.0) / 2.0);
}

dnd_character_t make_dnd_character(void) {
    dnd_character_t res;
    res.strength = ability();
    res.dexterity = ability();
    res.constitution = ability();
    res.intelligence = ability();
    res.wisdom = ability();
    res.charisma = ability();
    res.hitpoints = 10 + modifier(res.constitution);
    return res;
}