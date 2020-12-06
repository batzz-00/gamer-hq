package main

import (
	"syscall/js"
	"time"
)

type Woodcutting struct {
	TargetLocation *Location
	Resource       *Resource
	Human          *Human
}

func (woodcutting *Woodcutting) Parse(human *Human, args []js.Value) (Actionable, error) {
	locationID := js.Value.String(args[0])
	resourceID := js.Value.String(args[1])
	woodcutting.Human = human
	location := woodcutting.Human.Group.World.GetLocationByID(locationID)

	woodcutting.Resource = location.GetResourceByResourceID(resourceID)

	woodcutting.TargetLocation = location
	return woodcutting, nil
}

// Do(human *Human, location *Location)

// Do(human *Human, location *Location)
func (woodcutting *Woodcutting) PreAction() error {
	woodcutting.Human.Target = woodcutting.TargetLocation
	woodcutting.Human.SetStatus(BUSY, woodcutting)
	return nil
}

func (woodcutting *Woodcutting) PostAction() error {
	AddMultiple(Gather(woodcutting.Resource, 3), &woodcutting.Human.Inventory)
	woodcutting.Human.SetStatus(IDLE, nil)
	return nil
}

func (woodcutting *Woodcutting) Type() Skill {
	return Skill{
		Name: "Woodcutting",
	}
}

func (woodcutting *Woodcutting) Time() time.Duration {
	return time.Millisecond * 1500
}

func (woodcutting *Woodcutting) Check() error {
	if woodcutting.Resource.Amount == 0 {
		return &EmptyResource{}
	}
	if woodcutting.Human.IsBusy() {
		return &HumanIsBusy{}
	}
	if woodcutting.Human.Group.KnowsLocation(woodcutting.TargetLocation.ID.String()) == nil {
		return &UnknownLocation{}
	}
	return nil
}
