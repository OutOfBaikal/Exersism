#include "bob.h"
#include <stdbool.h>
#include <stddef.h>
#include <string.h>
#include <stdlib.h>

bool containsLetters(char *input);
bool isUpper(char *input);

char *hey_bob(char *greeting) {
    size_t begin = 0;
    size_t len = strlen(greeting);
    if (len == 0) {
        return "Fine. Be that way!";
    }
    size_t end = len - 1;

    // Пропустить пробелы в начале
    while (begin < len && (greeting[begin] == ' ' || greeting[begin] == '\t' || greeting[begin] == '\n' || greeting[begin] == '\r')) {
        ++begin;
    }
    // Пропустить пробелы в конце
    while (end > begin && (greeting[end] == ' ' || greeting[end] == '\t' || greeting[end] == '\n' || greeting[end] == '\r')) {
        --end;
    }

    // Если строка пустая после обрезки пробелов
    if (end < begin) {
        return "Fine. Be that way!";
    }

    // выделить память для очищенной строки и скопировать
    size_t cleared_len = end - begin + 1;
    char *cleared = (char *)malloc(cleared_len + 1);
    if (!cleared)
        return "Whatever."; // если не удалось выделить память

    memcpy(cleared, &greeting[begin], cleared_len);
    cleared[cleared_len] = '\0';

    // Проверяем условие с вопросительным знаком
    size_t cleared_len_nonull = strlen(cleared);
    bool has_letters = containsLetters(cleared);
    bool shouting = isUpper(cleared) && has_letters;

    char *result = NULL;

    if (cleared[cleared_len_nonull - 1] == '?') {
        if (shouting) {
            result = "Calm down, I know what I'm doing!";
        } else {
            result = "Sure.";
        }
    } else if (shouting) {
        result = "Whoa, chill out!";
    } else {
        result = "Whatever.";
    }

    free(cleared);
    return result;
}

bool containsLetters(char *input) {
    for (size_t i = 0; i < strlen(input); ++i) {
        if ((input[i] >= 'A' && input[i] <= 'Z') || (input[i] >= 'a' && input[i] <= 'z')) {
            return true;
        }
    }
    return false;
}

bool isUpper(char *input) {
    for (size_t i = 0; i < strlen(input); ++i) {
        if (input[i] >= 'a' && input[i] <= 'z') {
            return false;
        }
    }
    return true;
}