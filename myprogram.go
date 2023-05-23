package mypackage

import "fmt"

type MyProgram struct {
	SomeDependency DoThing
}

func (mp MyProgram) Action(blah string) string {
	word := mp.SomeDependency.Do(blah)
	phrase := fmt.Sprintf("%s goodbye", word)
	return phrase
}
