#include "collatz_conjecture.h"

int steps(int start) {
    if (start <= 0) {
        // Вернуть ошибку или исключение
        return -1; // или какое-то другое значение, обозначающее ошибку
    }
    
    int res = 0;
    while (start != 1) {
        if (start % 2 == 1) {
            start = 3 * start + 1; 
        } else {
            start /= 2;
        }
        ++res;
    }

    return res;
}