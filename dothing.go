package mypackage

import (
	"fmt"
	"github.com/stretchr/testify/mock"
)

type DoThing interface {
	Do(foo string) string
}

// The _real_ DoThing

type RealDoThing struct{}

func (*RealDoThing) Do(foo string) string {
	return fmt.Sprintf("REAL DoThing %s", foo)
}

// The _mock_ DoThing
type MockDoThing struct {
	mock.Mock
}

func (m *MockDoThing) Do(foo string) string {
	args := m.Called(foo)
	return args.String(0)
}
