#include "beer_song.h"
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>

void recite(uint8_t start_bottles, uint8_t take_down, char **song) {
    int line_index = 0;
    int current = start_bottles;

    for (int verse = 0; verse < take_down; ++verse) {
        // Первая строка куплета
        if (current > 2) {
            sprintf(song[line_index++], "%d bottles of beer on the wall, %d bottles of beer.", current, current);
            sprintf(song[line_index++], "Take one down and pass it around, %d bottles of beer on the wall.", current - 1);
        } else if (current == 2) {
            sprintf(song[line_index++], "2 bottles of beer on the wall, 2 bottles of beer.");
            sprintf(song[line_index++], "Take one down and pass it around, 1 bottle of beer on the wall.");
        } else if (current == 1) {
            sprintf(song[line_index++], "1 bottle of beer on the wall, 1 bottle of beer.");
            sprintf(song[line_index++], "Take it down and pass it around, no more bottles of beer on the wall.");
        } else { // current == 0, если вдруг
            sprintf(song[line_index++], "No more bottles of beer on the wall, no more bottles of beer.");
            sprintf(song[line_index++], "Go to the store and buy some more, 99 bottles of beer on the wall.");
        }

        current--;

        // Добавляем пустую строку между куплетами, кроме последнего
        if (verse != take_down - 1) {
            song[line_index][0] = '\0'; // пустая строка
            line_index++;
        }
    }

    //song_out = song;
}