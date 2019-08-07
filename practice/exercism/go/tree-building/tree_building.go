package tree

import (
	"fmt"
	"sort"
)

// Record stores its own ID the its Parent ID
type Record struct {
	ID     int
	Parent int
}

// Node is the start of a branch
type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	nodes := make(map[int]*Node)
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	for i := range records {
		if records[i].ID != i {
			return nil, fmt.Errorf("unexpected record id: %d, expected: %d", records[i].ID, i)
		}
		if i == 0 && records[i].Parent != 0 {
			return nil, fmt.Errorf("root node should not have a parent (%d)", records[i].Parent)
		}
		if i != 0 && records[i].Parent >= i {
			return nil, fmt.Errorf("parent id (%d) should be lower than its own id (%d)", records[i].Parent, i)
		}
		nodes[i] = &Node{ID: i}
		if i != 0 {
			nodes[records[i].Parent].Children = append(nodes[records[i].Parent].Children, nodes[i])
		}
	}
	return nodes[0], nil
}
