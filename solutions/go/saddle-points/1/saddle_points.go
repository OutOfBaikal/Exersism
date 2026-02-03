package matrix

import (
    "bufio"
    "strings"
    "errors"
    "strconv"
)

// Define the Matrix and Pair types here.
type Matrix struct {
    data [][]int
}

type Pair struct {
    first, second int
}

func New(s string) (*Matrix, error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
    var matrixData [][]int

    for scanner.Scan() {
        line := scanner.Text()
        rowElements := strings.Fields(line)
        row := make([]int, len(rowElements))

        for j, elemStr := range rowElements {
            num, err := strconv.Atoi(elemStr)
            if err != nil {
                return nil, errors.New("Invalid input format")
            }
            row[j] = num
        }

        matrixData = append(matrixData, row)
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return &Matrix {
        data: matrixData,
    }, nil
}

func (m *Matrix) Saddle() []Pair {
	var results []Pair
    nRows := len(m.data)
    if nRows == 0 {
		return results
	}
    nCols := len(m.data[0])

    // Обходим каждый элемент матрицы
    for i := 0; i < nRows; i++ {
        for j := 0; j < nCols; j++ {
            currentValue := m.data[i][j]
            
            // Сначала проверим, что текущий элемент максимальный в своем ряду
            maxInRow := true
            for k := 0; k < nCols; k++ {
                if m.data[i][k] > currentValue {
                    maxInRow = false
                    break
                }
            }
            
            // Потом проверим, что текущий элемент минимальный в своем столбце
            minInCol := true
            for l := 0; l < nRows; l++ {
                if m.data[l][j] < currentValue {
                    minInCol = false
                    break
                }
            }
            
            // Если оба условия выполняются, добавляем координату в результат
            if maxInRow && minInCol {
                results = append(results, Pair{i + 1, j + 1})
            }
        }
    }

    return results
}
