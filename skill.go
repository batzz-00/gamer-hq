package main

import (
	"fmt"
	"time"
)

type MissingSkill struct {
	msg string
}

func NewMissingSkill(human *Human, skill Actionable) MissingSkill {
	return MissingSkill{msg: fmt.Sprintf("%s does not know the skill %s", human.Name, skill.Type().Name)}
}

func (MissingSkill MissingSkill) Error() string {
	return MissingSkill.msg
}

type Actionable interface {
	Type() Skill
	Time() time.Duration
	Check() error
	PreAction() error
	PostAction() error
}

type Skill struct {
	Name string
}

type Skillable interface {
	AddSkill(Actionable)
	Skills() []Actionable
}

func HasSkill(skillable Skillable, check Actionable) bool {
	for _, skill := range skillable.Skills() {
		if skill.Type().Name == check.Type().Name {
			return true
		}
	}
	return false
}

func getSkillByName(name string) Actionable {
	if skill, ok := skills[name]; ok {
		return skill
	}
	return nil
}
