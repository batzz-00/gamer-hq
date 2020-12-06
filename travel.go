package main

import (
	"errors"
	"syscall/js"
	"time"
)

type Travel struct {
	TargetLocation *Location
	Human          *Human
}

func (travel *Travel) Parse(human *Human, args []js.Value) (Actionable, error) {
	locationID := js.Value.String(args[1])
	travel.Human = human
	travel.TargetLocation = travel.Human.Group.World.GetLocationByName(locationID)
	return travel, nil
}

// Do(human *Human, location *Location)
func (travel *Travel) PreAction() error {
	travel.Human.Target = travel.TargetLocation
	travel.Human.SetStatus(BUSY, travel)
	return nil
}

func (travel *Travel) PostAction() error {
	travel.Human.Location = travel.TargetLocation
	travel.Human.Target = nil
	travel.Human.SetStatus(IDLE, nil)
	return nil
}

func (travel *Travel) Type() Skill {
	return Skill{
		Name: "Travelling",
	}
}

func (travel *Travel) Time() time.Duration {
	return time.Millisecond * 1500
}

func (travel *Travel) Check() error {
	location := travel.Human.Group.KnowsLocation(travel.TargetLocation.ID.String())
	if location == nil {
		return errors.New("Attempting to send human to location that is unknown")
	}
	return nil
}
