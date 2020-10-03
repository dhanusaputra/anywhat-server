package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResolver(t *testing.T) {
	r := NewResolver(nil, nil)
	assert.NotNil(t, r, t.Name)
}
