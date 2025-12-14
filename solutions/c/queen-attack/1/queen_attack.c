#include "queen_attack.h"
#include <stdlib.h>

attack_status_t can_attack(position_t queen_1, position_t queen_2) {
    // Проверяем, что позиции находятся в пределах доски и не совпадают
    if ((queen_1.row > 7) || (queen_1.column > 7) ||
        (queen_2.row > 7) || (queen_2.column > 7) ||
        (queen_1.row == queen_2.row && queen_1.column == queen_2.column)) {
        return INVALID_POSITION;
    }

    // Проверка на одну линию (горизонталь или вертикаль)
    if (queen_1.row == queen_2.row || queen_1.column == queen_2.column) {
        return CAN_ATTACK;
    }

    // Проверка на одну диагональ
    if (abs(queen_1.row - queen_2.row) == abs(queen_1.column - queen_2.column)) {
        return CAN_ATTACK;
    }

    // Королевы не могут атаковать друг друга
    return CAN_NOT_ATTACK;
}