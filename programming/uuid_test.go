package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUuidWithHyphen(t *testing.T) {
	uuitWithHyphen := NewUuid(false)

	assert.Len(t, uuitWithHyphen, 36)
	assert.Contains(t, uuitWithHyphen, "-")
}

func TestNewUuidWithoutHyphen(t *testing.T) {
	uuitWithHyphen := NewUuid(true)

	assert.Len(t, uuitWithHyphen, 32)
	assert.NotContains(t, uuitWithHyphen, "-")
}