package yacht

func Score(dice []int, category string) int {
    var res int
	switch category {
        case "ones": {
            for _, x := range dice {
                if x == 1 {
                	res += 1
                }
            }
        }
        case "twos": {
            for _, x := range dice {
                if x == 2 {
                	res += 2
                }
            }
        }
        case "threes": {
            for _, x := range dice {
                if x == 3 {
                	res += 3
                }
            }
        }
        case "fours": {
            for _, x := range dice {
                if x == 4 {
                	res += 4
                }
            }
        }
        case "fives": {
            for _, x := range dice {
                if x == 5 {
                	res += 5
                }
            }
        }
        case "sixes": {
            for _, x := range dice {
                if x == 6 {
                	res += 6
                }
            }
        }
        case "full house": {
            dict := make(map[int]int, 0)
            for _, x := range dice {
            	dict[x]++
            }
            var f1, f2 bool
            for _, x := range dict {
                if x >= 3 {
                    f1 = true
                } else if x >= 2 {
                    f2 = true
                }
            }
            if f1 && f2 {
                for _, x := range dice {
            		res += x
            	}
            }
        }
        case "four of a kind": {
            dict := make(map[int]int, 0)
            for _, x := range dice {
            	dict[x]++
            }
            var card int
            for k, x := range dict {
                if x >= 4 {
                    card = k
                }
            }
            res = card * 4
        }
        case "little straight": {
            dict := make(map[int]int, 0)
            for _, x := range dice {
            	dict[x]++
            }
            var count int
            var sum int
            for k, v := range dict {
                if v >= 1 {
                    count++
                    sum += k
                }
            }
            if count == 5 && sum == 15 {
                res = 30
            }
        }
        case "big straight": {
            dict := make(map[int]int, 0)
            for _, x := range dice {
            	dict[x]++
            }
            var count int
            var sum int
            for k, v := range dict {
                if v == 1 {
                    count++
                    sum += k
                }
            }
            if count == 5 && sum == 20 {
                res = 30
            }
        }
        case "choice": {
            for _, x := range dice {
            	res += x
            }
        }
        case "yacht": {
            dict := make(map[int]int, 0)
            for _, x := range dice {
            	dict[x]++
            }
            var card bool
            for _, x := range dict {
                if x >= 5 {
                    card = true
                    break
                }
            }
            if card {
            	res = 50
            }
        }
        default: {
            
        }
    }
    return res
}
