package schedulre

import (
	"errors"
	"sync"

	"github.com/lkeix/go-concurrency-scheduler/concurrency"
	"github.com/lkeix/go-concurrency-scheduler/internal"
)

type Scheduler struct {
	wg             *sync.WaitGroup
	dependencyTree *internal.DependenceTree
}

type AtOnceScheduler struct {
	wg             *sync.WaitGroup
	dependencyTree *internal.DependenceTree
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		wg:             &sync.WaitGroup{},
		dependencyTree: internal.NewDepsTree(),
	}
}

func (s *Scheduler) Insert(child *concurrency.Executor, parents ...*concurrency.Executor) {
	s.dependencyTree.Insert(child, parents...)
}

func (s *Scheduler) Do() {
	walk(s.dependencyTree.Tree, s.wg)
	s.wg.Wait()
}

func NewAtOnceScheduler() *AtOnceScheduler {
	return &AtOnceScheduler{
		wg:             &sync.WaitGroup{},
		dependencyTree: internal.NewDepsTree(),
	}
}

func (s *AtOnceScheduler) Insert(child *concurrency.Executor, parents ...*concurrency.Executor) {
	s.dependencyTree.Insert(child, parents...)
}

func (s *AtOnceScheduler) Do() {
	s.wg.Add(len(s.dependencyTree.Place) - 1)
	s.walk(s.dependencyTree.Tree, s.wg)
	s.wg.Wait()

	defer func() {
		if err := recover(); err != nil {
			panic(errors.New("Exist illegal dependency"))
		}
	}()
}

func (s *AtOnceScheduler) walk(n *internal.Node, wg *sync.WaitGroup) {
	if n.Executor != nil {
		executor := *n.Executor
		go func(wg *sync.WaitGroup) {
			wait(n.Chans)
			executor.Exec()
			wg.Done()
			n.Chan <- true
		}(wg)
	}

	for i := 0; i < len(n.Children); i++ {
		s.walk(n.Children[i], wg)
	}
}

func walk(n *internal.Node, wg *sync.WaitGroup) {
	if n.Executor != nil {
		executor := *n.Executor
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			wait(n.Chans)
			executor.Exec()
			wg.Done()
			n.Chan <- true
		}(wg)
	}

	for i := 0; i < len(n.Children); i++ {
		walk(n.Children[i], wg)
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
