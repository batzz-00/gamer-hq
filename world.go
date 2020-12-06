package main

type World struct {
	Locations []*Location
	Groups    []*Group
}

func NewWorld(locations []*Location, Groups []*Group) World {
	return World{
		Groups:    Groups,
		Locations: locations,
	}
}

// func (world *World) SendToScout() {
// 	skill := &Scouting{}
// 	member := world.Groups[0].GetRandomMemberWithSkill(skill)
// 	if member == nil {
// 		fmt.Println("Everyone in this group is busy, or does not have the skill to scout.")
// 		return
// 	}

// 	location := world.GetUnknownLocation(world.Groups[0])
// 	if location == nil {
// 		fmt.Println("No locations left to discover")
// 		return
// 	}

// 	DoAction(member, skill)
// }

func (world *World) GetLocationByName(name string) *Location {
	for _, location := range world.Locations {
		if location.Name == name {
			return location
		}
	}
	return nil
}

func (world *World) GetLocationByID(ID string) *Location {
	for _, location := range world.Locations {
		if location.ID.String() == ID {
			return location
		}
	}
	return nil
}

func (world *World) AddGroup(group *Group) {
	group.World = world
	world.Groups = append(world.Groups, group)
}

func (world *World) GetUnknownLocation(group *Group) *Location {
	for _, worldLocation := range world.Locations {
		known := false
		for _, groupLocation := range group.AllKnownOrTargetedLocations() {
			if groupLocation.Name == worldLocation.Name {
				known = true
				break
			}
		}
		if !known {
			return worldLocation
		}
	}
	return nil
}

func SerializeWorld(world *World) map[string]interface{} {
	worldMap := make(map[string]interface{})

	worldMap["locations"] = SerializeLocations(world.Locations)

	return worldMap
}
