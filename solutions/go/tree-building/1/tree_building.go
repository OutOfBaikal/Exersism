package tree

import (
    "errors"
    "sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
    if len(records) == 0 {
		return nil, nil
	}

	// Step 1: Validate and collect records
	nodeMap := make(map[int]*Node, len(records))
	hasRoot := false

	for _, r := range records {
		// ID must be in [0, len(records)-1]
		if r.ID < 0 || r.ID >= len(records) {
			return nil, errors.New("record ID out of range")
		}

		// Check for duplicate ID
		if nodeMap[r.ID] != nil {
			return nil, errors.New("duplicate record ID found")
		}

		// Parent must be < ID, except for root (ID==0, Parent==0)
		if r.ID == 0 {
			if r.Parent != 0 {
				return nil, errors.New("root node's parent must be 0")
			}
			hasRoot = true
		} else {
			if r.Parent >= r.ID {
				return nil, errors.New("parent ID must be less than node ID for non‑root nodes")
			}
		}

		// Create node
		nodeMap[r.ID] = &Node{ID: r.ID}
	}

	if !hasRoot {
		return nil, errors.New("no root node (ID=0, Parent=0) found")
	}

	// Step 2: Build tree by linking parents and children
	for _, r := range records {
		if r.ID == 0 {
			continue // root has no parent to link
		}

		parentNode := nodeMap[r.Parent]
		if parentNode == nil {
			return nil, errors.New("parent node not found in records")
		}

		parentNode.Children = append(parentNode.Children, nodeMap[r.ID])
	}

	// Step 3: Sort children by ID (once, post‑construction)
	var sortChildren func(*Node)
	sortChildren = func(n *Node) {
		sort.Slice(n.Children, func(i, j int) bool {
			return n.Children[i].ID < n.Children[j].ID
		})
		for _, child := range n.Children {
			sortChildren(child)
		}
	}

	root := nodeMap[0]
	sortChildren(root)

	return root, nil
}
