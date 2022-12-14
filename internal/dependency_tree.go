package internal

import (
	"github.com/lkeix/go-concurrency-scheduler/concurrency"
)

type Node struct {
	Chan     chan bool
	Chans    []chan bool
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

func (dt *DependenceTree) Insert(from *concurrency.Executor, tos ...*concurrency.Executor) {
	child := &Node{
		Executor: from,
		Children: make([]*Node, 0),
		Chan:     make(chan bool),
	}

	dt.Place[from] = child

	if tos == nil {
		child.Parent = dt.Place[dt.root]
		dt.Place[dt.root].Children = append(dt.Place[dt.root].Children, child)
		return
	}

	child.Chans = make([]chan bool, len(tos))

	for i := 0; i < len(tos); i++ {
		_, ok := dt.Place[tos[i]]
		if !ok {
			panic("parent func doesn't exist")
		}

		child.Parent = dt.Place[tos[i]]
		child.Chans[i] = dt.Place[tos[i]].Chan
		dt.Place[tos[i]].Children = append(dt.Place[tos[i]].Children, child)
	}
}
