package schedulre

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

func (s *Scheduler) Insert(child concurrency.Executor, parents ...concurrency.Executor) {
	s.Insert(child, parents...)
}

func (s *Scheduler) Do() {
	walk(s.dependencyTree.Tree)
}

func walk(n *internal.Node) {
	for i := 0; i < len(n.Children); i++ {
		executor := *n.Children[i].Executor
		go func(i int) {
			wait(n.Chans)
			executor.Exec()
			n.Chan <- true
		}(i)
		walk(n)
	}
}

func wait(chans []chan bool) {
	if len(chans) == 0 {
		return
	}

	ends := make([]bool, 0)

	for len(ends) != len(chans) {
		for i := 0; i < len(chans); i++ {
			select {
			case <-chans[i]:
				ends = append(ends, true)
			default:
				continue
			}
		}
	}
}
