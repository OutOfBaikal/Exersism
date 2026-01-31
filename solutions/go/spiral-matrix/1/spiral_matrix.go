package spiralmatrix

func SpiralMatrix(size int) [][]int {
	res := make([][]int, size)
    for i := range res {
        res[i] = make([]int, size)
    }
    i_min, i_max, j_min, j_max := 0, size, 0, size
    num := 1
    
    for i_min < i_max && j_min < j_max {
        for j := j_min; j < j_max; j++ {
			res[i_min][j] = num
			num++
		}
		i_min++
        for i := i_min; i < i_max; i++ {
			res[i][j_max-1] = num
			num++
		}
		j_max--
        
        if i_min < i_max {
			for j := j_max - 1; j >= j_min; j-- {
				res[i_max-1][j] = num
				num++
			}
			i_max--
		}

		if j_min < j_max {
			for i := i_max - 1; i >= i_min; i-- {
				res[i][j_min] = num
				num++
			}
			j_min++
		}
    }

    return res
}
