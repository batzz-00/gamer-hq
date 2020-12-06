package main

import (
	"errors"
	"fmt"
	"miner/events"
	"sort"
	"syscall/js"

	"github.com/google/uuid"
)

var skills map[string]Actionable

func init() {
	skills = make(map[string]Actionable)
	skills["Scouting"] = &Scouting{}
	skills["Mining"] = &Mining{}
	skills["Woodcutting"] = &Woodcutting{}
	skills["Travel"] = &Travel{}
}

func main() {

	done := make(chan struct{})
	// Mining := Mining{}

	// Mine := Location{Name: "Mine", Skill: &Mining}

	// Mine2 := Location{Name: "Mine 2", Skill: &Mining}

	Batu := NewHuman("Batu", []Actionable{skills["Scouting"], skills["Mining"]})
	Gazza := NewHuman("Gartosh Gazza", []Actionable{skills["Scouting"]})
	Calum := NewHuman("Calum Calcium", []Actionable{})

	Group1 := Group{ID: uuid.New(), Name: "Group 1"}
	Group1.AddMember(&Batu)
	Group1.AddMember(&Gazza)
	Group1.AddMember(&Calum)

	// World := NewWorld([]*Location{&Mine, &Mine2}, )
	World := GenerateWorldWithLocations()
	World.AddGroup(&Group1)
	// js
	doc := js.Global().Get("document")

	js.Global().Set("game", make(map[string]interface{}))
	module := js.Global().Get("game")

	jsWrapper := newJsWrapper(&World, doc)
	GenerateWorldWithLocations()

	module.Set("status", js.FuncOf(jsWrapper.status))
	module.Set("scout", js.FuncOf(jsWrapper.scout))
	module.Set("send", js.FuncOf(jsWrapper.send))
	module.Set("act", js.FuncOf(jsWrapper.act))

	events.AddEvent("Member")
	events.AddEvent("Error")

	jsWrapper.HandleEvents()

	<-done
}

func (jsWrapper *jsWrapper) update(event events.Event) {
	detail := make(map[string]interface{})
	data := make(map[string]interface{})
	data["event"] = event.Name
	data["data"] = event.Data
	detail["detail"] = data
	customEvent := js.Global().Get("CustomEvent").New(event.Name, detail)
	js.Global().Get("window").Call("dispatchEvent", customEvent)
}

func (jsWrapper *jsWrapper) register(event string) {

}

func (jsWrapper *jsWrapper) HandleEvents() {
	go func() {
		for {
			select {
			case event := <-events.Instance().HookChannel:
				jsWrapper.update(event)
			}
		}
	}()
}

func (jsWrapper *jsWrapper) status(this js.Value, args []js.Value) interface{} {
	status := make(map[string]interface{})
	var groups []interface{}
	for _, group := range jsWrapper.World.Groups {
		groupMap := make(map[string]interface{})
		groupMap["name"] = group.Name
		groupMap["id"] = group.ID.String()

		locations := []interface{}{}
		for _, location := range group.GetKnownLocations() {
			locations = append(locations, SerializeLocation(location.Location))
		}
		groupMap["knownLocations"] = locations

		var members []interface{}
		for _, member := range group.Members {
			members = append(members, SerializeHuman(member))
		}
		groupMap["members"] = members
		groups = append(groups, groupMap)
	}

	status["groups"] = groups
	status["world"] = SerializeWorld(jsWrapper.World)

	return ValueOf(status)
}

func ValueOf(x interface{}) js.Value {
	switch x := x.(type) {
	case map[string]interface{}:
		keys := make([]string, 0, len(x))
		for key := range x {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		o := js.Global().Get("Object").New()
		for _, k := range keys {
			o.Set(k, x[k])
		}
		return o
	default:
		return js.ValueOf(x)
	}
}

func (jsWrapper *jsWrapper) send(this js.Value, args []js.Value) interface{} {
	memberName := js.Value.String(args[0])
	location := js.Value.String(args[1])

	context := ActionContext{World: jsWrapper.World, Location: jsWrapper.World.GetLocationByName(location)}
	member := jsWrapper.World.Groups[0].GetMemberByName(memberName)
	fmt.Println(context)

	go Send(member, jsWrapper.World.GetLocationByName(location))
	return nil
}

func (jsWrapper *jsWrapper) act(this js.Value, args []js.Value) interface{} {
	member := jsWrapper.World.Groups[0].GetMemberByName(js.Value.String(args[0]))
	if member == nil {
		events.Emit("Error", "Member is not")
		return nil
	}

	skill, err := jsWrapper.parseArgs(js.Value.String(args[1]), member, args[2:])
	if err != nil {
		events.Emit("Error", err.Error())
		return nil
	}

	go DoAction(member, skill)
	return nil
}

func (jsWrapper *jsWrapper) scout(this js.Value, args []js.Value) interface{} {
	name := js.Value.String(args[0])
	fmt.Println("?")
	go Scout(name, jsWrapper.World)
	return nil
}

func gameLoop() {

}

func (jsWrapper *jsWrapper) parseArgs(skillType string, human *Human, args []js.Value) (Actionable, error) {
	fmt.Println(args)
	switch skillType {
	case "Mining":
		return (&Mining{}).Parse(human, args)
	case "Scouting":
		return (&Scouting{}).Parse(human, args)
	case "Woodcutting":
		return (&Woodcutting{}).Parse(human, args)
	}
	return nil, errors.New("No args passed")
}
