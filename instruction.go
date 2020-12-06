package main

type ActionContext struct {
	Instructions []Instruction
	Location     *Location
	World        *World
	Resource     *Resource
}

type Instruction struct {
	Name string
	Data interface{}
}

type Instructable interface {
	Set(Instruction) error
}
