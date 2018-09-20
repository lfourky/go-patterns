package persistence

import "fmt"

type MockPersister struct {
}

// Persist (in this case) mocks an actual persisting implementation
func (m MockPersister) Persist(args ...interface{}) error {
	fmt.Println("Persisting! Content:", args)
	return nil
}
