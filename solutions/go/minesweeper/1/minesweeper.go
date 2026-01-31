package minesweeper

// Annotate returns an annotated board
func Annotate(board []string) []string {
	if len(board) == 0 {
        return []string{}
    }
    rows := len(board)
    cols := len(board[0])
    result := make([]string, rows)
    directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
    for i := 0; i < rows; i++ {
        line := make([]byte, 0, cols)
        for j := 0; j < cols; j++ {
            if board[i][j] == '*' {
                line = append(line, '*')
            } else {
                count := 0
                for _, d := range(directions) {
                    new_i := i + d[0]
                    new_j := j + d[1]
                    if new_i >= 0 && new_i < rows && new_j >= 0 && new_j < cols && board[new_i][new_j] == '*' {
                        count++
                    } 
                }
                if count > 0 {
                    line = append(line, byte(count) + '0')
                } else {
                    line = append(line, ' ')
                }
            }
        }
        result[i] = string(line)
    }
    return result
}
