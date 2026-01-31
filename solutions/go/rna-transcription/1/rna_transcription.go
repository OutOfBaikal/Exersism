package strand

func ToRNA(dna string) string {
	res := make([]byte, 0)
    for i := 0; i < len(dna); i++ {
        if dna[i] == 'G' {
            res = append(res, 'C')
        } else if dna[i] == 'C' {
            res = append(res, 'G')
        } else if dna[i] == 'T' {
            res = append(res, 'A')
        } else if dna[i] == 'A' {
            res = append(res, 'U')
        }
    }
    return string(res)
}
