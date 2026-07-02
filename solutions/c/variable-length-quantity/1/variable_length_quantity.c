#include "variable_length_quantity.h"

#define MIN(a, b) (((a) < (b))?(a):(b))

int encode(const uint32_t *integers, size_t integers_len, uint8_t *output)
{
   // write to `output`, return final output's length
   // `output` buffer should be enough to hold the full result
    size_t idx = 0;
    for (size_t i = 0; i < integers_len; ++i) {
        uint32_t current_int = integers[i]; 
        uint8_t bytes[5] = {0};  // Максимум 5 байт для uint32_t
        int byte_count = 0;
        
        do {
            bytes[byte_count++] = current_int & 0x7F;
            current_int >>= 7;
        } while (current_int > 0);
        
        // Теперь добавляем флаги продолжения (0x80) ко всем байтам, кроме последнего
        for (int j = byte_count - 1; j >= 0; --j) {
            if (j > 0) {
                bytes[j] |= 0x80;  // Флаг продолжения
            }
            output[idx++] = bytes[j];
        }
    }

    return idx;
}

int decode(const uint8_t *bytes, size_t buffer_len, uint32_t *output)
{
   // write to `output`, return final output's length
   // return -1 if error
   // `output` buffer should be enough to hold the full result
    size_t idx_in = 0;
    size_t idx_out = 0;

     while(idx_in < buffer_len) {
        uint32_t result = 0;
        uint8_t byte;
        int byte_count = 0;
        
        do {
            if (idx_in >= buffer_len) {
                return -1; 
            }
            
            byte = bytes[idx_in++];
            byte_count++;
            // if (byte_count > 5 || (byte_count == 5 && (byte & 0x70) != 0)) {
            //     return -1; 
            // }
            result = (result << 7) | (uint32_t)(byte & 0x7F);
        } while(byte & 0x80); 
        
        output[idx_out++] = result;
    }
    
    return (int)idx_out; 
}
