package gobpmn_counter

import (
	"reflect"
	"strings"

	"github.com/deemount/gobpmnCounter/internals/utils"
	gobpmn_reflection "github.com/deemount/gobpmnReflection"
)

// Quantities ...
type Quantities struct {
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

// In ...
func (q *Quantities) In(p interface{}) interface{} {

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
					q.countPool(field, name)
					q.countMessage(field, name)
					q.countElements(name)
				}
			}
		}
	case len(ref.Anonym) == 0:
		for _, field := range ref.Rflct {
			q.countProcess(field)
			q.countElements(field)
		}
	}

	q.countWords()

	return q

}

/*
 * @private
 */

// countPool ...
func (q *Quantities) countPool(field, reflectionField string) {
	if strings.ToLower(field) == "pool" {
		if strings.Contains(reflectionField, "Process") {
			q.Process++
		}
		if strings.Contains(reflectionField, "ID") {
			q.Participant++
			q.Shape++
		}
	}
}

// countMessage ...
func (q *Quantities) countMessage(field, reflectionField string) {
	if strings.ToLower(field) == "message" {
		if strings.Contains(reflectionField, "Message") {
			q.Message++
			q.Edge++
		}
	}
}

// countProcess ...
func (q *Quantities) countProcess(field string) {
	if strings.Contains(field, "Process") {
		q.Process++
	}
}

// countElements ...
func (q *Quantities) countElements(field string) {

	if utils.After(field, "From") == "" {

		switch true {
		case strings.Contains(field, "StartEvent"):
			q.StartEvent++
		case strings.Contains(field, "EndEvent"):
			q.EndEvent++
		case strings.Contains(field, "BusinessRuleTask"):
			q.BusinessRuleTask++
		case strings.Contains(field, "ManualTask"):
			q.ManualTask++
		case strings.Contains(field, "ReceiveTask"):
			q.ReceiveTask++
		case strings.Contains(field, "ScriptTask"):
			q.ScriptTask++
		case strings.Contains(field, "SendTask"):
			q.SendTask++
		case strings.Contains(field, "ServiceTask"):
			q.ServiceTask++
		case strings.Contains(field, "Task"):
			q.Task++
		case strings.Contains(field, "UserTask"):
			q.UserTask++
		}

		q.Shape++

	}
}

// countWords ...
func (q Quantities) countWords() {
	l := 0
	length := len(q.Words)
	for i := 0; i < length; i++ {
		l += len(q.Words[i])
	}
}
