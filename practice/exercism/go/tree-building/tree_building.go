package tree

import (
	"errors"
	"fmt"
	"reflect"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

func BuildOld(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	root := &Node{}
	todo := []*Node{root}
	n := 1
	for {
		if len(todo) == 0 {
			break
		}
		newTodo := []*Node(nil)
		for _, c := range todo {
			for _, r := range records {
				if r.Parent == c.ID {
					if r.ID < c.ID {
						return nil, errors.New("a")
					} else if r.ID == c.ID {
						if r.ID != 0 {
							return nil, fmt.Errorf("b")
						}
					} else {
						n++
						switch len(c.Children) {
						case 0:
							nn := &Node{ID: r.ID}
							c.Children = []*Node{nn}
							newTodo = append(newTodo, nn)
						case 1:
							nn := &Node{ID: r.ID}
							if c.Children[0].ID < r.ID {
								c.Children = []*Node{c.Children[0], nn}
								newTodo = append(newTodo, nn)
							} else {
								c.Children = []*Node{nn, c.Children[0]}
								newTodo = append(newTodo, nn)
							}
						default:
							nn := &Node{ID: r.ID}
							newTodo = append(newTodo, nn)
						breakpoint:
							for range []bool{false} {
								for i, cc := range c.Children {
									if cc.ID > r.ID {
										a := make([]*Node, len(c.Children)+1)
										copy(a, c.Children[:i])
										copy(a[i+1:], c.Children[i:])
										copy(a[i:i+1], []*Node{nn})
										c.Children = a
										break breakpoint
									}
								}
								c.Children = append(c.Children, nn)
							}
						}
					}
				}
			}
		}
		todo = newTodo
	}
	if n != len(records) {
		return nil, Mismatch{}
	}
	if err := chk(root, len(records)); err != nil {
		return nil, err
	}
	return root, nil
}

func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	err := validateData(records)
	if err != nil {
		return nil, err
	}


	tree := branch(&Node{}, records)
	return tree, nil
}

func validateData(records []Record) error {

	err := rootCheck(records)
	if err != nil {
		return err
	}

	err = duplicateCheck(records)
	if err != nil {
		return err
	}

	err = continuous(records)
	if err != nil {
		return err
	}

	return nil
}

// rootCheck checks for the presence of a valid root node in the records
func rootCheck(records []Record) error {
	for _, r := range records {
		if r.ID == 0 { // root node
			if r.Parent != 0 {
				return errors.New("Root node cannot have a parent id")
			}
			return nil
		}
	}

	return errors.New("No root node found")
}

func duplicateCheck(records []Record) error {

	for ai, av := range records {
		for bi, bv := range records {
			if ai == bi { // skip same record comparison
				continue
			}
			if reflect.DeepEqual(av, bv) {
				return errors.New("Duplicate records found")
			}
		}
	}

	return nil
}

// continuous checks for missing ID in record sequence
func continuous(records []Record) error {

	var highestID int

	for _, r := range records {
		if r.ID > highestID {
			highestID = r.ID
		}
	}

	for i := 0; i < highestID; i++ {
		for _, r := range records {
			if r.ID == i {
				break
			}
			return errors.New("non continuous")
		}
	}

	return nil
}

// orphanCheck checks for orphaned records
func orphanCheck(records []Record) error {

	fmt.Println("########################################")

	for ai, av := range records {

		if av.ID == 0 { // skip root
			continue
		}

		fmt.Println("Looking for parents of", av)
		parent := false
		for bi, bv := range records {
			if ai == bi { // skip same record comparison
				continue
			}
			fmt.Print("\t", bv)
			if av.Parent == bv.ID {
				fmt.Println("\tfound!")
				parent = true
			}
			if parent == true {
				break
			}
		}
		if parent == false {
			fmt.Println("\tNOT found!")
			return errors.New("Found orphaned record")
		}
	}

	return nil
}

// branch returns the parent with child nodes attached
func branch(root *Node, records []Record) *Node {

	if len(records) <  1 {
		return root
	}

	for _, r := range records {
		if r.Parent == root.ID && r.ID != 0 {
			root.Children  = append(root.Children, &Node{ID: r.ID})
			root.Children = sortChildren(root.Children)
		}
	}

	// recursively add branches
	for _, c := range root.Children {
		branch(c, records)
	}

	return root
}

// sortChildren does a bubble sort a slice of Node based on Node.ID
func sortChildren(children []*Node) []*Node {

	l := len(children)

	for i := 0; i < l; i++ {
		for j := 0; j < (l - 1 - i); j++ {
			if children[j].ID > children[j+1].ID {
				children[j], children[j+1] = children[j+1], children[j]
			}
		}
	}

	return children
}

func chk(n *Node, m int) (err error) {
	if n.ID > m {
		return fmt.Errorf("z")
	} else if n.ID == m {
		return fmt.Errorf("y")
	} else {
		for i := 0; i < len(n.Children); i++ {
			err = chk(n.Children[i], m)
			if err != nil {
				return
			}
		}
		return
	}
}
