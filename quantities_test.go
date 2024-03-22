package gobpmn_counter_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/deemount/gobpmnCounter/internals/utils"
)

// TestQuantities ...
type TestQuantities struct {
	StartEvent int
}

// TestUtilsAfter ...
func TestUtilsAfter(t *testing.T) {
	t.Run("TestUtilsAfter", func(t *testing.T) {
		field := "FromStartEvent"
		a := "StartEvent"
		b := utils.After(field, "From")
		assert.Equal(t, a, b, "The two words should be StartEvent")
	})
}

// TestCountElements ...
func TestCountElements(t *testing.T) {
	t.Run("TestCountElements", func(t *testing.T) {
		field := "StartEvent"
		q := TestQuantities{}
		if utils.After(field, "From") == "" {
			switch true {
			case strings.Contains(field, "StartEvent"):
				q.StartEvent++
			}
			assert.Equal(t, 1, q.StartEvent, "The StartEvent should be 1")
		}
	})
}
