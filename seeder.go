package main

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

func GenerateResourceTypes() []Resource {
	resources := []Resource{}
	resources = append(resources, NewResource("Oak", skills["Woodcutting"], rand.Intn(10)))
	resources = append(resources, NewResource("Willow", skills["Woodcutting"], rand.Intn(10)))
	resources = append(resources, NewResource("Coal", skills["Mining"], rand.Intn(10)))
	resources = append(resources, NewResource("Iron", skills["Mining"], rand.Intn(10)))
	resources = append(resources, NewResource("Stone", skills["Mining"], rand.Intn(10)))
	return resources
}

func CopyResource(resource Resource) Resource {
	return Resource{Name: resource.Name, Skill: resource.Skill, Amount: rand.Intn(10), ID: uuid.New()}
}

func GenerateEntities() []Entity {
	resources := GenerateResourceTypes()
	entities := []Entity{}

	entities = append(entities, Entity{Name: "Oak Tree", Resources: []Resource{CopyResource(resources[0])}})
	entities = append(entities, Entity{Name: "Willow Tree", Resources: []Resource{CopyResource(resources[1])}})
	entities = append(entities, Entity{Name: "Coal Mine", Resources: []Resource{CopyResource(resources[2])}})
	entities = append(entities, Entity{Name: "Iron Mine", Resources: []Resource{CopyResource(resources[3])}})
	entities = append(entities, Entity{Name: "Rocks", Resources: []Resource{CopyResource(resources[4])}})

	entitiesToKeep := []Entity{}
	for _, entity := range entities {
		if val := rand.Intn(2); val == 1 {
			entitiesToKeep = append(entitiesToKeep, entity)
		}
	}
	return entitiesToKeep
}

func GenerateWorldWithLocations() World {

	locations := []*Location{}
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			id := uuid.New()
			Entities := GenerateEntities()
			locations = append(locations, &Location{ID: id, Entities: Entities, Name: fmt.Sprintf("location %d-%d", x, y), Position: Position{X: x, Y: y}})
		}
	}

	world := World{Locations: locations}
	DoWorldPathing(&world)
	return world
}

func DoWorldPathing(world *World) {
	for _, location := range world.Locations {
		location.Connected = GetConnectedLocations(world, location.Position.X, location.Position.Y)
	}
}

func GetConnectedLocations(world *World, x int, y int) []*ConnectedLocation {
	connected := []*ConnectedLocation{GetLocationByPos(world, x, y+1, NORTH), GetLocationByPos(world, x, y-1, SOUTH), GetLocationByPos(world, x-1, y, EAST), GetLocationByPos(world, x+1, y, WEST)}

	nonNilLocation := []*ConnectedLocation{}
	for _, connected := range connected {
		if connected != nil {
			nonNilLocation = append(nonNilLocation, connected)
		}
	}
	return nonNilLocation
}

func GetLocationByPos(world *World, x int, y int, direction CompassDirection) *ConnectedLocation {
	for _, location := range world.Locations {
		if location.Position.X == x && location.Position.Y == y {
			return &ConnectedLocation{Direction: direction, Location: location}
		}
	}
	return nil
}
