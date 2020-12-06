package main

type Status string

const (
	BUSY Status = "BUSY"
	IDLE Status = "IDLE"
)

type Statusable interface {
	Status() Status
	SetStatus(Status, Actionable)
}
