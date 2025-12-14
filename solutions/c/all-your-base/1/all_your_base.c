#include "all_your_base.h"
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include <errno.h>

int toDecimal(int16_t input_base, int8_t *digits, size_t input_length);
int8_t* fromDecimal(int16_t output_base, int input_number, size_t *output_length);

size_t rebase(int8_t *digits, int16_t input_base, int16_t output_base, size_t input_length) {
    if (input_base < 2 || output_base < 2 || input_length == 0)
        return 0;
    
    int decimal_value = toDecimal(input_base, digits, input_length);
    if (decimal_value < 0)
        return 0;
    size_t output_length;
    int8_t *output = fromDecimal(output_base, decimal_value, &output_length);
    if (!output) 
        return 0;

    memcpy(digits, output, output_length * sizeof(int8_t));
    free(output);
    return output_length;
}

int toDecimal(int16_t input_base, int8_t *digits, size_t input_length) {
    if (input_length == 0) return 0;
    int res = 0;
    
    size_t i = 0;
    while (i < input_length && digits[i] == 0) {
        i++;
    }

    if (i == input_length) return 0;
    for (; i < input_length; i++) {
        if (digits[i] < 0 || digits[i] >= input_base) {
            errno = EINVAL; // Установим код ошибки для некорректных цифр
            return -1;
        }
        res = res * input_base + digits[i]; // Умножаем на основание и добавляем цифру
    }
    
    return res;
}

int8_t* fromDecimal(int16_t output_base, int input_number, size_t *output_length) {
    if (input_number == 0) {
        *output_length = 1; // Should return one digit for zero
        int8_t *res = malloc(*output_length * sizeof(int8_t));
        if (!res) return NULL; // Check memory allocation
        res[0] = 0; // Representing zero
        return res;
    }
    int8_t *res = malloc(32 * sizeof(int8_t));
    if (!res) return NULL; 

    size_t length = 0;
    
    while (input_number > 0) {
        res[length++] = input_number % output_base; // Add remainder
        input_number /= output_base; // Integer division
    }

    for (size_t i = 0; i < length / 2; i++) {
        int8_t temp = res[i];
        res[i] = res[length - i - 1];
        res[length - i - 1] = temp;
    }

    size_t non_zero_index = 0;
    while (non_zero_index < length && res[non_zero_index] == 0) {
        non_zero_index++;
    }

    if (non_zero_index == length) { // If all are zeros
        *output_length = 1;
        res[0] = 0; // Ensure to return single zero in this case
    } else {
        // Return the valid portion of the array
        *output_length = length;  
        res = realloc(res, *output_length * sizeof(int8_t)); // Resize to actual length
    }
    return res;
}