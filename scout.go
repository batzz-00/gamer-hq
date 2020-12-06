package main

import (
	"fmt"
	"syscall/js"
	"time"
)

type Scouting struct {
	TargetLocation   *Location
	TargetLocationID string
	Human            *Human
}

func (scouting *Scouting) Parse(human *Human, args []js.Value) (Actionable, error) {
	fmt.Println("parse")
	scouting.Human = human
	return scouting, nil
}

// Do(human *Human, location *Location)
func (scouting *Scouting) PreAction() error {
	fmt.Println(scouting.Human)
	scouting.TargetLocation = scouting.Human.Group.World.GetUnknownLocation(scouting.Human.Group)
	scouting.Human.Target = scouting.TargetLocation
	scouting.Human.SetStatus(BUSY, scouting)
	return nil
}

func (scouting *Scouting) PostAction() error {
	scouting.Human.Discover(scouting.TargetLocation)
	scouting.Human.Target = nil
	scouting.Human.SetStatus(IDLE, nil)
	return nil
}

func (scouting *Scouting) Type() Skill {
	return Skill{
		Name: "Scouting",
	}
}

func (scouting *Scouting) Time() time.Duration {
	return time.Millisecond * 50
}

func (scouting *Scouting) Check() error {
	return nil
}
