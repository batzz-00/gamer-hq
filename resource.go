package main

import (
	"sync"

	"github.com/google/uuid"
)

type Resource struct {
	Name   string
	Skill  Actionable
	Amount int
	ID     uuid.UUID
	Lock   sync.Mutex
}

type EmptyResource struct {
}

func (e *EmptyResource) Error() string {
	return "Resource is empty!"
}

type Gatherable interface {
	Gather(amount int) []Item
}

func (resource *Resource) Gather(amount int) []Item {
	resource.Lock.Lock()
	defer resource.Lock.Unlock()
	diff := 0
	if resource.Amount-amount < 0 {
		diff = resource.Amount
	} else {
		diff = amount
	}

	items := []Item{}
	for i := 1; i <= diff; i++ {
		items = append(items, Item{Name: resource.Name})
	}

	resource.Amount = resource.Amount - diff
	return items
}

func NewResource(name string, skill Actionable, amount int) Resource {
	return Resource{
		Name:   name,
		Skill:  skill,
		Amount: amount,
		ID:     uuid.New(),
	}
}

type Entity struct {
	Name      string
	Resources []Resource
}

func SerializeEntities(entities []Entity) []interface{} {
	var entitySlice []interface{}

	for _, entity := range entities {
		entitySlice = append(entitySlice, SerializeEntity(entity))
	}

	return entitySlice
}

func SerializeEntity(entity Entity) map[string]interface{} {
	object := make(map[string]interface{})

	object["name"] = entity.Name
	object["resources"] = SerializeResources(entity.Resources)
	return object
}

func SerializeResources(resources []Resource) []interface{} {
	var resourceSlice []interface{}

	for _, resource := range resources {
		resourceSlice = append(resourceSlice, SerializeResource(resource))
	}

	return resourceSlice
}

func SerializeResource(resource Resource) map[string]interface{} {
	object := make(map[string]interface{})

	object["name"] = resource.Name
	object["amount"] = resource.Amount
	object["id"] = resource.ID.String()

	object["skill"] = nil
	if resource.Skill != nil {
		object["skill"] = resource.Skill.Type().Name
	}
	// object["skill"] = resource.Skill.Type().Name

	return object
}
