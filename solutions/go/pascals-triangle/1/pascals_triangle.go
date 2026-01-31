package pascal

// import "fmt" 

func Triangle(n int) [][]int {
    res := make([][]int, 0)
    res = append(res, []int{1})
	for i := 1; i < n; i++ {
        data := make([]int, 0)
        data = append(data, 1)
        for j := 0; j < i - 1; j++ {
            data = append(data, res[i - 1][j] + res[i - 1][j + 1])
        }
        data = append(data, 1)
        // fmt.Printf("%d\n", n)
        // for _, x := range data {
        //     fmt.Printf("%d, ", x)
        // }
        // fmt.Println("")
        res = append(res, data)
    }

    return res
}
