package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type Human struct {
	Name          string
	Group         *Group
	skills        []Actionable
	currentAction Actionable
	Location      *Location
	Target        *Location
	Discovered    map[string]DiscoveredLocation
	status        Status
	Lock          sync.Mutex

	Inventory Inventory
}

type HumanIsBusy struct{}

func (h *HumanIsBusy) Error() string {
	return "Human is currently busy with something"
}

type DiscoveredLocation struct {
	Location  *Location
	Timestamp time.Time
}

func NewHuman(name string, skills []Actionable) Human {
	return Human{
		Name:       name,
		skills:     skills,
		Discovered: make(map[string]DiscoveredLocation),
		status:     IDLE,
	}
}

func (human *Human) Join(group *Group) {
	group.Members = append(group.Members, human)
	human.Group = group
}

func (human *Human) Status() Status {
	return human.status
}

func (human *Human) IsBusy() bool {
	return human.status == BUSY
}

func (human *Human) SetStatus(status Status, action Actionable) {
	human.status = status
	human.currentAction = action
}

func (human *Human) AddSkill(skill Actionable) {
	human.skills = append(human.skills, skill)
}

func (human *Human) Skills() []Actionable {
	return human.skills
}

func (human *Human) Knows(location *Location) bool {
	if _, ok := human.Discovered[location.Name]; ok {
		return true
	}
	return false
}

func (human *Human) Discover(location *Location) {
	human.Discovered[location.Name] = DiscoveredLocation{
		Location:  location,
		Timestamp: time.Now(),
	}
}

func SerializeHuman(human *Human) map[string]interface{} {
	statusMap := make(map[string]interface{})
	statusMap["name"] = human.Name
	statusMap["status"] = fmt.Sprintf("%s", human.status)
	statusMap["inventory"] = SerializeInventory(&human.Inventory)
	statusMap["action"] = nil
	statusMap["groupId"] = nil

	if human.Group != nil {
		statusMap["groupId"] = human.Group.ID.String()
	}

	if human.currentAction != nil {
		statusMap["action"] = fmt.Sprintf("%s", human.currentAction.Type().Name)
	}

	discoveredLocations := []DiscoveredLocation{}
	for _, location := range human.Discovered {
		discoveredLocations = append(discoveredLocations, location)
	}

	sort.Slice(discoveredLocations, func(i, j int) bool {
		return discoveredLocations[i].Timestamp.Before(discoveredLocations[j].Timestamp)
	})

	visible := []interface{}{}
	for _, location := range discoveredLocations {
		visible = append(visible, location.Location.Name)
	}

	skills := []interface{}{}

	for _, skill := range human.skills {
		skills = append(skills, skill.Type().Name)
	}

	statusMap["location"] = SerializeLocation(human.Location)
	statusMap["skills"] = skills
	statusMap["discovered"] = visible
	return statusMap
}
