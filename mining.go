package main

import (
	"fmt"
	"syscall/js"
	"time"
)

type Mining struct {
	TargetLocation *Location
	Resource       *Resource
	Human          *Human
}

func (Mining *Mining) Type() Skill {
	return Skill{
		Name: "Mining",
	}
}

func (mining *Mining) Parse(human *Human, args []js.Value) (Actionable, error) {
	locationID := js.Value.String(args[0])
	resourceID := js.Value.String(args[1])
	mining.Human = human
	location := mining.Human.Group.World.GetLocationByID(locationID)

	mining.Resource = location.GetResourceByResourceID(resourceID)

	mining.TargetLocation = location
	return mining, nil
}

// Do(human *Human, location *Location)
func (mining *Mining) PreAction() error {
	fmt.Println("STARTED MINING")
	mining.Human.Target = mining.TargetLocation
	mining.Human.SetStatus(BUSY, mining)
	return nil
}

func (mining *Mining) PostAction() error {
	fmt.Println("lol")
	AddMultiple(Gather(mining.Resource, 3), &mining.Human.Inventory)
	mining.Human.Target = nil
	mining.Human.SetStatus(IDLE, nil)
	return nil
}

func (Mining *Mining) Time() time.Duration {
	return time.Millisecond * 2500
}

func (mining *Mining) Check() error {
	if mining.Resource.Amount == 0 {
		return &EmptyResource{}
	}
	if mining.Human.IsBusy() {
		return &HumanIsBusy{}
	}
	if mining.Human.Group.KnowsLocation(mining.TargetLocation.ID.String()) == nil {
		return &UnknownLocation{}
	}
	return nil
}

// func (woodcutting *Woodcutting) PreAction() error {
// 	woodcutting.Human.Target = woodcutting.TargetLocation
// 	woodcutting.Human.SetStatus(BUSY, woodcutting)
// 	return nil
// }

// func (woodcutting *Woodcutting) PostAction() error {
// 	AddMultiple(Gather(woodcutting.Resource, 3), &woodcutting.Human.Inventory)
// 	woodcutting.Human.SetStatus(IDLE, nil)
// 	return nil
// }
