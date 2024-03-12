package gobpmn_counter

import (
	"reflect"
	"strings"

	"github.com/deemount/gobpmnCounter/internals/utils"
	gobpmn_reflection "github.com/deemount/gobpmnReflection"
)

// Count ...
type Count struct {
	Process          int
	Participant      int
	Message          int
	StartEvent       int
	EndEvent         int
	BusinessRuleTask int
	ManualTask       int
	ReceiveTask      int
	ScriptTask       int
	SendTask         int
	ServiceTask      int
	Task             int
	UserTask         int
	Flow             int
	Shape            int
	Edge             int
	Words            map[int][]string
}

// Counter ...
func (c *Count) Counter(p interface{}) interface{} {

	ref := gobpmn_reflection.New(p)
	ref.Interface().Allocate().Maps().Assign()

	switch true {
	case len(ref.Anonym) > 0:
		for _, field := range ref.Anonym {
			n := ref.Temporary.FieldByName(field)
			for i := 0; i < n.NumField(); i++ {
				name := n.Type().Field(i).Name
				switch n.Field(i).Kind() {
				case reflect.Struct:
					c.countPool(field, name)
					c.countMessage(field, name)
					c.countElements(name)
				}
			}
		}
	case len(ref.Anonym) == 0:
		for _, field := range ref.Rflct {
			c.countProcess(field)
			c.countElements(field)
		}
	}

	c.countWords()

	return c

}

/*
 * @Base
 */

// countPool ...
func (c *Count) countPool(field, reflectionField string) {
	if strings.ToLower(field) == "pool" {
		if strings.Contains(reflectionField, "Process") {
			c.Process++
		}
		if strings.Contains(reflectionField, "ID") {
			c.Participant++
			c.Shape++
		}
	}
}

// countMessage ...
func (c *Count) countMessage(field, reflectionField string) {
	if strings.ToLower(field) == "message" {
		if strings.Contains(reflectionField, "Message") {
			c.Message++
			c.Edge++
		}
	}
}

/*
 * @Processes
 */

// countProcess ...
func (c *Count) countProcess(field string) {
	if strings.Contains(field, "Process") {
		c.Process++
	}
}

/*
 * @Elements
 */

func (c *Count) countElements(field string) {

	if utils.After(field, "From") == "" {

		switch true {
		case strings.Contains(field, "StartEvent"):
			c.StartEvent++
		case strings.Contains(field, "EndEvent"):
			c.EndEvent++
		case strings.Contains(field, "BusinessRuleTask"):
			c.BusinessRuleTask++
		case strings.Contains(field, "ManualTask"):
			c.ManualTask++
		case strings.Contains(field, "ReceiveTask"):
			c.ReceiveTask++
		case strings.Contains(field, "ScriptTask"):
			c.ScriptTask++
		case strings.Contains(field, "SendTask"):
			c.SendTask++
		case strings.Contains(field, "ServiceTask"):
			c.ServiceTask++
		case strings.Contains(field, "Task"):
			c.Task++
		case strings.Contains(field, "UserTask"):
			c.UserTask++
		}

		c.Shape++

	}
}

/*
 * @Words
 */

// countWords ...
func (c Count) countWords() {
	l := 0
	length := len(c.Words)
	for i := 0; i < length; i++ {
		l += len(c.Words[i])
	}
}
