package mypackage

import (
	"fmt"
	"github.com/stretchr/testify/mock"
)

type DoThing interface {
	Do(foo string) string
}

type RealDependency struct{}

func (*RealDependency) Do(foo string) string {
	return fmt.Sprintf("REAL Dependency was called with arg: %s", foo)
}

type MockDependency struct {
	mock.Mock
}

func (m *MockDependency) Do(foo string) string {
	args := m.Called(foo)
	return args.String(0)
}
