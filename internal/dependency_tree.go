package internal

import "github.com/lkeix/go-concurrency-scheduler/concurrency"

type Node struct {
	Parent   *Node
	Executor *concurrency.Executor
	Children []*Node
}

type DependenceTree struct {
	Tree  *Node
	root  *concurrency.Executor
	Place map[*concurrency.Executor]*Node
}

func NewDepsTree() *DependenceTree {
	var rootExecutor concurrency.Executor
	root := &Node{}
	dt := &DependenceTree{
		Tree:  root,
		Place: make(map[*concurrency.Executor]*Node),
		root:  &rootExecutor,
	}
	dt.Place[&rootExecutor] = root
	return dt
}

func (dt *DependenceTree) Insert(from, to *concurrency.Executor) {
	child := &Node{
		Executor: from,
		Children: make([]*Node, 0),
	}

	dt.Place[from] = child
	if to == nil {
		child.Parent = dt.Place[dt.root]
		dt.Place[dt.root].Children = append(dt.Place[dt.root].Children, child)
		return
	}

	parent, ok := dt.Place[to]
	if !ok {
		panic("parent func doesn't insert")
	}

	child.Parent = parent
	parent.Children = append(parent.Children, child)
}