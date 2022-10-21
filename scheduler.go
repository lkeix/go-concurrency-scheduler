package gocs

import (
	"github.com/lkeix/go-concurrency-scheduler/concurrency"
	"github.com/lkeix/go-concurrency-scheduler/internal"
)

type Scheduler struct {
	dependencyTree *internal.DependenceTree
}

func New() *Scheduler {
	return &Scheduler{
		dependencyTree: internal.NewDepsTree(),
	}
}

func (s *Scheduler) Insert(child, parent concurrency.Fn) {
	s.Insert(child, parent)
}

func (s *Scheduler) Do() {
	walk(s.dependencyTree.Tree)
}

func walk(n *internal.Node) {
	for _, child := range n.Children {
		executor := *child.Executor
		if len(child.Children) == 0 {
			go executor.Exec()
			continue
		}

		go func(children []*internal.Node) {
			if len(children) > 0 {
				executor.Exec()
			}
			walk(n)
		}(child.Children)
	}
}
