package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File
// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    res := 0
    for _, occupied := range cb[file] {
        if occupied {
            res++
        }
    }
    return res
	panic("Please implement CountInFile()")
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
     res := 0
    // Adjust rank to be zero-indexed
    if rank < 1 || rank > 8 {
        return 0 // Invalid rank
    }
    for _, file := range cb {
        if file[rank-1] { // rank-1 because slices are zero-indexed
            res++
        }
    }
    return res
	panic("Please implement CountInRank()")
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
     return len(cb) * 8
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    res := 0
    for _, file := range cb {
        for _, occupied := range file {
            if occupied {
                res++
            }
        }
    }
    return res
}
