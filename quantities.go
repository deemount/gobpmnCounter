package gobpmn_counter

import (
	"reflect"
	"strings"

	"github.com/deemount/gobpmnCounter/internals/utils"
	gobpmn_reflection "github.com/deemount/gobpmnReflection"
)

// Quantities holds all the quantities of the BPMN elements
// in the BPMN model. It is used to count the number of elements
type Quantities struct {
	// How many processes are in the BPMN model
	Process int
	// How many participants are in the BPMN model
	// This will be counted by the number of pools
	Participant int
	// How many messages are in the BPMN model
	// This will be counted by the number of edges
	// and if the string contains the word "Message"
	Message int
	// How many elements are in the BPMN model
	ComplexGateway    int
	EventBasedGateway int
	ExclusiveGateway  int
	InclusiveGateway  int
	ParallelGateway   int
	// How many events are in the BPMN model
	BoundaryEvent          int
	EndEvent               int
	IntermediateCatchEvent int
	IntermediateThrowEvent int
	StartEvent             int
	// How many tasks are in the BPMN model
	BusinessRuleTask int
	ManualTask       int
	ReceiveTask      int
	ScriptTask       int
	SendTask         int
	ServiceTask      int
	Task             int
	UserTask         int
	// How many flows are in the BPMN model
	// This will be counted by the number of edges
	// and if the string contains a preposition, like
	// "From" or "To"
	Flow int
	// How many shapes and edges are in the BPMN model
	Shape int
	Edge  int
	// How many words are in the BPMN model
	Words map[int][]string
}

// In is a copy of the reflection package, but with the
// ability to count the number of elements in the BPMN model.
func (q *Quantities) In(p interface{}) *Quantities {

	ref := gobpmn_reflection.New(p)
	ref.Interface().Allocate().Maps().Assign()

	switch true {

	// If the BPMN model is a pool and has embedded structs
	case len(ref.Anonym) > 0:
		for _, field := range ref.Anonym {
			n := ref.Temporary.FieldByName(field)
			for i := 0; i < n.NumField(); i++ {
				name := n.Type().Field(i).Name
				switch n.Field(i).Kind() {
				case reflect.Struct:
					q.countPool(field, name)
					q.countMessage(field, name)
					q.countFlow(name)
					q.countElements(name)
				}
			}
		}

	// If the BPMN model is not a pool and has no embedded structs
	case len(ref.Anonym) == 0:
		for _, field := range ref.Rflct {
			q.countProcess(field)
			q.countFlow(field)
			q.countElements(field)
		}
	}

	// Count the number of words in the BPMN model
	q.countWords()

	return q

}

/*
 * @pprivate
 */

// countPool counts the number of processes and participants in the BPMN model.
// A pool is structured as a process and has participants and messages.
// Ruleset:
//   - If the field contains the word "Pool" and the reflection field contains the word "Process"
//     then it is a process.
//   - If the field contains the word "Pool" and the reflection field contains the word "ID"
//     then it is a participant.
//
// Note:
// The word "Pool" is case insensitive.
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

// countFlow counts all the flows in the BPMN model.
// Ruleset:
//   - If the field contains the word "From" then it is a flow.
func (q *Quantities) countFlow(field string) {
	if strings.Contains(field, "From") {
		q.Flow++
		q.Edge++
	}
}

// countElements counts all the elements in the BPMN model
// and increments the counter for each element.
// Ruleset:
//   - If the field contains one of the words below and without the word "From"
//     then it is an element.
//   - If the field contains the word from one of the words below
func (q *Quantities) countElements(field string) {

	if utils.After(field, "From") == "" {

		switch true {

		// events
		case strings.Contains(field, "StartEvent"):
			q.StartEvent++
		case strings.Contains(field, "BoundaryEvent"):
			q.BoundaryEvent++
		case strings.Contains(field, "IntermediateCatchEvent"):
			q.IntermediateCatchEvent++
		case strings.Contains(field, "IntermediateThrowEvent"):
			q.IntermediateThrowEvent++
		case strings.Contains(field, "EndEvent"):
			q.EndEvent++

		// gateways
		case strings.Contains(field, "ComplexGateway"):
			q.ComplexGateway++
		case strings.Contains(field, "EventBasedGateway"):
			q.EventBasedGateway++
		case strings.Contains(field, "ExclusiveGateway"):
			q.ExclusiveGateway++
		case strings.Contains(field, "InclusiveGateway"):
			q.InclusiveGateway++
		case strings.Contains(field, "ParallelGateway"):
			q.ParallelGateway++

		// tasks
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

		// each element in the switch has a shape
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
