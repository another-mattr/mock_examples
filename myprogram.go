package mypackage

type MyProgram struct {
	Foo DoThing
}

func (mp MyProgram) Action(blah string) string {
	return mp.Foo.Do(blah)
}
