package main

import (
	"fmt"
	"miner/events"
	"time"
)

func JoinGroup(groupable Groupable, group *Group) {
	groupable.Join(group)

	fmt.Println("Group Joined")
}

func IsBusy(statusable Statusable) bool {
	return statusable.Status() == BUSY
}

func DoAction(human *Human, skill Actionable) {
	err := skill.Check()
	if err != nil {
		events.Emit("Error", err.Error())
		return
	}

	err = skill.PreAction()
	if err != nil {
		events.Emit("Error", err.Error())
		return
	}

	events.Emit("Member", SerializeHuman(human))
	time.Sleep(skill.Time())

	err = skill.PostAction()
	if err != nil {
		events.Emit("Error", err.Error())
		return
	}

	events.Emit("Member", SerializeHuman(human))
}

func Send(human *Human, location *Location) {
	human.Location = location
	time.Sleep(time.Millisecond * 2500)
	events.Emit("Member", SerializeHuman(human))
	fmt.Println(fmt.Sprintf("Moved %s to %s", human.Name, location.Name))
}

func Add(Item Item, Inventory Addable) {
	Inventory.Add(Item)
}

func AddMultiple(Items []Item, Inventory Addable) {
	Inventory.AddMultiple(Items)
}

func Gather(resource Gatherable, amount int) []Item {
	return resource.Gather(amount)
}

func SendToScout(world *World) {
	skill := &Scouting{}
	member := world.Groups[0].GetRandomMemberWithSkill(skill)
	if member == nil {
		events.Emit("Error", "Everyone in this group is busy, or does not have the skill to scout.")
		return
	}

	location := world.GetUnknownLocation(world.Groups[0])
	if location == nil {
		events.Emit("Error", "No locations left to discover")
		return
	}

	DoAction(member, skill)
}

func Scout(name string, world *World) {
	member, err := world.Groups[0].GetMemberByNameAndSkill(name, skills["Scouting"])

	if err != nil {
		events.Emit("Error", err.Error())
		return
	}

	if member.status == BUSY {
		events.Emit("Error", fmt.Sprintf("%s is busy!", name))
		return
	}

	location := world.GetUnknownLocation(world.Groups[0])
	if location == nil {
		events.Emit("Error", "No locations left to discover")
		return
	}
	fmt.Println("what")
	skill := &Scouting{Human: member, TargetLocation: location}
	fmt.Println(skill)
	DoAction(member, skill)
}
