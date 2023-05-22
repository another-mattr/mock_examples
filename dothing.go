package mypackage

import (
	"fmt"
	"github.com/stretchr/testify/mock"
)

type DoThing interface {
	Do(foo string) string
}

type RealDoThing struct{}

func (*RealDoThing) Do(foo string) string {
	return fmt.Sprintf("REAL DoThing %s", foo)
}

type MockDoThing struct {
	mock.Mock
}

func (m *MockDoThing) Do(foo string) string {
	args := m.Called(foo)
	return args.String(0)
}
