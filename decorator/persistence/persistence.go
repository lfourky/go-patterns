package persistence

type Persister interface {
	Persist(args ...interface{}) error
}
