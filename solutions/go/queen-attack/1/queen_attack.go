package queenattack

import (
    "math"
    "errors"
)

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if len(whitePosition) != 2 || len(blackPosition) != 2 ||
    	(whitePosition[0] - 'a') > 7 || (whitePosition[1] - '1') > 7 ||
        (blackPosition[0] - 'a') > 7 || (blackPosition[1] - '1') > 7 ||
        (whitePosition[0] == blackPosition[0] && whitePosition[1] == blackPosition[1]) {
        return false, errors.New("INVALID_POSITION")
    }
    if whitePosition[0] == blackPosition[0] || whitePosition[1] == blackPosition[1] {
        return true, nil
    }
    if math.Abs(float64(int(whitePosition[0]) - int(blackPosition[0]))) == math.Abs(float64(int(whitePosition[1]) - int(blackPosition[1]))) {
        return true, nil
    }
    
    return false, nil
}
