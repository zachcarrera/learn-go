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
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	nodes := make([]*Node, len(records))
	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.ID == r.Parent {
			return nil, errors.New("Records not in sequence or has a bad parend")
		}
		nodes[i] = &Node{ID: r.ID}
		if i != 0 {
			parent := nodes[r.Parent]
			parent.Children = append(parent.Children, nodes[i])
		}
	}
	return nodes[0], nil
}
