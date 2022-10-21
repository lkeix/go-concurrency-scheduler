package concurrency

type Executor interface {
	Name() string
	Exec()
}
