package gameoflife

func Tick(matrix [][]int) [][]int {
    dirs := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
    n := len(matrix)
    if n == 0 {
        return matrix
    }
    m := len(matrix[0])
    if m == 0 {
        return matrix
    }
    next := make([][]int, n)
    for i := range next {
        next[i] = make([]int, m)
    }
	for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            live := 0
            for _, d := range dirs {
                ni, nj := i + d[0], j + d[1]
                if ni >= 0 && ni < n && nj >= 0 && nj < m {
                    live += matrix[ni][nj]
                }
            }
            if matrix[i][j] == 1 {
                if live == 2 || live == 3 {
                    next[i][j] = 1
                }
            } else {
                if live == 3 {
                    next[i][j] = 1
                }
            }
        }
    }

    return next
}
