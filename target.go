package main

type TargetInfo struct {
	Target     interface{}
	TargetType string
}

type Targetable interface {
	Type() TargetInfo
}
