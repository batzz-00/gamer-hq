package main

import "syscall/js"

type jsWrapper struct {
	World  *World
	Events map[string]js.Value
	Doc    js.Value
}

func newJsWrapper(World *World, Doc js.Value) jsWrapper {
	return jsWrapper{
		World:  World,
		Doc:    Doc,
		Events: make(map[string]js.Value),
	}
}
