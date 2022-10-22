package schedulre

import (
	"sync"

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
	wg := &sync.WaitGroup{}
	walk(s.dependencyTree.Tree, wg)
	wg.Wait()
}

func walk(n *internal.Node, wg *sync.WaitGroup) {
	for _, child := range n.Children {
		executor := *child.Executor
		if len(child.Children) == 0 {
			wg.Add(1)
			go func() {
				executor.Exec()
				wg.Done()
			}()
			continue
		}

		go func(children []*internal.Node) {
			if len(children) > 0 {
				wg.Add(1)
				executor.Exec()
				wg.Done()
			}
			walk(n, wg)
		}(child.Children)
	}
}
