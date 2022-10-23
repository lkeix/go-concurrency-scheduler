package schedulre

import (
	"fmt"
	"testing"

	"github.com/lkeix/go-concurrency-scheduler/concurrency"
)

type (
	a1 struct{}
	a2 struct{}
	a3 struct{}
)

func Newa1() concurrency.Executor {
	return &a1{}
}

func Newa2() concurrency.Executor {
	return &a1{}
}

func Newa3() concurrency.Executor {
	return &a3{}
}

func (o *a1) Exec() {
	fmt.Printf("a1\n")
}

func (o *a2) Exec() {
	fmt.Printf("a2\n")
}

func (o *a3) Exec() {
	fmt.Printf("a3\n")
}

func (o *a1) Name() string {
	return ""
}

func (o *a2) Name() string {
	return ""
}

func (o *a3) Name() string {
	return ""
}

func TestInsert(t *testing.T) {
	s := New()
	a1 := Newa1()
	a2 := Newa2()
	a3 := Newa3()
	s.Insert(&a1)
	s.Insert(&a2, &a1)
	s.Insert(&a3)
}

func TestDo(t *testing.T) {
	s := New()
	a1 := Newa1()
	a2 := Newa2()
	a3 := Newa3()
	s.Insert(&a1)
	s.Insert(&a2, &a1)
	s.Insert(&a3)
	s.Do()
}
