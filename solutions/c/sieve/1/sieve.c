#include "sieve.h"
#include <math.h>
#include <stdbool.h>
#include <stdlib.h>

uint32_t sieve(uint32_t limit, uint32_t *primes, size_t max_primes) {
    if (limit < 2 || primes == NULL || max_primes == 0) {
        return 0;  // Нет простых чисел или некорректный ввод
    }

    // Выделяем флаг-массив: индекс = число, значение = составное (true) / простое (false)
    bool *flags = malloc(sizeof(bool) * (limit + 1));
    if (!flags) {
        return 0;  // Ошибка выделения памяти
    }

    // Инициализируем: все числа считаются простыми
    for (uint32_t i = 2; i <= limit; ++i) {
        flags[i] = false;
    }

    // Решето Эратосфена: отмечаем составные числа
    for (uint32_t i = 2; i * i <= limit; ++i) {
        if (!flags[i]) {  // i — простое
            // Начинаем с i*i, т.к. меньшие кратные уже отмечены
            for (uint32_t j = i * i; j <= limit; j += i) {
                flags[j] = true;
            }
        }
    }

    // Собираем простые числа в выходной массив, следя за границей
    uint32_t count = 0;
    for (uint32_t i = 2; i <= limit; ++i) {
        if (!flags[i]) {  // i — простое
            if (count >= max_primes) {
                // Буфер переполнен — прекращаем
                free(flags);
                return count;  // Возвращаем то, что успели собрать
            }
            primes[count++] = i;
        }
    }

    free(flags);
    return count;  // Количество найденных простых чисел
}