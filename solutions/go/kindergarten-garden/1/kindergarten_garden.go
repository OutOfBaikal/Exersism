package kindergarten

import(
    "fmt"
    "strings"
    "sort"
)

// Define the Garden type here.
type Child struct {
    Name   string
    Plants []string
}

type Garden struct {
    Children []Child
}
// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
    for i := 0; i < len(children) - 1; i++ {
        for j := i + 1; j < len(children); j++ {
            if children[i] == children[j] {
                return nil, fmt.Errorf("duplicate name")
            }
        }
    }
    if diagram[0] != '\n' {
        return nil, fmt.Errorf("wrong diagram format")
    }
    for _, v := range(diagram) {
        if (v < 'A' || v > 'Z') && v != '\n' {
            return nil, fmt.Errorf("invalid cup codes")
        }
    }
    
	data := map[byte]string{
        'V': "violets",
        'R': "radishes",
        'C': "clover",
        'G': "grass",
    }
    
    cur_student := make([]string, len(children))
	copy(cur_student, children)
    sort.Strings(cur_student)

    // Clean the diagram string, trim whitespace and split into lines
    lines := strings.Split(strings.TrimSpace(diagram), "\n")
    if len(lines) != 2 {
        return nil, fmt.Errorf("diagram must have exactly two rows")
    }

    row1 := lines[0]
    row2 := lines[1]

    if len(row1) != len(row2) {
        return nil, fmt.Errorf("diagram rows must be of equal length")
    }

    // Each child has 4 plants: 2 from row1 and 2 from row2
    // So total plants per row must be 2 * len(children)
    if len(row1) != 2*len(children) {
        return nil, fmt.Errorf("diagram row length does not match number of children")
    }

    garden := &Garden{
        Children: make([]Child, len(children)),
    }

    for i, childName := range cur_student {
        plants := []string{
            data[row1[2*i]],
            data[row1[2*i+1]],
            data[row2[2*i]],
            data[row2[2*i+1]],
        }
        garden.Children[i] = Child{
            Name:   childName,
            Plants: plants,
        }
    }

    return garden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	for _, x := range(g.Children) {
        if x.Name == child {
            return x.Plants, true
        }
    }
    return []string{}, false
}
