#include "high_scores.h"
/// Return the latest score.
int32_t latest(const int32_t *scores, size_t scores_len) {
    int32_t min = INT32_MAX;
    for (size_t i = 0; i < scores_len; ++i) {
        if (scores[i] < min && scores[i] > 0) {
            min = scores[i];
        }
    }
    return min;
}

/// Return the highest score.
int32_t personal_best(const int32_t *scores, size_t scores_len) {
    if (scores_len == 0) return 0;
    int32_t max = scores[0];
    for (size_t i = 1; i < scores_len; ++i) {
        if (scores[i] > max) {
            max = scores[i];
        }
    }
    return max;
}

/// Write the highest scores to `output` (in non-ascending order).
/// Return the number of scores written.
size_t personal_top_three(const int32_t *scores, size_t scores_len,
                          int32_t *output) {    
    const size_t s = 3;
    int32_t max[s];
    for (size_t i = 0; i < s; ++i) {
        max[i] = 0;
    }
    for (size_t i = 0; i < scores_len; ++i) {
        for (size_t j = 0; j < s; ++j) {
            if (max[j] < scores[i]) {
                for (size_t k = s - 1; k > j; --k) {
                    max[k] = max[k - 1];
                }
                max[j] = scores[i];
                break;
            }
        }
    }

    size_t count = scores_len < s ? scores_len : s;
    for (size_t i = 0; i < count; ++i) {
        output[i] = max[i];
    }

    return count;
                          }