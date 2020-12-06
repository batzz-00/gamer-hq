package main

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type Groupable interface {
	Join(Group *Group)
}

type Group struct {
	ID      uuid.UUID
	Name    string
	Mutex   sync.Mutex
	Members []*Human
	World   *World
}

type UnknownLocation struct {
}

func (u *UnknownLocation) Error() string {
	return "Location is not known!"
}

func (group *Group) GetKnownLocations() []DiscoveredLocation {
	locations := []DiscoveredLocation{}
	for _, member := range group.Members {
		for _, location := range member.Discovered {
			locations = append(locations, location)
		}
	}
	return locations
}

func (group *Group) AddMember(human *Human) {
	group.Mutex.Lock()
	defer group.Mutex.Unlock()
	group.Members = append(group.Members, human)
	human.Group = group
}

func (group *Group) GetRandomMemberWithSkill(skill Actionable) *Human {
	if len(group.Members) == 0 {
		return nil
	}

	group.Mutex.Lock()
	defer group.Mutex.Unlock()

	for _, member := range group.Members {
		if !IsBusy(member) && HasSkill(member, skill) {
			return member
		}
	}
	return nil
}

func (group *Group) GetMemberByName(Name string) *Human {
	for _, member := range group.Members {
		if member.Name == Name {
			return member
		}
	}
	return nil
}

func (group *Group) GetMemberByNameAndSkill(Name string, skill Actionable) (*Human, error) {
	for _, member := range group.Members {
		if member.Name == Name {
			if HasSkill(member, skill) {
				return member, nil
			}
			return nil, NewMissingSkill(member, skill)
		}
	}
	return nil, errors.New("No members with given name")
}

func (group *Group) GetRandomMember() *Human {
	if len(group.Members) == 0 {
		return nil
	}

	group.Mutex.Lock()
	defer group.Mutex.Unlock()

	for _, member := range group.Members {
		if !IsBusy(member) {
			return member
		}
	}
	return nil
}

func (group *Group) AllKnownOrTargetedLocations() []*Location {
	uniqueLocations := make(map[string]*Location)
	for _, member := range group.Members {
		for _, knownLocation := range member.Discovered {
			uniqueLocations[knownLocation.Location.Name] = knownLocation.Location
		}
		if member.Target != nil {
			uniqueLocations[member.Target.Name] = member.Target
		}
	}

	locations := []*Location{}
	for _, location := range uniqueLocations {
		locations = append(locations, location)
	}
	return locations
}

func (group *Group) KnowsLocation(id string) *Location {
	for _, location := range group.AllKnownOrTargetedLocations() {
		if location.ID.String() == id {
			return location
		}
	}
	return nil
}
