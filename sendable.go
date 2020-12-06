package main

type Sendable interface {
	Send() bool
	Knows(*Location) bool
	Discover(*Location)
}
