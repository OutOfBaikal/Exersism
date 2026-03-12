package lineup

import "fmt"

func Format(name string, number int) string {
	end := number % 10
    tail := number % 100
    if tail != 11 && tail != 12 && tail != 13 {
        switch end {
            case 1:
            	return fmt.Sprintf("%s, you are the %dst customer we serve today. Thank you!", name, number)
            case 2:
            	return fmt.Sprintf("%s, you are the %dnd customer we serve today. Thank you!", name, number)
            case 3:
            	return fmt.Sprintf("%s, you are the %drd customer we serve today. Thank you!", name, number)
        	default:
            	return fmt.Sprintf("%s, you are the %dth customer we serve today. Thank you!", name, number)
        }
    }
    return fmt.Sprintf("%s, you are the %dth customer we serve today. Thank you!", name, number)
}
