package internal

import (
	"fmt"
	"testing"

	"github.com/lkeix/go-concurrency-scheduler/concurrency"
)

type a struct{}

func Newa() concurrency.Executor {
	return &a{}
}

func (ea *a) Exec() {
	fmt.Println("aa")
}

func (ea *a) Name() string {
	return ""
}

func TestInsertAtRootChild(t *testing.T) {
	dtree := NewDepsTree()
	a := Newa()
	dtree.Insert(&a)
	if *dtree.Tree.Children[0].Executor != a {
		t.Errorf(`
    Error: cannot insert executor at root child
      expected: %v
      Actual: %v
    `, a, dtree.Tree.Children[0].Executor)
	}
}

func TestPlace(t *testing.T) {
	dtree := NewDepsTree()
	a1 := Newa()
	a2 := Newa()
	a3 := Newa()

	dtree.Insert(&a1)
	if *dtree.Place[&a1].Executor != a1 {
		t.Errorf(`
    Error: cannot insert executor at root child
      expected: %v
      Actual: %v
    `, a1, dtree.Place[&a1].Executor)
	}

	dtree.Insert(&a2, &a1)
	dtree.Insert(&a3)

	if *dtree.Place[&a1].Children[0].Executor != a2 {
		t.Errorf(`
    Error: cannot insert a2 executor after a1
      expected: %v
      Actual: %v
    `, a2, dtree.Place[&a1].Children[0].Executor)
	}
}

func TestInsertAtSomeExecutorChild(t *testing.T) {
	dtree := NewDepsTree()
	a1 := Newa()
	a2 := Newa()
	a3 := Newa()
	dtree.Insert(&a1)
	dtree.Insert(&a2, &a1)
	dtree.Insert(&a3)

	if *dtree.Tree.Children[0].Children[0].Executor != a2 {
		t.Errorf(`
    Error: cannot insert executor at root child
      expected: %v
      Actual: %v
    `, a2, dtree.Tree.Children[0].Children[0].Executor)
	}

	if *dtree.Tree.Children[1].Executor != a3 {
		t.Errorf(`
    Error: cannot insert executor at root child
      expected: %v
      Actual: %v
    `, a3, dtree.Tree.Children[1].Executor)
	}
}
