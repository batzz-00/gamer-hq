package main

import (
	"github.com/google/uuid"
)

type Location struct {
	ID        uuid.UUID
	Name      string
	Entities  []Entity
	Connected []*ConnectedLocation
	Position  Position
}

type NoLocation struct {
}

func (n *NoLocation) Error() string {
	return "Location does not exist"
}

type Position struct {
	X int
	Y int
}

type ConnectedLocation struct {
	Location  *Location
	Direction CompassDirection
}

type CompassDirection string

var (
	NORTH CompassDirection = "NORTH"
	WEST  CompassDirection = "WEST"
	SOUTH CompassDirection = "SOUTH"
	EAST  CompassDirection = "EAST"
)

// type Comparable interface {
// 	Comparator() interface{}
// }

func (location *Location) Comparator() interface{} {
	return location.ID
}

func (location *Location) GetResourceByResourceID(resourceID string) *Resource {
	for _, entity := range location.Entities {
		for i := range entity.Resources {
			if entity.Resources[i].ID.String() == resourceID {
				return &entity.Resources[i]
			}
		}
	}
	return nil
}

func SerializeConnectedLocations(connectedLocations []*ConnectedLocation) []interface{} {
	var connectedLocationSlice []interface{}

	for _, connectedLocation := range connectedLocations {
		connectedLocationSlice = append(connectedLocationSlice, SerializeConnectedLocation(connectedLocation))
	}

	return connectedLocationSlice
}

func SerializeConnectedLocation(connectedLocation *ConnectedLocation) map[string]interface{} {
	object := make(map[string]interface{})

	if connectedLocation.Location != nil {
		object["locationId"] = connectedLocation.Location.ID.String()
	}
	object["direction"] = string(connectedLocation.Direction)

	return object
}

func SerializeLocations(locations []*Location) []interface{} {
	var locationSlice []interface{}

	for _, location := range locations {
		locationSlice = append(locationSlice, SerializeLocation(location))
	}

	return locationSlice
}

func SerializeLocation(location *Location) map[string]interface{} {
	object := make(map[string]interface{})
	if location == nil {
		return nil
	}
	object["id"] = location.ID.String()
	object["name"] = location.Name
	object["entities"] = SerializeEntities(location.Entities)
	object["position"] = SerializePosition(location.Position)
	object["connected"] = SerializeConnectedLocations(location.Connected)

	return object
}

func SerializePosition(position Position) map[string]interface{} {
	object := make(map[string]interface{})

	object["x"] = position.X
	object["y"] = position.Y

	return object
}
