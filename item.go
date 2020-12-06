package main

import (
	"sync"

	"github.com/google/uuid"
)

type Item struct {
	ID   uuid.UUID
	Name string
}

type Stack struct {
	Stack []Item
}

type Inventory struct {
	Items []Stack
	Lock  sync.Mutex
}

type Addable interface {
	Add(item Item)
	AddMultiple(item []Item)
}

func (inventory *Inventory) Add(item Item) {
	inventory.Lock.Lock()
	defer inventory.Lock.Unlock()
	for i := range inventory.Items {
		if inventory.Items[i].Stack[0].Name == item.Name {
			inventory.Items[i].Stack = append(inventory.Items[i].Stack, item)
			return
		}
	}
	inventory.Items = append(inventory.Items, Stack{Stack: []Item{item}})
}

func (inventory *Inventory) AddMultiple(items []Item) {
	if len(items) == 0 {
		return
	}
	inventory.Lock.Lock()
	defer inventory.Lock.Unlock()
	for i := range inventory.Items {
		if inventory.Items[i].Stack[0].Name == items[0].Name {
			inventory.Items[i].Stack = append(inventory.Items[i].Stack, items...)
			return
		}
	}
	inventory.Items = append(inventory.Items, Stack{Stack: items})
}

func (inventory *Inventory) AddToInventory(item Item) {
	inventory.Lock.Lock()
	inventory.Add(item)
	defer inventory.Lock.Unlock()
}

func SerializeInventory(inventory *Inventory) []interface{} {
	stacks := []interface{}{}
	for _, item := range inventory.Items {
		stackMap := make(map[string]interface{})
		stackMap["name"] = item.Stack[0].Name
		stackMap["amount"] = len(item.Stack)
		stacks = append(stacks, stackMap)
	}
	return stacks
}
