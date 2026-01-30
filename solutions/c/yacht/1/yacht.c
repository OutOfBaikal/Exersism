#include "yacht.h"
#include <string.h>

// typedef enum {
//    ONES,
//    TWOS,
//    THREES,
//    FOURS,
//    FIVES,
//    SIXES,
//    FULL_HOUSE,
//    FOUR_OF_A_KIND,
//    LITTLE_STRAIGHT,
//    BIG_STRAIGHT,
//    CHOICE,
//    YACHT
// } category_t;

// Вспомогательная функция: подсчитать количество каждого номинала (1..6)
static void count_values(dice_t *dice, int counts[7]) {
    memset(counts, 0, 7 * sizeof(int));
    for (int i = 0; i < 5; ++i) {
        if (dice->faces[i] >= 1 && dice->faces[i] <= 6) {
            counts[dice->faces[i]]++;
        }
    }
}

// Вспомогательная функция: сумма элементов массива
static int sum_array(const dice_t *dice) {
    int sum = 0;
    for (int i = 0; i < 5; ++i) {
        sum += dice->faces[i];
    }
    return sum;
}

int score(dice_t dice, category_t category) {
    int res = 0;
    int counts[7];  // counts[1]..counts[6] — количество кубиков с номиналом 1..6
    int valid;
    
    switch (category) {
        case ONES:
            for (int i = 0; i < 5; ++i) {
                if (dice.faces[i] == 1) res += 1;
            }
            break;

        case TWOS:
            for (int i = 0; i < 5; ++i) {
                if (dice.faces[i] == 2) res += 2;
            }
            break;

        case THREES:
            for (int i = 0; i < 5; ++i) {
                if (dice.faces[i] == 3) res += 3;
            }
            break;

        case FOURS:
            for (int i = 0; i < 5; ++i) {
                if (dice.faces[i] == 4) res += 4;
            }
            break;

        case FIVES:
            for (int i = 0; i < 5; ++i) {
                if (dice.faces[i] == 5) res += 5;
            }
            break;

        case SIXES:
            for (int i = 0; i < 5; ++i) {
                if (dice.faces[i] == 6) res += 6;
            }
            break;

        case FULL_HOUSE:
            count_values(&dice, counts);
            // Ищем тройку и пару
            int has_three = 0, has_two = 0;
            for (int val = 1; val <= 6; ++val) {
                if (counts[val] == 3) has_three = 1;
                if (counts[val] == 2) has_two = 1;
            }
            if (has_three && has_two) {
                res = sum_array(&dice);
            }
            break;

        case FOUR_OF_A_KIND:
            count_values(&dice, counts);
            for (int val = 1; val <= 6; ++val) {
                if (counts[val] >= 4) {
                    res = val * 4;
                    break;
                }
            }
            break;

        case LITTLE_STRAIGHT:
            count_values(&dice, counts);
            // Проверяем, что есть ровно по одному кубику 1,2,3,4,5
            valid = 1;
            for (int val = 1; val <= 5; ++val) {
                if (counts[val] != 1) {
                    valid = 0;
                    break;
                }
            }
            if (valid) res = 30;
            break;

        case BIG_STRAIGHT:
            count_values(&dice, counts);
            // Проверяем, что есть ровно по одному кубику 2,3,4,5,6
            valid = 1;
            for (int val = 2; val <= 6; ++val) {
                if (counts[val] != 1) {
                    valid = 0;
                    break;
                }
            }
            if (valid) res = 30;
            break;

        case CHOICE:
            res = sum_array(&dice);
            break;

        case YACHT:
            count_values(&dice, counts);
            for (int val = 1; val <= 6; ++val) {
                if (counts[val] == 5) {
                    res = 50;
                    break;
                }
            }
            break;

        default:
            res = 0;  // На случай неизвестного category
            break;
    }
    return res;
}