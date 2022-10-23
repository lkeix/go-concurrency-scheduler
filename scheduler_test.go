package schedulre

import (
	"testing"

	"github.com/lkeix/go-concurrency-scheduler/concurrency"
)

type a struct{}

func Newa() concurrency.Executor {
	return &a{}
}

func (o *a) Exec() {

}

func (o *a) Name() string {
	return ""
}

func TestInsert(t *testing.T) {
	s := New()
	a1 := Newa()
	s.Insert(&a1)
}

func TestDo(t *testing.T) {
	s := New()
	a1 := Newa()
	s.Insert(&a1)
	s.Do()
}
