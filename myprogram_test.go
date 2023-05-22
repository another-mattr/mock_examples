package mypackage

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestMyProgramReal(t *testing.T) {
	// Instantiate MyProgram with the RealDependency (no mocking necessary)
	dep := RealDependency{}
	mp := MyProgram{SomeDependency: dep}

	// Check console for the output
	fmt.Println(mp.Action("hey"))
}

func TestMyProgramMock(t *testing.T) {
	//instantiate our MockDependency
	myMock := MockDependency{}
	//Register inputs/outputs for our mock
	myMock.On("Do", "foo").Return("bar")
	myMock.On("Do", "baz").Return("bozo")
	myMock.On("Do", mock.Anything).Return("wildcard!!!")
	//Pass our mock as a dependency to our program
	mp := MyProgram{SomeDependency: &myMock}

	//Exercise the Action function (check console for results)
	fmt.Println(mp.Action("foo"))
	fmt.Println(mp.Action("baz"))
	fmt.Println(mp.Action("sup?"))
}
