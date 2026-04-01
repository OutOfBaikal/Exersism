#include "matching_brackets.h"
#include <stdlib.h>
#include <string.h>

bool is_paired(const char *input) {
    size_t len = strlen(input);
    char *stack = malloc(sizeof(char) * len);
    size_t j = 0;
    char top;
    for (size_t i = 0; i < len; ++i) {
        switch(input[i]) {
            case '[':
            case '{':
            case '(': 
                stack[j++] = input[i];
                break;
                
            case ']' :
                    if (j == 0) {
                        free(stack);
                        return false;
                    }
                    top = stack[j - 1];
                    if (top != '[') {
                        free(stack);
                        return false;
                    }
                    --j;
                    break;
            case '}' :
                    if (j == 0) {
                        free(stack);
                        return false;
                    }
                    top = stack[j - 1];
                    if (top != '{') {
                        free(stack);
                        return false;
                    }
                    --j;
                    break;
            case ')' :
                    if (j == 0) {
                        free(stack);
                        return false;
                    }
                    top = stack[j - 1];
                    if (top != '(') {
                        free(stack);
                        return false;
                    }
                    --j;
                    break;
        }
    }
    free(stack);
    return j == 0;
}