package mypackage

type MyProgram struct {
	SomeDependency DoThing
}

func (mp MyProgram) Action(blah string) string {
	return mp.SomeDependency.Do(blah)
}
