package matrix

import (
    "strings"
    "strconv"
    "errors"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {
	rows := strings.Split(s, "\n")
    
    matrix := [][]int{}
    for _, row := range rows {
        cols := strings.Fields(row)
        intRow := []int{}
        for _, col := range cols {
            num, err := strconv.Atoi(col)
            if err != nil {
            	return nil, err    
            }
            intRow = append(intRow, num)
        }
        matrix = append(matrix, intRow)
    }
    n := len(matrix[0])
    for _, x := range matrix{
        if len(x) != n {
            return nil, errors.New("uneven_rows")
        }
    }
    return matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	if len(m) == 0 {
        return [][]int{}
    }
    cols := make([][]int, len(m[0])) // Создаём срез для столбцов
    for i := 0; i < len(m); i++ {
        for j := 0; j < len(m[i]); j++ {
            cols[j] = append(cols[j], m[i][j])
        }
    }
    return cols
}

func (m Matrix) Rows() [][]int {
    if len(m) == 0 {
        return [][]int{}
    }
    rows := make([][]int, len(m))
    for i := 0; i < len(m); i++ {
        for j := 0; j < len(m[i]); j++ {
            rows[i] = append(rows[i], m[i][j])
        }
    }
	return rows
}

func (m Matrix) Set(row, col, val int) bool {
    if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) {
        return false
    }
	m[row][col] = val
    return true
}
