package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pf ProgrammingFunctions = ProgrammingFunctions{}

func TestNewUuidWithHyphen(t *testing.T) {
	uuitWithHyphen := pf.NewUuid(false)

	assert.Len(t, uuitWithHyphen, 36)
	assert.Contains(t, uuitWithHyphen, "-")
}

func TestNewUuidWithoutHyphen(t *testing.T) {
	uuitWithHyphen := pf.NewUuid(true)

	assert.Len(t, uuitWithHyphen, 32)
	assert.NotContains(t, uuitWithHyphen, "-")
}