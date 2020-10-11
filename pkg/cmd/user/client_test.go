package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	got := NewClient("mock")
	assert.NotNil(t, got)
}
